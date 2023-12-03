package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

const (
	authCodeResponse = `{
    "access_token": "APP_USR-761615972200605-061812-79870010a2b5027f9bdd7500a00fe6f8-146322322",
    "token_type": "Bearer",
    "expires_in": 21600,
    "scope": "offline_access read write",
    "user_id": 146322322,
    "refresh_token": "TG-648f2c088bff590001df3ce5-146322322"
	}`

	authCodeResponseError = `{
    "cause": [],
    "error": "invalid_grant",
    "error_description": "Error validating grant. Your authorization code or refresh token may be expired or it was already used",
    "status": 400
}`
	getUserResponse = `{
    "id": 1496809856,
    "nickname": "TESTUSER2074800394",
    "country_id": "AR",
    "address": {
        "city": "Palermo",
        "state": "AR-C"
    },
    "user_type": "normal",
    "site_id": "MLA",
    "permalink": "http://perfil.mercadolibre.com.ar/TESTUSER2074800394",
    "seller_reputation": {
        "level_id": null,
        "power_seller_status": null,
        "transactions": {
            "period": "historic",
            "total": 0
        }
    },
    "status": {
        "site_status": "active"
    }
}`

	getSitesResponse = `
[
    {
        "default_currency_id": "CRC",
        "id": "MCR",
        "name": "Costa Rica"
    },
    {
        "default_currency_id": "BRL",
        "id": "MLB",
        "name": "Brasil"
    },
    {
        "default_currency_id": "HNL",
        "id": "MHN",
        "name": "Honduras"
    }
]`

	getCategoriesResponse = `[
    {
        "id": "MLA5725",
        "name": "Accesorios para Veículos"
    },
    {
        "id": "MLA1512",
        "name": "Agro"
    },
    {
        "id": "MLA1403",
        "name": "Alimentos y Bebidas"
    }
]`
)

func TestAuthClient(t *testing.T) {

	t.Run("should call create auth successfully", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPost, "http://mocktest/oauth/token",
			httpmock.NewStringResponder(http.StatusOK, authCodeResponse))

		client := NewMercadoLivre("http://mocktest", time.Second)

		response, err := client.CreateToken(context.Background(), "client_id", "client_secret", "http://localhost", "TG-648f8999952b710001817e36-146322322")

		assert.Nil(t, err)
		assert.Equal(t, "APP_USR-761615972200605-061812-79870010a2b5027f9bdd7500a00fe6f8-146322322", response.AccessToken)
		assert.Equal(t, "TG-648f2c088bff590001df3ce5-146322322", response.RefreshToken)
		assert.Equal(t, "Bearer", response.TokenType)
		assert.Equal(t, 21600, response.ExpiresIn)
		assert.Equal(t, "offline_access read write", response.Scope)
		assert.Equal(t, 146322322, response.UserId)
	})

	t.Run("should call create auth with error", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder(http.MethodPost, "http://mocktest/oauth/token",
			httpmock.NewStringResponder(http.StatusBadRequest, authCodeResponseError))
		client := NewMercadoLivre("http://mocktest", time.Second)

		_, err := client.CreateToken(context.Background(), "client_id", "client_secret", "http://localhost", "TG-648f8999952b710001817e36-146322322")

		assert.EqualError(t, err, "status code: 400")

	})

	t.Run("should return an error when unable to make the request", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPost, "http://mocktest/oauth/token",
			httpmock.NewErrorResponder(errors.New("bla")))
		client := NewMercadoLivre("http://mocktest", time.Second)

		_, err := client.CreateToken(context.Background(), "client_id", "client_secret", "http://localhost", "auth-code")
		assert.EqualError(t, err, `failed to execute http request. Error: Post "http://mocktest/oauth/token": bla`)
	})
}

func TestGetUser(t *testing.T) {

	t.Run("should call get users successfully", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodGet, fmt.Sprintf("http://mocktest/users/%s", "12344"),
			httpmock.NewStringResponder(http.StatusOK, getUserResponse))

		client := NewMercadoLivre("http://mocktest", time.Second)

		user, err := client.GetUser(context.Background(), "12344")

		assert.Nil(t, err)
		assert.Equal(t, 1496809856, user.ID)
		assert.Equal(t, "active", user.Status.SiteStatus)
	})
}

func TestGetSites(t *testing.T) {

	t.Run("should call get sites successfully", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodGet, "http://mocktest/sites",
			httpmock.NewStringResponder(http.StatusOK, getSitesResponse))

		client := NewMercadoLivre("http://mocktest", time.Second)

		sites, err := client.GetSites(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, 3, len(sites))
		assert.Equal(t, "MCR", sites[0].ID)
		assert.Equal(t, "Costa Rica", sites[0].Name)
		assert.Equal(t, "CRC", sites[0].DefaultCurrencyID)
	})
}

func TestGetCategories(t *testing.T) {

	t.Run("should call get categories successfully", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodGet, "http://mocktest/sites/MLB/categories",
			httpmock.NewStringResponder(http.StatusOK, getCategoriesResponse))

		client := NewMercadoLivre("http://mocktest", time.Second)

		categories, err := client.GetCategories(context.Background(), "MLB")

		assert.Nil(t, err)
		assert.Equal(t, 3, len(categories))
		assert.Equal(t, "MLA5725", categories[0].ID)
		assert.Equal(t, "Accesorios para Veículos", categories[0].Name)
	})
}

//func TestName(t *testing.T) {
//
//	c := new(mockks.MLClientMock)
//	authToken := &client.AuthTokenResponse{
//		AccessToken:  "access_token",
//		RefreshToken: "refresh_token",
//		ExpiresIn:    10,
//		Scope:        "scope",
//		TokenType:    "token_type",
//	}
//
//	c.On("CreateToken", "token").Return(authToken, nil)
//
//	s := NewToken(c, time.Second*1)
//
//	tokenResponse, err := s.CreateToken("token")
//	//time.Sleep(time.Second * 5)
//	assert.Equal(t, "access_token", tokenResponse.AccessToken)
//	assert.Nil(t, err)
//	assert.True(t, c.AssertExpectations(t))
//
//}
