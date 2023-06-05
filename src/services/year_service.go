package services

import (
	"golang-web-api/api/dtos"
	"golang-web-api/config"
	"golang-web-api/data/db"
	"golang-web-api/data/models"
	"golang-web-api/pkg/logging"

	"github.com/gin-gonic/gin"
)

type PersianYearService struct {
	base *BaseService[models.PersianYear, dtos.CreatePersianYearRequest, dtos.UpdatePersianYearRequest, dtos.PersianYearResponse]
}

func NewPersianYearService(cfg *config.Config) *PersianYearService {
	return &PersianYearService{
		base: &BaseService[models.PersianYear, dtos.CreatePersianYearRequest, dtos.UpdatePersianYearRequest, dtos.PersianYearResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// Create
func (s *PersianYearService) Create(ctx *gin.Context, req *dtos.CreatePersianYearRequest) (*dtos.PersianYearResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *PersianYearService) Update(ctx *gin.Context, id int, req *dtos.UpdatePersianYearRequest) (*dtos.PersianYearResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *PersianYearService) Delete(ctx *gin.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *PersianYearService) GetById(ctx *gin.Context, id int) (*dtos.PersianYearResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *PersianYearService) GetByFilter(ctx *gin.Context, req *dtos.PaginationInputWithFilter) (*dtos.PagedList[dtos.PersianYearResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
