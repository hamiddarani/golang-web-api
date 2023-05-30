package routers

import (
	"golang-web-api/api/handlers"
	"golang-web-api/api/middlewares"
	"golang-web-api/config"

	"github.com/gin-gonic/gin"
)

func Users(r *gin.RouterGroup, cfg *config.Config) {
	handler := handlers.NewUserHandler(cfg)

	r.POST("/send-otp", middlewares.OtpLimiter(cfg), handler.SendOtp)
}
