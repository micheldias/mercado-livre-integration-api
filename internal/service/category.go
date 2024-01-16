package service

import (
	"context"
	"mercado-livre-integration/internal/infrastructure/client"
	"mercado-livre-integration/internal/model"
)

type CategoryService interface {
	GetCategories(ctx context.Context, appId int, siteID string) (model.Categories, error)
}

type category struct {
	mercadoLivreClient client.MercadoLivre
	appService         ApplicationService
}

func NewCategory(mercadoLivreClient client.MercadoLivre) CategoryService {
	return category{
		mercadoLivreClient: mercadoLivreClient,
	}

}

func (c category) GetCategories(ctx context.Context, appId int, siteID string) (categories model.Categories, err error) {
	app, err := c.appService.GetAppByID(ctx, appId)
	if err != nil {
		return model.Categories{}, err
	}

	mlCategories, err := c.mercadoLivreClient.GetCategories(ctx, app.ClientID, siteID)
	if err != nil {
		return categories, err
	}
	return categories.From(mlCategories), nil
}
