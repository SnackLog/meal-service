package meal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (mc *MealController) GetId(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}
