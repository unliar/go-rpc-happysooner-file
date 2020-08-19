package utils

import (
	"context"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/v3/logger"
)

func LogWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		logger.Infof("[wrapper] server request: %v", req.Method())
		err := fn(ctx, req, rsp)
		return err
	}
}
