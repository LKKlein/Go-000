// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"Week04/internal/biz"
	"Week04/internal/data"
	"Week04/internal/service"
)

// Injectors from wire.go:

func InitializeCommentService() *service.CommentSvc {
	commentRepo := data.NewCommentRepo()
	mockCommentRepo := biz.NewMockCommentRepo(commentRepo)
	commentSvc := service.NewCommentSvc(mockCommentRepo)
	return commentSvc
}
