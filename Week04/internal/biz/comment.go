package biz

import "log"

type Comment struct {
	ID      int64
	Content string
}

type CommentRepo interface {
	GetCommentByID(id int64) (*Comment, error)
	SaveComment(c *Comment) int64
}

type MockCommentRepo struct {
	repo CommentRepo
}

// 依赖注入
func NewMockCommentRepo(repo CommentRepo) *MockCommentRepo {
	log.Println("new comment biz")
	return &MockCommentRepo{repo: repo}
}

// 具体实现逻辑
func (mc *MockCommentRepo) Save(o *Comment) int64 {
	id := mc.repo.SaveComment(o)
	return id
}
