package handler

import (
	"context"
	"github.com/unliar/happysooner-proto/file"
)

type FileHandler struct {
}

func (f *FileHandler) PostImage(ctx context.Context, req *file.PostFileRequest, res *file.PostFileResponse) error {
	return nil
}
