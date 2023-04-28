package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"note_extension/ent"
	"note_extension/pkg/config"
)

func init() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}

func main() {
	mySQLConfig, err := config.NewMySQLConfig()
	if err != nil {
		zap.S().Errorw("error when call config.NewMySQLConfig", "error", err)
		return
	}
	client, err := ent.Open("mysql", mySQLConfig.ConnectionString())
	if err != nil {
		zap.S().Errorw("error when call ent.Open", "error", err, "driver", "mysql", "connection_string", mySQLConfig.ConnectionString())
		return
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		zap.S().Errorw("error when call client.Schema.Create", "error", err)
		return
	}
	zap.S().Info("success")
}
