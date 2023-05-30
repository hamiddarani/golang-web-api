package services

import (
	"golang-web-api/api/dtos"
	"golang-web-api/common"
	"golang-web-api/config"
	"golang-web-api/constants"
	"golang-web-api/data/db"
	"golang-web-api/data/models"
	"golang-web-api/pkg/logging"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type UserService struct {
	logger       logging.Logger
	cfg          *config.Config
	otpService   *OtpService
	database     *gorm.DB
	tokenService *TokenService
}

func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDb()
	logger := logging.NewLogger(cfg)
	return &UserService{
		cfg:          cfg,
		database:     database,
		logger:       logger,
		tokenService: NewTokenService(cfg),
		otpService:   NewOtpService(cfg),
	}
}

func (s *UserService) RegisterLoginByMobileNumber(req *dtos.RegisterLoginByMobileRequest) (*dtos.TokenDetail, error) {
	err := s.otpService.ValidateOtp(req.MobileNumber, req.Otp)
	if err != nil {
		return nil, err
	}
	exists, err := s.existsByMobileNumber(req.MobileNumber)
	if err != nil {
		return nil, err
	}

	u := models.User{MobileNumber: req.MobileNumber}

	if exists {
		var user models.User
		err = s.database.
			Model(&models.User{}).
			Where("mobile_number = ?", u.MobileNumber).
			Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
				return tx.Preload("Role")
			}).
			Find(&user).Error
		if err != nil {
			return nil, err
		}
		tdto := tokenDto{UserId: user.Id, MobileNumber: user.MobileNumber}

		if len(*user.UserRoles) > 0 {
			for _, ur := range *user.UserRoles {
				tdto.Roles = append(tdto.Roles, ur.Role.Name)
			}
		}

		token, err := s.tokenService.GenerateToken(&tdto)
		if err != nil {
			return nil, err
		}
		return token, nil

	}

	roleId, err := s.getDefaultRole()
	if err != nil {
		s.logger.Error(logging.Postgres, logging.DefaultRoleNotFound, err.Error(), nil)
		return nil, err
	}

	tx := s.database.Begin()
	err = tx.Create(&u).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return nil, err
	}
	err = tx.Create(&models.UserRole{RoleId: roleId, UserId: u.Id}).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return nil, err
	}
	tx.Commit()

	var user models.User
	err = s.database.
		Model(&models.User{}).
		Where("mobile_number = ?", u.MobileNumber).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).
		Find(&user).Error
	if err != nil {
		return nil, err
	}
	tdto := tokenDto{UserId: user.Id, MobileNumber: user.MobileNumber}

	if len(*user.UserRoles) > 0 {
		for _, ur := range *user.UserRoles {
			tdto.Roles = append(tdto.Roles, ur.Role.Name)
		}
	}

	token, err := s.tokenService.GenerateToken(&tdto)
	if err != nil {
		return nil, err
	}
	return token, nil

}

func (s *UserService) RefreshToken(req *dtos.RefreshToken) (*dtos.AccessToken, error) {
	claimMap, err := s.tokenService.GetClaims(req.RefreshToken)
	if err != nil {
		return nil, err
	}
	td := &dtos.AccessToken{}
	td.AccessTokenExpireTime = time.Now().Add(s.cfg.JWT.AccessTokenExpireDuration * time.Minute).Unix()

	atc := jwt.MapClaims{}

	atc[constants.UserIdKey] = claimMap[constants.UserIdKey]
	atc[constants.MobileNumberKey] = claimMap[constants.MobileNumberKey]
	atc[constants.RolesKey] = claimMap[constants.RolesKey]
	atc[constants.ExpireTimeKey] = td.AccessTokenExpireTime

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)

	td.AccessToken, err = at.SignedString([]byte(s.cfg.JWT.Secret))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (s *UserService) SendOtp(req *dtos.GetOtpRequest) error {
	otp := common.GenerateOtp()
	err := s.otpService.SetOtp(otp, req.MobileNumber)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) existsByMobileNumber(mobileNumber string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("mobile_number = ?", mobileNumber).
		Find(&exists).
		Error; err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserService) getDefaultRole() (roleId int, err error) {

	if err = s.database.Model(&models.Role{}).
		Select("id").
		Where("name = ?", constants.DefaultRoleName).
		First(&roleId).Error; err != nil {
		return 0, err
	}
	return roleId, nil
}
