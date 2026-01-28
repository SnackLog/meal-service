package meal

import (
	"log"
	"net/http"

	"github.com/SnackLog/meal-service/internal/database/models"
	"github.com/SnackLog/meal-service/internal/database/queries"
	"github.com/SnackLog/meal-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

type mealPostResponse struct {
	MealId int `json:"meal_id"`
}

// Post godoc
// @Summary      Create a meal
// @Description  Create a new meal
// @Tags         meals
// @Accept       json
// @Produce      json
// @Param        meal  body      models.Meal  true  "Meal Data"
// @Success      201   {object}  map[string]int
// @Failure      400   {object}  handlers.Error
// @Failure      401   {object}  handlers.Error
// @Failure      500   {object}  handlers.Error
// @Security     BearerAuth
// @Router       /meal [post]
func (mc *MealController) Post(c *gin.Context) {
	var requestBody models.Meal
	if err := c.ShouldBind(&requestBody); err != nil {
		log.Println("Error binding request body:", err)
		c.JSON(http.StatusBadRequest, handlers.Error{Error: "Invalid request body."})
		return
	}

	username := c.GetString("username")
	mealID, err := queries.CreateMeal(mc.DB, username, requestBody)
	if err != nil {
		log.Println("Error creating meal:", err)
		c.JSON(http.StatusInternalServerError, handlers.Error{Error: "Failed to create meal"})
		return
	}

	c.JSON(http.StatusCreated, mealPostResponse{MealId: mealID})
}
