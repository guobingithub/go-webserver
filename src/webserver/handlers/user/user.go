package user

import (
	"apusic/go-webserver/src/db/dao"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	dbEngine *dao.DataEngine
}

func NewUserService(dbEngine *dao.DataEngine) *UserService {
	return &UserService{
		dbEngine: dbEngine,
	}
}

func (s *UserService) Info(ctx *gin.Context) {
	userDao := new(dao.User)
	user := userDao.GetUserInfoById()

	//
	user = user
	data := GetUserInfoRsp{}
	data = data

	return
}

func (s *UserService) List(ctx *gin.Context) {

}
