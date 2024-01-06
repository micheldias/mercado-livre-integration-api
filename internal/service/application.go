package service

import (
	"context"
	"mercado-livre-integration/internal/model"
	"mercado-livre-integration/internal/repository"
)

type ApplicationService interface {
	GetAppByID(ctx context.Context, id int) (model.Application, error)
	GetApps(ctx context.Context) ([]model.Application, error)
	SaveApp(ctx context.Context, app model.Application) (model.Application, error)
}

func NewApplicationService(repo repository.ApplicationRepository) ApplicationService {
	return appService{
		repo: repo,
	}
}

type appService struct {
	repo repository.ApplicationRepository
}

func (a appService) GetAppByID(ctx context.Context, id int) (model.Application, error) {
	return a.repo.GetByID(ctx, id)
}

func (a appService) GetApps(ctx context.Context) ([]model.Application, error) {
	return a.repo.Get(ctx)
}

func (a appService) SaveApp(ctx context.Context, app model.Application) (model.Application, error) {
	return a.repo.Save(ctx, app)
}
