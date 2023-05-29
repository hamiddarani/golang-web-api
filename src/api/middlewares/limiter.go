package middlewares

import (
	"golang-web-api/api/helper"
	"net/http"

	"github.com/didip/tollbooth/v7"
	"github.com/gin-gonic/gin"
)

func LimitByRequest() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)
	return func(ctx *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithError(nil, false, 429, err))
			return
		}
		ctx.Next()

	}
}
