package services

import (
	"golang-web-api/api/dtos"
	"golang-web-api/config"
	"golang-web-api/data/db"
	"golang-web-api/data/models"
	"golang-web-api/pkg/logging"

	"github.com/gin-gonic/gin"
)

type CarModelPriceHistoryService struct {
	base *BaseService[models.CarModelPriceHistory, dtos.CreateCarModelPriceHistoryRequest, dtos.UpdateCarModelPriceHistoryRequest, dtos.CarModelPriceHistoryResponse]
}

func NewCarModelPriceHistoryService(cfg *config.Config) *CarModelPriceHistoryService {
	return &CarModelPriceHistoryService{
		base: &BaseService[models.CarModelPriceHistory, dtos.CreateCarModelPriceHistoryRequest, dtos.UpdateCarModelPriceHistoryRequest, dtos.CarModelPriceHistoryResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// Create
func (s *CarModelPriceHistoryService) Create(ctx *gin.Context, req *dtos.CreateCarModelPriceHistoryRequest) (*dtos.CarModelPriceHistoryResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelPriceHistoryService) Update(ctx *gin.Context, id int, req *dtos.UpdateCarModelPriceHistoryRequest) (*dtos.CarModelPriceHistoryResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelPriceHistoryService) Delete(ctx *gin.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelPriceHistoryService) GetById(ctx *gin.Context, id int) (*dtos.CarModelPriceHistoryResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarModelPriceHistoryService) GetByFilter(ctx *gin.Context, req *dtos.PaginationInputWithFilter) (*dtos.PagedList[dtos.CarModelPriceHistoryResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
