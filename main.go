package main

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/SnackLog/meal-service/internal/database"
	"github.com/SnackLog/meal-service/internal/handlers/meal"
	"github.com/SnackLog/meal-service/internal/health"
	"github.com/gin-gonic/gin"

	authLib "github.com/SnackLog/auth-lib"
	databaseConfig "github.com/SnackLog/database-config-lib"
	serviceConfig "github.com/SnackLog/service-config-lib"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

// @title           Meal Service API
// @version         1.0
// @description     This is the API for the Meal Service.
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	loadConfigs()
	migrateDatabase()

	db := connectDB()
	defer db.Close()

	setupEndpoints(db)
}

// setupEndpoints sets up endpoints for the gin router engine
func setupEndpoints(db *sql.DB) {
	engine := gin.Default()

	setupMealEndpoints(engine, db)
	setupHealthEndpoints(engine, db)

	engine.Run(":80")

}

// setupHealthEndpoints sets up the healthcheck endpoint
func setupHealthEndpoints(engine *gin.Engine, db *sql.DB) {
	hc := health.HealthController{DB: db}
	engine.GET("/health", hc.GetHealthStatus)
}

// setupMealEndpoints sets up the meal endpoints
func setupMealEndpoints(engine *gin.Engine, db *sql.DB) {
	mc := meal.MealController{DB: db}
	engine.GET("/meal/:id", authLib.Authentication, mc.GetID)
	engine.POST("/meal", authLib.Authentication, mc.Post)
	engine.DELETE("/meal/:id", authLib.Authentication, mc.Delete)
}

// connectDB connects to the DB and returns the *sql.DB object returned
func connectDB() *sql.DB {
	db, err := database.Connect(databaseConfig.GetDatabaseConnectionString())
	if err != nil {
		panic(err)
	}
	return db
}

// loadConfigs loads the configurations for service and database from ENV
func loadConfigs() {
	err := serviceConfig.LoadConfig()
	if err != nil {
		panic(err)
	}

	err = databaseConfig.LoadConfig()
	if err != nil {
		panic(err)
	}
}

func migrateDatabase() {
	err := doMigrations()
	if err != nil {
		panic(fmt.Sprintf("Database migration failed: %v", err))
	}
}

// migrationFiles embeds SQL migration files.
//
//go:embed db/migrations/*.sql
var migrationFiles embed.FS

// doMigrations performs database migrations using embedded SQL files.
func doMigrations() error {
	migrationDriver, err := iofs.New(migrationFiles, "db/migrations")
	if err != nil {
		return fmt.Errorf("Failed to create iofs driver: %v", err)
	}

	m, err := migrate.NewWithSourceInstance(
		"iofs",
		migrationDriver,
		databaseConfig.GetDatabaseConnectionString(),
	)

	if err != nil {
		return fmt.Errorf("Failed to create migrate instance: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("Failed to run migrations: %v", err)
	}

	return nil
}
