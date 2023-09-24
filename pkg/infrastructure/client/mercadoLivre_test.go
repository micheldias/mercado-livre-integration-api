package client

import (
	"errors"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
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
)

func TestAuthClient(t *testing.T) {

	t.Run("should call create auth successfully", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPost, "http://mocktest/oauth/token",
			httpmock.NewStringResponder(http.StatusOK, authCodeResponse))

		client := NewMercadoLivre("client_id", "client_secret", "http://localhost", "http://mocktest")

		response, err := client.CreateToken("TG-648f8999952b710001817e36-146322322")

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
		client := NewMercadoLivre("client_id", "client_secret", "http://localhost", "http://mocktest")

		_, err := client.CreateToken("TG-648f8999952b710001817e36-146322322")

		assert.EqualError(t, err, "status code: 400")

	})

	t.Run("should return an error when unable to make the request", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(http.MethodPost, "http://mocktest/oauth/token",
			httpmock.NewErrorResponder(errors.New("bla")))
		client := NewMercadoLivre("client_id", "client_secret", "http://localhost", "http://mocktest")

		_, err := client.CreateToken("auth-code")
		assert.EqualError(t, err, `failed to execute http request. Error: Post "http://mocktest/oauth/token": bla`)
	})
}
