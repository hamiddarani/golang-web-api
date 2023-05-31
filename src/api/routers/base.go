package routers

import (
	"golang-web-api/api/handlers"
	"golang-web-api/config"

	"github.com/gin-gonic/gin"
)

func Country(router *gin.RouterGroup, cfg *config.Config) {
	handler := handlers.NewCountryHandler(cfg)

	router.POST("/", handler.Create)
	router.PUT("/:id", handler.Update)
	router.DELETE("/:id", handler.Delete)
	router.GET("/:id", handler.GetById)
}