package service

import (
	"mercado-livre-integration/internal/infrastructure/client"
	"mercado-livre-integration/internal/model"
)

type CategoryService interface {
	GetCategories(siteID string) (model.Categories, error)
}

type category struct {
	mercadoLivreClient client.MercadoLivre
}

func NewCategory(mercadoLivreClient client.MercadoLivre) CategoryService {
	return category{
		mercadoLivreClient: mercadoLivreClient,
	}

}

func (c category) GetCategories(siteID string) (categories model.Categories, err error) {
	mlCategories, err := c.mercadoLivreClient.GetCategories(siteID)
	if err != nil {
		return categories, err
	}
	return categories.From(mlCategories), nil
}
