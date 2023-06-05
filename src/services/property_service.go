package services

import (
	"golang-web-api/api/dtos"
	"golang-web-api/config"
	"golang-web-api/data/db"
	"golang-web-api/data/models"
	"golang-web-api/pkg/logging"

	"github.com/gin-gonic/gin"
)

type PropertyService struct {
	base *BaseService[models.Property, dtos.CreatePropertyRequest, dtos.UpdatePropertyRequest, dtos.PropertyResponse]
}

func NewPropertyService(cfg *config.Config) *PropertyService {
	return &PropertyService{
		base: &BaseService[models.Property, dtos.CreatePropertyRequest, dtos.UpdatePropertyRequest, dtos.PropertyResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{{string: "Category"}},
		},
	}
}

// Create
func (s *PropertyService) Create(ctx *gin.Context, req *dtos.CreatePropertyRequest) (*dtos.PropertyResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *PropertyService) Update(ctx *gin.Context, id int, req *dtos.UpdatePropertyRequest) (*dtos.PropertyResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *PropertyService) Delete(ctx *gin.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *PropertyService) GetById(ctx *gin.Context, id int) (*dtos.PropertyResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *PropertyService) GetByFilter(ctx *gin.Context, req *dtos.PaginationInputWithFilter) (*dtos.PagedList[dtos.PropertyResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
