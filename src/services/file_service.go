package services

import (
	"golang-web-api/api/dtos"
	"golang-web-api/config"
	"golang-web-api/data/db"
	"golang-web-api/data/models"
	"golang-web-api/pkg/logging"

	"github.com/gin-gonic/gin"
)

type FileService struct {
	base *BaseService[models.File, dtos.CreateFileRequest, dtos.UpdateFileRequest, dtos.FileResponse]
}

func NewFileService(cfg *config.Config) *FileService {
	return &FileService{
		base: &BaseService[models.File, dtos.CreateFileRequest, dtos.UpdateFileRequest, dtos.FileResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// Create
func (s *FileService) Create(ctx *gin.Context, req *dtos.CreateFileRequest) (*dtos.FileResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *FileService) Update(ctx *gin.Context, id int, req *dtos.UpdateFileRequest) (*dtos.FileResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *FileService) Delete(ctx *gin.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *FileService) GetById(ctx *gin.Context, id int) (*dtos.FileResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *FileService) GetByFilter(ctx *gin.Context, req *dtos.PaginationInputWithFilter) (*dtos.PagedList[dtos.FileResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
