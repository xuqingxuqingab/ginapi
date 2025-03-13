package learn

import (
	"ginapi/app/api/learn"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Learn1(c *gin.Context) {
	param := learn.Learn1Req{}
	err := c.ShouldBind(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}
