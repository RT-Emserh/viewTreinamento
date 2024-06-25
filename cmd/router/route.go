package router

import (
	"github.com/gin-gonic/gin"
	"viewportal.com/cmd/api/controller"
	limitbucket "viewportal.com/pkg/middlewares"
)

func Router(r *gin.Engine) {
	r.Use(limitbucket.RateLimiter())
	r.GET("/viewbook/:ebook", controller.GetEbook)
	r.GET("/healthy", controller.Healthy)
}
