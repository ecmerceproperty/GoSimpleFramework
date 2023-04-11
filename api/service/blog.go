package service

import (
	"blog/api/repository"
	"blog/models"
)

type PostService struct {
	repository repository.PostRepository
}

func NewPostService(r repository.PostRepository) PostService {
	return PostService{
		repository: r,
	}
}

func (p PostService) FindAll(post models.Post, keyword string) (*[]models.Post, int64, error) {
	return p.repository.FindAll(post, keyword)
}

func (p PostService) Save(post models.Post) error {
	return p.repository.Save(post)
}

func (p PostService) Update(post models.Post) error {
	return p.repository.Update(post)
}

func (p PostService) Delete(id int64) error {
	var post models.Post
	post.ID = id
	return p.repository.Delete(post)
}

func (p PostService) Find(post models.Post) (models.Post, error) {
	return p.repository.Find(post)
}
