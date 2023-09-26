package service

import (
	"mercado-livre-integration/pkg/domain/token/model"
	"mercado-livre-integration/pkg/infrastructure/client"
	"time"
)

type TokenService interface {
	Create(authCode string) (model.Token, error)
}
type token struct {
	mercadoLivreClient client.MercadoLivre
	executeTimes       time.Duration
}

func NewToken(mercadoLivreClient client.MercadoLivre) TokenService {
	tokenService := token{
		mercadoLivreClient: mercadoLivreClient,
	}

	return tokenService
}

func (t token) Create(authCode string) (model.Token, error) {
	token, err := t.mercadoLivreClient.CreateToken(authCode)
	if err != nil {
		return model.Token{}, err
	}

	domainToken := model.Token{
		AccessToken:     token.AccessToken,
		RefreshToken:    token.RefreshToken,
		ExpireInSeconds: token.ExpiresIn,
	}

	return domainToken, err
}
