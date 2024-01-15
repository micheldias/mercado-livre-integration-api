package mockks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"mercado-livre-integration/internal/infrastructure/client"
)

type MLClientMock struct {
	mock.Mock
}

func (m MLClientMock) GetUser(ctx context.Context, userID string) (client.User, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(client.User), args.Error(1)
}

func (m MLClientMock) CreateToken(ctx context.Context, clientID, clientSecret, redirectURL, authCode string) (client.AuthTokenResponse, error) {
	args := m.Called(ctx, clientID, clientSecret, redirectURL, authCode)
	return args.Get(0).(client.AuthTokenResponse), args.Error(1)
}

func (m MLClientMock) GetSites(ctx context.Context) (client.Sites, error) {
	args := m.Called(ctx)
	return args.Get(0).(client.Sites), args.Error(1)
}
func (m MLClientMock) GetCategories(ctx context.Context, appId int, siteID string) (client.Categories, error) {
	args := m.Called(ctx, appId, siteID)
	return args.Get(0).(client.Categories), args.Error(1)
}
func (m MLClientMock) CreateProduct(ctx context.Context, product client.ProductRequest) (client.ProductResponse, error) {
	args := m.Called(ctx, product)
	return args.Get(0).(client.ProductResponse), args.Error(1)
}

var _ client.MercadoLivre = (*MLClientMock)(nil)
