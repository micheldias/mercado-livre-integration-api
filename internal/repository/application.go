package repository

import (
	"context"
	"gorm.io/gorm"
	"mercado-livre-integration/internal/model"
)

type ApplicationRepository interface {
	GetByID(ctx context.Context, id int) (model.Application, error)
	Get(ctx context.Context) ([]model.Application, error)
	Save(ctx context.Context, app model.Application) (model.Application, error)
}

func NewApplicationRepository(dbConnection *gorm.DB) ApplicationRepository {
	return appRepository{
		dbConnection: dbConnection,
	}
}

type appRepository struct {
	dbConnection *gorm.DB
}

func (a appRepository) GetByID(ctx context.Context, id int) (model.Application, error) {
	var app model.Application
	if err := a.dbConnection.WithContext(ctx).First(&app, id).Error; err != nil {
		return app, err
	}
	return app, nil
}

func (a appRepository) Get(ctx context.Context) ([]model.Application, error) {
	var apps []model.Application
	if err := a.dbConnection.WithContext(ctx).Find(apps).Error; err != nil {
		return apps, err
	}
	return apps, nil
}

func (a appRepository) Save(ctx context.Context, app model.Application) (model.Application, error) {
	if err := a.dbConnection.WithContext(ctx).Create(app).Error; err != nil {
		return app, err
	}
	return app, nil
}
