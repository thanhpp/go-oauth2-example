package ggoauth2

import (
	"encoding/json"
	"io/ioutil"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	Web Web `json:"web"`
}

type Web struct {
	ClientID                string   `json:"client_id"`
	ProjectID               string   `json:"project_id"`
	AuthURI                 string   `json:"auth_uri"`
	TokenURI                string   `json:"token_uri"`
	AuthProviderX509CertURL string   `json:"auth_provider_x509_cert_url"`
	ClientSecret            string   `json:"client_secret"`
	RedirectUris            []string `json:"redirect_uris"`
}

func NewConfigFromFile(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func NewOAuth2ConfigFromFile(path string) (*oauth2.Config, error) {
	cfg, err := NewConfigFromFile(path)
	if err != nil {
		return nil, err
	}

	return &oauth2.Config{
		ClientID:     cfg.Web.ClientID,
		ClientSecret: cfg.Web.ClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  cfg.Web.RedirectUris[0],
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/user.birthday.read",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}, nil
}
