package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"

	"apusic/go-webserver/src"
)

func main() {
	app := cli.NewApp()

	app.Name = "WebServer"
	app.Version = "1.0.0"
	app.Description = "Golang web服务器示例"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config,c",
			Usage: "配置文件",
			Value: "./config/config.yaml",
		},
	}
	server := src.NewServer()
	app.Action = server.Start

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
