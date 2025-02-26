package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Data   any    `json:"data"`
	Result bool   `json:"result"`
}

func JsonToFront(c *gin.Context, ApiResponse *ApiResponse) {
	c.JSON(http.StatusOK, ApiResponse)
}
