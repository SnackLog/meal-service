package meal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (mc *MealController) Delete(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}
