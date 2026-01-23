package meal

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SnackLog/meal-service/internal/database/queries"
	"github.com/gin-gonic/gin"
)

// Delete godoc
// @Summary      Delete a meal
// @Description  Delete a meal by ID
// @Tags         meals
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Meal ID"
// @Success      204  {object}  nil
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security     BearerAuth
// @Router       /meal/{id} [delete]
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
