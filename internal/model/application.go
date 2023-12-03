package model

type Application struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	RedirectURL string `json:"redirect_url"`
}

func (a Application) TableName() string {
	return "api.applications"
}
