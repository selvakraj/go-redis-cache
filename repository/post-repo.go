package repository

import "github.com/selvakraj/go-redis-cache/entity"

type PostRepository interface{
Save(post *entity.Post) (*entity.Post, error)
FindAll() ([]entity.Post, error)
FindById(id string) (*entity.Post, error)
Delete(post *entity.Post) error
}