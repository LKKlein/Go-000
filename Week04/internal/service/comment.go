package service

import (
	"context"
	"log"

	"Week04/api"
	"Week04/internal/biz"
)

type CommentSvc struct {
	c *biz.MockCommentRepo
	api.UnimplementedCommentServer
}

func NewCommentSvc(mc *biz.MockCommentRepo) *CommentSvc {
	log.Println("new comment service")
	return &CommentSvc{c: mc}
}

func (s *CommentSvc) WriteComment(ctx context.Context, req *api.CommentRequest) (*api.CommentReply, error) {
	comment := &biz.Comment{
		ID:      req.Id,
		Content: req.Content,
	}

	id := s.c.Save(comment)

	return &api.CommentReply{Id: id}, nil
}
