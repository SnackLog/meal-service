package meal

import (
	"log"
	"net/http"
	"strconv"

	"github.com/SnackLog/meal-service/internal/database/queries"
	"github.com/gin-gonic/gin"
)

// GetID godoc
// @Summary      Get a meal
// @Description  Get a meal by ID
// @Tags         meals
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Meal ID"
// @Success      200  {object}  models.Meal
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security     BearerAuth
// @Router       /meal/{id} [get]
func (mc *MealController) GetID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error converting meal ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid meal ID"})
		return
	}

	meal, err := queries.GetMealById(mc.DB, id, c.GetString("username"))
	if err != nil {
		log.Printf("Error retrieving meal by ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve meal"})
		return
	}
	if meal == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Meal not found"})
		return
	}

	c.JSON(http.StatusOK, meal)
}
