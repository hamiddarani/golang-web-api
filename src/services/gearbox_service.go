package services

import (
	"golang-web-api/api/dtos"
	"golang-web-api/config"
	"golang-web-api/data/db"
	"golang-web-api/data/models"
	"golang-web-api/pkg/logging"

	"github.com/gin-gonic/gin"
)

type GearboxService struct {
	base *BaseService[models.Gearbox, dtos.CreateGearboxRequest, dtos.UpdateGearboxRequest, dtos.GearboxResponse]
}

func NewGearboxService(cfg *config.Config) *GearboxService {
	return &GearboxService{
		base: &BaseService[models.Gearbox, dtos.CreateGearboxRequest, dtos.UpdateGearboxRequest, dtos.GearboxResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// Create
func (s *GearboxService) Create(ctx *gin.Context, req *dtos.CreateGearboxRequest) (*dtos.GearboxResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *GearboxService) Update(ctx *gin.Context, id int, req *dtos.UpdateGearboxRequest) (*dtos.GearboxResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *GearboxService) Delete(ctx *gin.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *GearboxService) GetById(ctx *gin.Context, id int) (*dtos.GearboxResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *GearboxService) GetByFilter(ctx *gin.Context, req *dtos.PaginationInputWithFilter) (*dtos.PagedList[dtos.GearboxResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
