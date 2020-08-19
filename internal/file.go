package internal

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/v3/logger"
	"github.com/unliar/happysooner-proto/file"
	"go-rpc-happysooner-file/internal/handler"
	"go-rpc-happysooner-file/internal/utils"
	"time"
)

func Init() {
	srv := micro.NewService(
		micro.Name("happysooner-file-rpc"),
		micro.Version("v1.0.0"),
		micro.RegisterInterval(15*time.Second),
		micro.RegisterTTL(time.Second*60),
		micro.BeforeStart(func() error {
			fmt.Println("文件服务开始启动")
			return nil
		}),
		micro.WrapHandler(utils.LogWrapper),
	)
	srv.Init()
	_ = file.RegisterFileSVHandler(srv.Server(), new(handler.FileHandler))

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
