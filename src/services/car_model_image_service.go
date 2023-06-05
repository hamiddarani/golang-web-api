package services

import (
	"golang-web-api/api/dtos"
	"golang-web-api/config"
	"golang-web-api/data/db"
	"golang-web-api/data/models"
	"golang-web-api/pkg/logging"

	"github.com/gin-gonic/gin"
)

type CarModelImageService struct {
	base *BaseService[models.CarModelImage, dtos.CreateCarModelImageRequest, dtos.UpdateCarModelImageRequest, dtos.CarModelImageResponse]
}

func NewCarModelImageService(cfg *config.Config) *CarModelImageService {
	return &CarModelImageService{
		base: &BaseService[models.CarModelImage, dtos.CreateCarModelImageRequest, dtos.UpdateCarModelImageRequest, dtos.CarModelImageResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Image"},
			},
		},
	}
}

// Create
func (s *CarModelImageService) Create(ctx *gin.Context, req *dtos.CreateCarModelImageRequest) (*dtos.CarModelImageResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelImageService) Update(ctx *gin.Context, id int, req *dtos.UpdateCarModelImageRequest) (*dtos.CarModelImageResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelImageService) Delete(ctx *gin.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelImageService) GetById(ctx *gin.Context, id int) (*dtos.CarModelImageResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarModelImageService) GetByFilter(ctx *gin.Context, req *dtos.PaginationInputWithFilter) (*dtos.PagedList[dtos.CarModelImageResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
