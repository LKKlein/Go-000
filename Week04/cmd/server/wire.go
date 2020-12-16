// +build wireinject

package main

import (
	"Week04/internal/biz"
	"Week04/internal/data"
	"Week04/internal/service"

	"github.com/google/wire"
)

func InitializeCommentService() *service.CommentSvc {
	wire.Build(service.NewCommentSvc, biz.NewMockCommentRepo, data.NewCommentRepo)
	return &service.CommentSvc{}
}
