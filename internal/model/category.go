package model

import "mercado-livre-integration/internal/infrastructure/client"

type Categories []Category

func (c Categories) From(categories client.Categories) []Category {
	domainCategories := make(Categories, len(categories))
	for i, cat := range categories {
		domainCategories[i] = Category{
			ID:   cat.ID,
			Name: cat.Name,
		}
	}
	return domainCategories
}

type Category struct {
	ID   string
	Name string
}
