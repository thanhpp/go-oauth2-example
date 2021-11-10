package httpserver

import (
	"encoding/base64"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/go-oauth2-example/infrastructure/apdaters/ggoauth2"
	"golang.org/x/oauth2"
)

const (
	oauthStateKey = "oauthstate"
)

type GoogleOAuthCtrl struct {
	oauth2Cfg           *oauth2.Config
	oauthExpireDuration time.Duration
}

func NewGoogleOAuthCtrl(oauth2Cfg *oauth2.Config) GoogleOAuthCtrl {
	return GoogleOAuthCtrl{
		oauth2Cfg:           oauth2Cfg,
		oauthExpireDuration: time.Hour * 24 * 365,
	}
}

func (ctrl GoogleOAuthCtrl) randomState() string {
	b := make([]byte, 16)
	seeded := rand.New(rand.NewSource(time.Now().Unix()))
	seeded.Read(b)

	return base64.URLEncoding.EncodeToString(b)
}

func (ctrl *GoogleOAuthCtrl) LoginHandler(c *gin.Context) {
	oauthState := ctrl.randomState()
	log.Println("Login - New oauth state", oauthState)

	// set oauth2 state into the cookie
	c.SetCookie(
		oauthStateKey,
		oauthState,
		int(time.Now().Add(ctrl.oauthExpireDuration).Unix()),
		"",
		"",
		false,
		false,
	)

	// build redirect URL
	redirectURL := ctrl.oauth2Cfg.AuthCodeURL(oauthState)
	log.Println("Login - Redirect URL", redirectURL)

	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}

func (ctrl *GoogleOAuthCtrl) CallbackHandler(c *gin.Context) {
	// get oauth state from cookie
	oauthState, err := c.Cookie(oauthStateKey)
	if err != nil {
		resp := new(ErrorDTO)
		resp.SetCodeMsg(http.StatusNotAcceptable, "state not found")
		c.AbortWithStatusJSON(http.StatusNotAcceptable, resp)
		return
	}
	log.Println("Callback - state from cookie", oauthState)

	// get the state from form (URL) and compare with the state in the cookie
	formState := c.Request.FormValue("state")
	log.Println("Callback - state from form", c.Request.FormValue("state"))
	if formState != oauthState {
		resp := new(ErrorDTO)
		resp.SetCode(http.StatusNotAcceptable)
		resp.SetMessage("invalid oauth state")
		c.AbortWithStatusJSON(http.StatusNotAcceptable, resp)
		return
	}

	// get data
	formCode := c.Request.FormValue("code")
	log.Println("Callback - code from form", formCode)
	userData, err := ggoauth2.GetUserDataFromGoogle(
		c,
		formCode,
		ctrl.oauth2Cfg,
	)
	if err != nil {
		resp := new(ErrorDTO)
		resp.SetCode(http.StatusInternalServerError)
		resp.SetMessage("get google data error")
		c.AbortWithStatusJSON(http.StatusNotAcceptable, resp)
		return
	}

	log.Println("user data", string(userData))

	c.JSON(
		http.StatusOK,
		new(ErrorDTO).SetCode(http.StatusOK),
	)
}
