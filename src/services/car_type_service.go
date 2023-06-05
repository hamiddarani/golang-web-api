package services

import (
	"golang-web-api/api/dtos"
	"golang-web-api/config"
	"golang-web-api/data/db"
	"golang-web-api/data/models"
	"golang-web-api/pkg/logging"

	"github.com/gin-gonic/gin"
)

type CarTypeService struct {
	base *BaseService[models.CarType, dtos.CreateCarTypeRequest, dtos.UpdateCarTypeRequest, dtos.CarTypeResponse]
}

func NewCarTypeService(cfg *config.Config) *CarTypeService {
	return &CarTypeService{
		base: &BaseService[models.CarType, dtos.CreateCarTypeRequest, dtos.UpdateCarTypeRequest, dtos.CarTypeResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// Create
func (s *CarTypeService) Create(ctx *gin.Context, req *dtos.CreateCarTypeRequest) (*dtos.CarTypeResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarTypeService) Update(ctx *gin.Context, id int, req *dtos.UpdateCarTypeRequest) (*dtos.CarTypeResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarTypeService) Delete(ctx *gin.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarTypeService) GetById(ctx *gin.Context, id int) (*dtos.CarTypeResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarTypeService) GetByFilter(ctx *gin.Context, req *dtos.PaginationInputWithFilter) (*dtos.PagedList[dtos.CarTypeResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
