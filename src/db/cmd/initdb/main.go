package main

import (
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
	"xorm.io/core"

	"apusic/go-webserver/src/db/schema"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetReportCaller(true)
}

func main() {
	// sample: postgres://postgres:root@localhost:5432/testdb?sslmode=disable
	var dbUrl = os.Getenv("DATABASE_URL")

	if len(dbUrl) == 0 {
		log.Fatal("please set DATABASE_URL env")
	}

	log.Infof("Get DATABASE_URL: %s", dbUrl)

	var err error
	engine, err := xorm.NewEngine("postgres", dbUrl)
	if err != nil {
		log.WithError(err).Fatal("failed to allocate NewPostgreSQL")
	}
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)

	err = engine.Sync2(
		new(schema.Users),
		new(schema.Projects),
	)
	if err != nil {
		log.WithError(err).Fatal("failed to sync schema")
	} else {
		log.Info("init db, create tables completed")
	}
}
