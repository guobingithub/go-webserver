package project

import (
	"apusic/go-webserver/src/db/dao"
	"github.com/gin-gonic/gin"
)

type ProjectService struct {
	dbEngine *dao.DataEngine
}

func NewProjectService(dbEngine *dao.DataEngine) *ProjectService {
	return &ProjectService{
		dbEngine: dbEngine,
	}
}

func (s *ProjectService) Info(ctx *gin.Context) {

}

func (s *ProjectService) List(ctx *gin.Context) {

}
