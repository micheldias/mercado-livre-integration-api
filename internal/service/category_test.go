package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mercado-livre-integration/internal/infrastructure/client"
	mockks "mercado-livre-integration/internal/infrastructure/mocks"
	"testing"
)

func TestGetCategories(t *testing.T) {
	t.Run("get categories successfully", func(t *testing.T) {
		mockClient := mockks.MLClientMock{}
		mockClient.On("GetCategories", mock.Anything).Return(client.Categories{
			{ID: "1", Name: "bla"},
			{ID: "2", Name: "bla_bla"},
		}, nil)

		service := NewCategory(mockClient)
		categories, err := service.GetCategories("12343")
		assert.NoError(t, err)
		assert.Equal(t, 2, len(categories))
		assert.Equal(t, "1", categories[0].ID)
		assert.Equal(t, "bla", categories[0].Name)
		assert.Equal(t, "2", categories[1].ID)
		assert.Equal(t, "bla_bla", categories[1].Name)

	})
	t.Run("get categories successfully", func(t *testing.T) {
		mockClient := mockks.MLClientMock{}
		mockClient.On("GetCategories", mock.Anything).Return(client.Categories{}, nil)

		service := NewCategory(mockClient)
		categories, err := service.GetCategories("12343")
		assert.NoError(t, err)
		assert.Equal(t, 0, len(categories))
		assert.Nil(t, categories)

	})

}
