package api

import (
	"fmt"
	"golang-web-api/api/middlewares"
	"golang-web-api/api/routers"
	"golang-web-api/api/validations"
	"golang-web-api/config"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer(cfg *config.Config) {
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()
	RegisterValidators()
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest())

	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
		if err != nil {
			log.Print(err.Error())
		}
	}
}
