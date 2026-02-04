package meal

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SnackLog/meal-service/internal/database/models"
	"github.com/SnackLog/meal-service/internal/database/queries"
	"github.com/SnackLog/meal-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

type mealGetResponse struct {
	Meals []models.Meal `json:"meals"`
}

// Get a meal list by date
// @Summary Get meals by date
// @Description Get meals for a user filtered by date range
// @Tags meals
// @Accept json
// @Produce json
// @Param date_not_after query string false "Get meals not after this date (YYYY-MM-DD)"
// @Param date_not_before query string false "Get meals not before this date (YYYY-MM-DD)"
// @Success 200 {object} mealGetResponse
// @Failure 400 {object} handlers.Error
// @Failure 500 {object} handlers.Error
// @Security     BearerAuth
// @Router /meal [get]
func (mc *MealController) Get(c *gin.Context) {
	username := c.GetString("username")

	stringDateNotAfter := c.Query("date_not_after")
	stringDateNotBefore := c.Query("date_not_before")

	dateNotAfter, dateNotBefore, err := parseDates(stringDateNotAfter, c, stringDateNotBefore)
	if err != nil {
		log.Println("Error parsing dates:", err)
		return
	}

	meals, err := getMeals(mc.DB, username, dateNotAfter, dateNotBefore)
	if err != nil {
		log.Println("Error retrieving meals by date:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, mealGetResponse{Meals: meals})
}

func parseDates(stringDateNotAfter string, c *gin.Context, stringDateNotBefore string) (time.Time, time.Time, error) {
	var dateNotAfter, dateNotBefore time.Time
	var err error

	if stringDateNotAfter != "" {
		dateNotAfter, err = time.Parse("2006-01-02", stringDateNotAfter)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, handlers.Error{Error: fmt.Sprintf("invalid date_not_after format: %v", err)})
			return time.Time{}, time.Time{}, fmt.Errorf("invalid date_not_after format: %v", err)
		}
	}

	if stringDateNotBefore != "" {
		dateNotBefore, err = time.Parse("2006-01-02", stringDateNotBefore)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, handlers.Error{Error: fmt.Sprintf("invalid date_not_before format: %v", err)})
			return time.Time{}, time.Time{}, fmt.Errorf("invalid date_not_before format: %v", err)
		}
	}
	return dateNotAfter, dateNotBefore, nil
}

func getMeals(db *sql.DB, username string, dateNotAfter time.Time, dateNotBefore time.Time) ([]models.Meal, error) {
	if !dateNotAfter.IsZero() && !dateNotBefore.IsZero() {
		return queries.GetMealsByDate(db, username, dateNotAfter, dateNotBefore)
	} else if !dateNotBefore.IsZero() {
		return queries.GetMealsByDateNotBefore(db, username, dateNotBefore)
	} else if !dateNotAfter.IsZero() {
		return queries.GetMealsByDateNotAfter(db, username, dateNotAfter)
	}
	return nil, fmt.Errorf("at least one of dateNotAfter or dateNotBefore must be provided")
}
