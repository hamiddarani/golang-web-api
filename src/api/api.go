package api

import (
	"golang-web-api/api/middlewares"
	"golang-web-api/api/routers"
	"golang-web-api/api/validations"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer() {
	r := gin.New()
	RegisterValidators()
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest())

	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}

	r.Run(":5000")
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
