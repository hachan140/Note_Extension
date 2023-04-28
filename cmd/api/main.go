package main

import (
	"fmt"
	"go.uber.org/zap"
	"note_extension/pkg/config"
)

func init() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}

func main() {
	zap.S().Infow("start process", "name", "api")

	mySQLConf, err := config.NewMySQLConfig()
	if err != nil {
		zap.S().Fatalw("error when call config.NewMySQLConfig", "error", err)
	}

	fmt.Println(mySQLConf.ConnectionString())
}
