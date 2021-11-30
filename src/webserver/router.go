package webserver

import (
	"apusic/go-webserver/src/config"
	"apusic/go-webserver/src/webserver/handlers/project"
	"apusic/go-webserver/src/webserver/handlers/user"
)

func (s *Server) registerHandlers(cfg *config.Config) {
	// middleware service
	// 统计、日志、数据库事务等中间件

	appV1 := s.app.Group("/api/v1")

	// user
	userRouter := appV1.Group("/user")
	{
		userSvc := user.NewUserService(s.dbEngine)
		userRouter.POST("/info", userSvc.Info)
		userRouter.POST("/list", userSvc.List)
	}

	// project
	projectRouter := appV1.Group("/project")
	{
		projectSvc := project.NewProjectService(s.dbEngine)
		projectRouter.POST("/info", projectSvc.Info)
		projectRouter.POST("/list", projectSvc.List)
	}
}
