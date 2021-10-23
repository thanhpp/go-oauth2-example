package ggoauth2

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

func OAuthGoogleCallback(w http.ResponseWriter, r *http.Request) {
	// readh oauthState from Cookie
	oauthState, err := r.Cookie(oauthStateKey)
	if err != nil {
		panic(err)
	}

	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

}

const GoogleOAuthAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func GetUserDataFromGoogle(ctx context.Context, code string, oauthCfg *oauth2.Config) ([]byte, error) {
	token, err := oauthCfg.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("Exchange code error :%v", err)
	}

	resp, err := http.Get(GoogleOAuthAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %v", err)
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read resp body %v", err)
	}

	return contents, nil
}
