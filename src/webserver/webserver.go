package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"apusic/go-webserver/src/config"
	"apusic/go-webserver/src/db/dao"
	"apusic/go-webserver/src/webserver/logwrap"
)

type Server struct {
	app      *gin.Engine
	dbEngine *dao.DataEngine
	logger   *logrus.Entry
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start(cfg *config.Config) (err error) {
	// 初始化日志库
	logger := logwrap.NewStructureLogger(cfg.WebServer.Logger)
	s.logger = logger.LogEntryWithFields(logrus.Fields{
		logwrap.ModuleName: "webserver",
	})

	// 初始化 gin app，并设置模式日志库
	gin.SetMode(gin.DebugMode)
	s.app = gin.Default()

	// gin路由初始化
	s.registerHandlers(cfg)

	go func() {
		s.checkError(s.app.Run(cfg.WebServer.Address), "run gin app failed")
	}()

	return
}

func (s *Server) checkError(err error, msg ...string) {
	if err != nil {
		if s.logger == nil {
			logrus.WithError(err).Fatal(msg)
		}
		s.logger.WithError(err).Fatal(msg)
	}

	return
}

func (s *Server) Stop() (err error) {
	return err
}
