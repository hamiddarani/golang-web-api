package handlers

import (
	"golang-web-api/api/dtos"
	"golang-web-api/api/helper"
	"golang-web-api/config"
	"golang-web-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(cfg *config.Config) *UserHandler {
	service := services.NewUserService(cfg)
	return &UserHandler{service: service}
}

// SendOtp godoc
// @Summary Send otp to user
// @Description Send otp to user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dtos.GetOtpRequest true "GetOtpRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/send-otp [post]
func (h *UserHandler) SendOtp(ctx *gin.Context) {
	req := new(dtos.GetOtpRequest)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}

	err = h.service.SendOtp(req)
	if err != nil {
		ctx.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	// Call internal SMS service
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 0))

}

// RegisterLoginByMobileNumber godoc
// @Summary RegisterLoginByMobileNumber
// @Description RegisterLoginByMobileNumber
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dtos.RegisterLoginByMobileRequest true "RegisterLoginByMobileRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/login [post]
func (h *UserHandler) RegisterLoginByMobileNumber(c *gin.Context) {
	req := new(dtos.RegisterLoginByMobileRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	token, err := h.service.RegisterLoginByMobileNumber(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(token, true, 0))
}

// RefreshToken godoc
// @Summary RefreshToken
// @Description RefreshToken
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dtos.RefreshToken true "RefreshToken"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/refresh-token [post]
func (h *UserHandler) RefreshToken(c *gin.Context){
	req := new(dtos.RefreshToken)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}

	token, err := h.service.RefreshToken(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(token, true, 0))
}
