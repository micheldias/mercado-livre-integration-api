package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"mercado-livre-integration/internal/infrastructure/client"
	"mercado-livre-integration/internal/model"
)

type AuthenticationService interface {
	Create(ctx context.Context, authCode string) (model.Token, error)
	GetUrlAuthentication(ctx context.Context) string
}
type authService struct {
	mercadoLivreClient client.MercadoLivre
}

func NewAuthenticationService(mercadoLivreClient client.MercadoLivre) AuthenticationService {
	return authService{
		mercadoLivreClient: mercadoLivreClient,
	}
}

func (t authService) Create(ctx context.Context, authCode string) (model.Token, error) {
	accessToken, err := t.mercadoLivreClient.CreateToken(ctx, authCode)
	if err != nil {
		return model.Token{}, err
	}

	return model.Token{
		AccessToken:     accessToken.AccessToken,
		RefreshToken:    accessToken.RefreshToken,
		ExpireInSeconds: accessToken.ExpiresIn,
	}, err
}

func (t authService) GetUrlAuthentication(ctx context.Context) string {
	return fmt.Sprintf("%s/authorization?response_type=code&client_id=%s&redirect_uri=%s&state=%s",
		viper.GetString("MERCADO_LIVRE_AUTH_URL"),
		viper.GetString("MERCADO_LIVRE_CLIENT_ID"),
		viper.GetString("MERCADO_LIVRE_REDIRECT_URL"),
		uuid.New().String(),
	)
}
