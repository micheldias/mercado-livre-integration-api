package service

import (
	"mercado-livre-integration/internal/infrastructure/client"
	"mercado-livre-integration/internal/model"
)

type TokenService interface {
	Create(authCode string) (model.Token, error)
}
type token struct {
	mercadoLivreClient client.MercadoLivre
}

func NewToken(mercadoLivreClient client.MercadoLivre) TokenService {
	tokenService := token{
		mercadoLivreClient: mercadoLivreClient,
	}

	return tokenService
}

func (t token) Create(authCode string) (model.Token, error) {
	accessToken, err := t.mercadoLivreClient.CreateToken(authCode)
	if err != nil {
		return model.Token{}, err
	}

	return model.Token{
		AccessToken:     accessToken.AccessToken,
		RefreshToken:    accessToken.RefreshToken,
		ExpireInSeconds: accessToken.ExpiresIn,
	}, err
}
