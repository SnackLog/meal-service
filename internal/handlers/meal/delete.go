package meal

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SnackLog/meal-service/internal/database/queries"
	"github.com/gin-gonic/gin"
)

func (mc *MealController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid meal ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid meal ID"})
		return
	}

	result, err := queries.DeleteMealById(mc.DB, id, c.GetString("username"))
	if err != nil {
		log.Printf("Error deleting meal: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete meal"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to determine deletion result"})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Meal not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
