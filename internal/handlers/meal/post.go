package meal

import (
	"log"
	"net/http"

	"github.com/SnackLog/meal-service/internal/database/models"
	"github.com/SnackLog/meal-service/internal/database/queries"
	"github.com/gin-gonic/gin"
)

func (mc *MealController) Post(c *gin.Context) {
	var requestBody models.Meal
	if err := c.ShouldBind(&requestBody); err != nil {
		log.Println("Error binding request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return
	}

	username := c.GetString("username")
	mealID, err := queries.CreateMeal(mc.DB, username, requestBody)
	if err != nil {
		log.Println("Error creating meal:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create meal."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"meal_id": mealID})
}
