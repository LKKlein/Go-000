package data

import (
	"Week04/internal/biz"
	"fmt"
	"log"
)

var _ biz.CommentRepo = (biz.CommentRepo)(nil)

type commentRepo struct {
	id   string
	text string
}

func NewCommentRepo() biz.CommentRepo {
	log.Println("new comment data")
	return new(commentRepo)
}

func (m *commentRepo) SaveComment(comment *biz.Comment) (id int64) {
	fmt.Printf("comment content: %s, id: %d saved!\n", comment.Content, comment.ID)
	return comment.ID
}

func (m *commentRepo) GetCommentByID(id int64) (*biz.Comment, error) {
	return &biz.Comment{
		ID:      id,
		Content: "we are family!",
	}, nil
}
