package service

import (
	"io"
)

type ImageRepository interface {
	Add(image io.Reader, name string) error
}

func NewImage(repo ImageRepository) *ImageService {
	return &ImageService{repo: repo}
}

type ImageService struct {
	repo ImageRepository
}

func (ms *ImageService) Add(image io.Reader, name string) error {
	return ms.repo.Add(image, name)
}
