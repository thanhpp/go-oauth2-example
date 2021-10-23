package ggoauth2

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

const (
	oauthStateKey = "oauthstate"
)

func GenerateStateCookie(w http.ResponseWriter, expire time.Time) string {
	var (
		b = make([]byte, 16)
	)
	rand.Read(b) // generate a random 16 bytes
	state := base64.URLEncoding.EncodeToString(b)

	// save the oauthstate to cookie
	cookie := http.Cookie{
		Name:    oauthStateKey,
		Value:   state,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)

	return state
}

func OAuthGoogleLogin(w http.ResponseWriter, r *http.Request, oauth2Cfg *oauth2.Config, expire time.Time) {
	// add a state into the cookie
	state := GenerateStateCookie(w, expire)
	redirectLink := oauth2Cfg.AuthCodeURL(state)

	http.Redirect(w, r, redirectLink, http.StatusTemporaryRedirect)
}
