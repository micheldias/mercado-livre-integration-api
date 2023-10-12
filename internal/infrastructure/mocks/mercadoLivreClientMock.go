package mockks

import (
	"github.com/stretchr/testify/mock"
	"mercado-livre-integration/internal/infrastructure/client"
)

type MLClientMock struct {
	mock.Mock
}

func (m MLClientMock) GetUser(userID string) (client.User, error) {
	args := m.Called(userID)
	return args.Get(0).(client.User), args.Error(1)
}

func (m MLClientMock) CreateToken(authCode string) (client.AuthTokenResponse, error) {
	args := m.Called(authCode)
	return args.Get(0).(client.AuthTokenResponse), args.Error(1)
}

func (m MLClientMock) RefreshToken(refreshToken string) (*client.AuthTokenResponse, error) {
	args := m.Called(refreshToken)
	return args.Get(0).(*client.AuthTokenResponse), args.Error(1)
}

func (m MLClientMock) GetSites() (client.Sites, error) {
	args := m.Called()
	return args.Get(0).(client.Sites), args.Error(1)
}
func (m MLClientMock) GetCategories(siteID string) (client.Categories, error) {
	args := m.Called(siteID)
	return args.Get(0).(client.Categories), args.Error(1)
}
func (m MLClientMock) CreateProduct(product client.ProductRequest) (client.ProductResponse, error) {
	args := m.Called(product)
	return args.Get(0).(client.ProductResponse), args.Error(1)
}

var _ client.MercadoLivre = (*MLClientMock)(nil)
