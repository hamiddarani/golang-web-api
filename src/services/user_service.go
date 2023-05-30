package services

import (
	"golang-web-api/api/dtos"
	"golang-web-api/common"
	"golang-web-api/config"
	"golang-web-api/data/db"
	"golang-web-api/pkg/logging"

	"gorm.io/gorm"
)

type UserService struct {
	logger     logging.Logger
	cfg        *config.Config
	otpService *OtpService
	database   *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDb()
	logger := logging.NewLogger(cfg)
	return &UserService{
		cfg:        cfg,
		database:   database,
		logger:     logger,
		otpService: NewOtpService(cfg),
	}
}

func (s *UserService) SendOtp(req *dtos.GetOtpRequest) error {
	otp := common.GenerateOtp()
	err := s.otpService.SetOtp(otp, req.MobileNumber)
	if err != nil {
		return err
	}
	return nil
}
