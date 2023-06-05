package services

import (
	"golang-web-api/api/dtos"
	"golang-web-api/config"
	"golang-web-api/constants"
	"golang-web-api/data/db"
	"golang-web-api/data/models"
	"golang-web-api/pkg/logging"

	"github.com/gin-gonic/gin"
)

type CarModelCommentService struct {
	base *BaseService[models.CarModelComment, dtos.CreateCarModelCommentRequest, dtos.UpdateCarModelCommentRequest, dtos.CarModelCommentResponse]
}

func NewCarModelCommentService(cfg *config.Config) *CarModelCommentService {
	return &CarModelCommentService{
		base: &BaseService[models.CarModelComment, dtos.CreateCarModelCommentRequest, dtos.UpdateCarModelCommentRequest, dtos.CarModelCommentResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "User"},
			},
		},
	}
}

// Create
func (s *CarModelCommentService) Create(ctx *gin.Context, req *dtos.CreateCarModelCommentRequest) (*dtos.CarModelCommentResponse, error) {
	req.UserId = int(ctx.Value(constants.UserIdKey).(float64))
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelCommentService) Update(ctx *gin.Context, id int, req *dtos.UpdateCarModelCommentRequest) (*dtos.CarModelCommentResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelCommentService) Delete(ctx *gin.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelCommentService) GetById(ctx *gin.Context, id int) (*dtos.CarModelCommentResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarModelCommentService) GetByFilter(ctx *gin.Context, req *dtos.PaginationInputWithFilter) (*dtos.PagedList[dtos.CarModelCommentResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
