package meal

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SnackLog/meal-service/internal/database/queries"
	"github.com/SnackLog/meal-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

// Delete godoc
// @Summary      Delete a meal
// @Description  Delete a meal by ID
// @Tags         meals
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Meal ID"
// @Success      204  "No Content"
// @Failure      400  {object}  handlers.Error
// @Failure      401  {object}  handlers.Error
// @Failure      404  {object}  handlers.Error
// @Failure      500  {object}  handlers.Error
// @Security     BearerAuth
// @Router       /meal/{id} [delete]
func (mc *MealController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Invalid meal ID: %v", err)
		c.JSON(http.StatusBadRequest, handlers.Error{Error: "Invalid meal ID"})
		return
	}

	result, err := queries.DeleteMealById(mc.DB, id, c.GetString("username"))
	if err != nil {
		log.Printf("Error deleting meal: %v", err)
		c.JSON(http.StatusInternalServerError, handlers.Error{Error: "Failed to delete meal"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		c.JSON(http.StatusInternalServerError, handlers.Error{Error: "Failed to determine deletion result"})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, handlers.Error{Error: "Meal not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
