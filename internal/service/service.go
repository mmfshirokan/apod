package service

import (
	"context"

	"github.com/mmfshirokan/apod/internal/model"
)

type InfoRepository interface {
	Add(ctx context.Context, ii model.ImageInfo) error
	Get(ctx context.Context, date string) (model.ImageInfo, error)
	GetAll(ctx context.Context) ([]model.ImageInfo, error)
}

type InfoService struct {
	repo InfoRepository
}

func NewInfo(repo InfoRepository) *InfoService {
	return &InfoService{repo: repo}
}

func (ns InfoService) Add(ctx context.Context, ii model.ImageInfo) error {
	return ns.repo.Add(ctx, ii)
}

func (ns InfoService) Get(ctx context.Context, date string) (model.ImageInfo, error) {
	return ns.repo.Get(ctx, date)
}

func (ns InfoService) GetAll(ctx context.Context) ([]model.ImageInfo, error) {
	return ns.repo.GetAll(ctx)
}
