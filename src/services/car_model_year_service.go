package services

import (
	"golang-web-api/api/dtos"
	"golang-web-api/config"
	"golang-web-api/data/db"
	"golang-web-api/data/models"
	"golang-web-api/pkg/logging"

	"github.com/gin-gonic/gin"
)

type CarModelYearService struct {
	base *BaseService[models.CarModelYear, dtos.CreateCarModelYearRequest, dtos.UpdateCarModelYearRequest, dtos.CarModelYearResponse]
}

func NewCarModelYearService(cfg *config.Config) *CarModelYearService {
	return &CarModelYearService{
		base: &BaseService[models.CarModelYear, dtos.CreateCarModelYearRequest, dtos.UpdateCarModelYearRequest, dtos.CarModelYearResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "PersianYear"},
			},
		},
	}
}

// Create
func (s *CarModelYearService) Create(ctx *gin.Context, req *dtos.CreateCarModelYearRequest) (*dtos.CarModelYearResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelYearService) Update(ctx *gin.Context, id int, req *dtos.UpdateCarModelYearRequest) (*dtos.CarModelYearResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelYearService) Delete(ctx *gin.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelYearService) GetById(ctx *gin.Context, id int) (*dtos.CarModelYearResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarModelYearService) GetByFilter(ctx *gin.Context, req *dtos.PaginationInputWithFilter) (*dtos.PagedList[dtos.CarModelYearResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
