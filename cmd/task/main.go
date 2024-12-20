package main

import (
	"context"
	"flag"

	"helloadmin/cmd/task/wire"
	"helloadmin/pkg/config"
	"helloadmin/pkg/log"
)

func main() {
	confFileName := "local.yml"
	envConf := flag.String("conf", "config/"+confFileName, "config path, eg: -conf ./config/"+confFileName)
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger := log.NewLog(conf)
	logger.Info("start task")
	app, cleanup, err := wire.NewWire(conf, logger)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}
