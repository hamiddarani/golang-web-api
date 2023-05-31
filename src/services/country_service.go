package services

import (
	"database/sql"
	"fmt"
	"golang-web-api/api/dtos"
	"golang-web-api/config"
	"golang-web-api/constants"
	"golang-web-api/data/db"
	"golang-web-api/data/models"
	"golang-web-api/pkg/logging"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CountryService struct {
	logger   logging.Logger
	database *gorm.DB
}

func NewCountryService(cfg *config.Config) *CountryService {
	return &CountryService{
		logger:   logging.NewLogger(cfg),
		database: db.GetDb(),
	}
}

func (s *CountryService) Create(ctx *gin.Context, req *dtos.CreateUpdateCountryRequest) (*dtos.CountryResponse, error) {
	country := models.Country{Name: req.Name}
	fmt.Printf("%s", ctx.Value(constants.UserIdKey))
	country.CreatedBy = int(ctx.Value(constants.UserIdKey).(float64))
	country.CreatedAt = time.Now().UTC()

	tx := s.database.WithContext(ctx).Begin()
	if err := tx.Create(&country).Error; err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	dto := &dtos.CountryResponse{Name: country.Name}
	return dto, nil
}

func (s *CountryService) Update(ctx *gin.Context, id int, req *dtos.CreateUpdateCountryRequest) (*dtos.CountryResponse, error) {
	updateMap := map[string]interface{}{
		"Name":        req.Name,
		"modified_by": &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true},
		"modified_at": &sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}

	tx := s.database.WithContext(ctx).Begin()
	if err := tx.Model(&models.Country{}).Where("id = ?", id).Updates(updateMap).Error; err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	dto, err := s.GetById(ctx, id)
	if err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return nil, err
	}
	return dto, nil
}

func (s *CountryService) Delete(ctx *gin.Context, id int) error {
	updateMap := map[string]interface{}{
		"deleted_by": &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true},
		"deleted_at": &sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}

	tx := s.database.WithContext(ctx).Begin()
	if err := tx.Model(&models.Country{}).Where("id = ?", id).Updates(updateMap).Error; err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return err
	}
	tx.Commit()
	return nil
}

func (s *CountryService) GetById(ctx *gin.Context, id int) (*dtos.CountryResponse, error) {
	country := &models.Country{}

	if err := s.database.Model(&models.Country{}).Where("id = ?", id).First(&country).Error; err != nil {
		return nil, err
	}
	dto := &dtos.CountryResponse{Id: country.Id, Name: country.Name}
	return dto, nil
}
