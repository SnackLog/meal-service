package meal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (mc *MealController) GetID(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}
