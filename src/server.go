package src

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v1"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"apusic/go-webserver/src/config"
	"apusic/go-webserver/src/db/dao"
	utconfig "apusic/go-webserver/src/utils/config"
	"apusic/go-webserver/src/webserver"
	"apusic/go-webserver/src/webserver/logwrap"
)

type Server struct {
	logger     *logrus.Entry
	dataEngine *dao.DataEngine
	webServer  *webserver.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start(ctx *cli.Context) error {
	defer s.Stop()
	defer func(s *Server) {
		if r := recover(); r != nil {
			debug.PrintStack()
			if s.logger == nil {
				logrus.WithField("panic", r).Fatal("Recovered from panic")
			}
			s.logger.WithField("panic", r).Fatal("Recovered from panic")
		}
	}(s)

	// 读取配置文件
	cfgPath := ctx.String("config")
	var cfg config.Config
	s.checkError(utconfig.ReadConfig(cfgPath, &cfg, nil))

	// 初始化日志模块
	s.logger = logwrap.NewStructureLogger(cfg.WebServer.Logger).
		LogEntryWithFields(logrus.Fields{
			logwrap.AppName:    "go-webserver",
			logwrap.ModuleName: "global",
		})

	// 初始化并启动db/redis/jenkins等等其他组件
	s.dataEngine = dao.NewDataEngine()
	s.webServer = webserver.NewServer()

	s.checkError(s.dataEngine.Start(cfg.DataBase.Url)).
		checkError(s.webServer.Start(&cfg))

	logrus.Info("app starting...")
	s.wait() // 等待系统退出信号
	logrus.Info("app starting222...")
	return nil
}

func (s *Server) checkError(err error) *Server {
	if err != nil {
		const msg = "Start webserver failed."
		if s.logger == nil {
			logrus.WithError(err).Fatal(msg)
		}
		s.logger.WithError(err).Fatal(msg)
	}
	return s
}

func (s *Server) wait() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGKILL,
		syscall.SIGABRT, syscall.SIGBUS, syscall.SIGSEGV)
	for sig := range c {
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM,
			syscall.SIGQUIT, syscall.SIGKILL,
			syscall.SIGABRT, syscall.SIGBUS, syscall.SIGSEGV:
			s.logger.WithField("signal", sig).
				Warn("webserver stop")
			return
		}
	}
}

func (s *Server) Stop() {
	if s.webServer != nil {
		s.webServer.Stop()
	}
}
