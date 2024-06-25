package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthy(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"server on": "Bom dia gliciaaaaaaaaaaaaaa",
	})
}
