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
	Create(ctx context.Context, applicationID int, authCode string) (model.Token, error)
	GetUrlAuthentication(ctx context.Context, applicationID int) (string, error)
}
type authService struct {
	mercadoLivreClient client.MercadoLivre
	appService         ApplicationService
}

func NewAuthenticationService(mercadoLivreClient client.MercadoLivre, service ApplicationService) AuthenticationService {
	return authService{
		mercadoLivreClient: mercadoLivreClient,
		appService:         service,
	}
}

func (t authService) Create(ctx context.Context, applicationID int, authCode string) (model.Token, error) {
	app, err := t.appService.GetAppByID(ctx, applicationID)
	if err != nil {
		return model.Token{}, err
	}

	accessToken, err := t.mercadoLivreClient.CreateToken(ctx, app.ClientID, app.Secret, app.RedirectURL, authCode)
	if err != nil {
		return model.Token{}, err
	}

	return model.Token{
		AccessToken:     accessToken.AccessToken,
		RefreshToken:    accessToken.RefreshToken,
		ExpireInSeconds: accessToken.ExpiresIn,
	}, err
}

func (t authService) GetUrlAuthentication(ctx context.Context, applicationID int) (string, error) {
	app, err := t.appService.GetAppByID(ctx, applicationID)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/authorization?response_type=code&client_id=%s&redirect_uri=%s&state=%s",
		viper.GetString("MERCADO_LIVRE_AUTH_URL"),
		app.ClientID,
		app.RedirectURL,
		uuid.New().String(),
	), err
}
