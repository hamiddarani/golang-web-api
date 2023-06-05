package services

import (
	"golang-web-api/api/dtos"
	"golang-web-api/config"
	"golang-web-api/data/db"
	"golang-web-api/data/models"
	"golang-web-api/pkg/logging"

	"github.com/gin-gonic/gin"
)

type ColorService struct {
	base *BaseService[models.Color, dtos.CreateColorRequest, dtos.UpdateColorRequest, dtos.ColorResponse]
}

func NewColorService(cfg *config.Config) *ColorService {
	return &ColorService{
		base: &BaseService[models.Color, dtos.CreateColorRequest, dtos.UpdateColorRequest, dtos.ColorResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// Create
func (s *ColorService) Create(ctx *gin.Context, req *dtos.CreateColorRequest) (*dtos.ColorResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *ColorService) Update(ctx *gin.Context, id int, req *dtos.UpdateColorRequest) (*dtos.ColorResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *ColorService) Delete(ctx *gin.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *ColorService) GetById(ctx *gin.Context, id int) (*dtos.ColorResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *ColorService) GetByFilter(ctx *gin.Context, req *dtos.PaginationInputWithFilter) (*dtos.PagedList[dtos.ColorResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
