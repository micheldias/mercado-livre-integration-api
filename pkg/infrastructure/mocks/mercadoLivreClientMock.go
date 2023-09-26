package mockks

import (
	"github.com/stretchr/testify/mock"
	"mercado-livre-integration/pkg/infrastructure/client"
)

type MLClientMock struct {
	mock.Mock
}

func (m MLClientMock) GetUser(userID string) (client.User, error) {
	args := m.Called(userID)
	return args.Get(0).(client.User), args.Error(1)
}

func (m MLClientMock) CreateToken(authCode string) (*client.AuthTokenResponse, error) {
	args := m.Called(authCode)
	return args.Get(0).(*client.AuthTokenResponse), args.Error(1)
}

func (m MLClientMock) RefreshToken(refreshToken string) (*client.AuthTokenResponse, error) {
	args := m.Called(refreshToken)
	return args.Get(0).(*client.AuthTokenResponse), args.Error(1)
}
