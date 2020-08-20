package db

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v3/logger"
	"go-rpc-happysooner-file/internal/service/ent"
	"go-rpc-happysooner-file/internal/service/ent/migrate"
	"os"
)
import _ "github.com/go-sql-driver/mysql"

var Client *ent.Client

func InitDBClient() {
	var err error
	x := os.Getenv("MYSQLLOCAL")
	dsn := fmt.Sprintf("%s/entc?charset=utf8mb4&parseTime=true&loc=Local", x)
	Client, err = ent.Open("mysql", dsn)
	if err != nil {
		logger.Fatal(err)
	}
	if err = Client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		logger.Fatalf("create schema error %v", err)
	}
	logger.Info("初始化DB")
	u, err := Client.File.Create().SetUID(1).SetURL("https://baidu.com").Save(context.Background())
	if err != nil {
		logger.Fatalf("create file error %v", err)
		return
	}
	fmt.Println(u)
}
