package services

import (
	"golang-web-api/api/dtos"
	"golang-web-api/config"
	"golang-web-api/data/db"
	"golang-web-api/data/models"
	"golang-web-api/pkg/logging"

	"github.com/gin-gonic/gin"
)

type PropertyCategoryService struct {
	base *BaseService[models.PropertyCategory, dtos.CreatePropertyCategoryRequest, dtos.UpdatePropertyCategoryRequest, dtos.PropertyCategoryResponse]
}

func NewPropertyCategoryService(cfg *config.Config) *PropertyCategoryService {
	return &PropertyCategoryService{
		base: &BaseService[models.PropertyCategory, dtos.CreatePropertyCategoryRequest, dtos.UpdatePropertyCategoryRequest, dtos.PropertyCategoryResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{{string: "Properties"}},
		},
	}
}

// Create
func (s *PropertyCategoryService) Create(ctx *gin.Context, req *dtos.CreatePropertyCategoryRequest) (*dtos.PropertyCategoryResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *PropertyCategoryService) Update(ctx *gin.Context, id int, req *dtos.UpdatePropertyCategoryRequest) (*dtos.PropertyCategoryResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *PropertyCategoryService) Delete(ctx *gin.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *PropertyCategoryService) GetById(ctx *gin.Context, id int) (*dtos.PropertyCategoryResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *PropertyCategoryService) GetByFilter(ctx *gin.Context, req *dtos.PaginationInputWithFilter) (*dtos.PagedList[dtos.PropertyCategoryResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
