package middlewares

import (
	"errors"
	"golang-web-api/api/helper"
	"golang-web-api/config"
	"golang-web-api/pkg/limiter"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func OtpLimiter(cfg *config.Config) gin.HandlerFunc {
	var limiter = limiter.NewIPRateLimiter(rate.Every(cfg.Otp.Limiter*time.Second), 1)
	return func(c *gin.Context) {
		limiter := limiter.GetLimiter(c.Request.RemoteAddr)
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithError(nil, false, -1, errors.New("not allowed")))
			c.Abort()
		} else {
			c.Next()
		}
	}
}
