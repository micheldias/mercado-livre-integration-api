package client

type AuthTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	UserId       int    `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
}

type AuthTokenErrorResponse struct {
	Cause            []interface{} `json:"cause"`
	Error            string        `json:"error"`
	ErrorDescription string        `json:"error_description"`
	Status           int           `json:"status"`
}
