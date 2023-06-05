package services

import (
	"golang-web-api/api/dtos"
	"golang-web-api/config"
	"golang-web-api/data/db"
	"golang-web-api/data/models"
	"golang-web-api/pkg/logging"

	"github.com/gin-gonic/gin"
)

type CarModelColorService struct {
	base *BaseService[models.CarModelColor, dtos.CreateCarModelColorRequest, dtos.UpdateCarModelColorRequest, dtos.CarModelColorResponse]
}

func NewCarModelColorService(cfg *config.Config) *CarModelColorService {
	return &CarModelColorService{
		base: &BaseService[models.CarModelColor, dtos.CreateCarModelColorRequest, dtos.UpdateCarModelColorRequest, dtos.CarModelColorResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Color"},
			},
		},
	}
}

// Create
func (s *CarModelColorService) Create(ctx *gin.Context, req *dtos.CreateCarModelColorRequest) (*dtos.CarModelColorResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelColorService) Update(ctx *gin.Context, id int, req *dtos.UpdateCarModelColorRequest) (*dtos.CarModelColorResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelColorService) Delete(ctx *gin.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelColorService) GetById(ctx *gin.Context, id int) (*dtos.CarModelColorResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarModelColorService) GetByFilter(ctx *gin.Context, req *dtos.PaginationInputWithFilter) (*dtos.PagedList[dtos.CarModelColorResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
