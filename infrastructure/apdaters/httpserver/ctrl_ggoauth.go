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

	c.SetCookie(
		oauthStateKey,
		oauthState,
		int(time.Now().Add(ctrl.oauthExpireDuration).Unix()),
		"",
		"",
		false,
		false,
	)

	redirectURL := ctrl.oauth2Cfg.AuthCodeURL(oauthState)

	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}

func (ctrl *GoogleOAuthCtrl) CallbackHandler(c *gin.Context) {
	// get oauth state from cookie
	oauthState, err := c.Request.Cookie(oauthStateKey)
	if err != nil {
		log.Println("state not found", err)
		resp := new(ErrorDTO)
		resp.SetCodeMsg(http.StatusNotAcceptable, "state not found")
		c.AbortWithStatusJSON(http.StatusNotAcceptable, resp)
		return
	}

	if c.Request.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth state")
		resp := new(ErrorDTO)
		resp.SetCode(http.StatusNotAcceptable)
		resp.SetMessage("invalid oauth state")
		c.AbortWithStatusJSON(http.StatusNotAcceptable, resp)
		return
	}

	userData, err := ggoauth2.GetUserDataFromGoogle(
		c,
		c.Request.FormValue("code"),
		ctrl.oauth2Cfg,
	)
	if err != nil {
		log.Println("get user data from google", err)
		resp := new(ErrorDTO)
		resp.SetCode(http.StatusInternalServerError)
		resp.SetMessage("get google data error")
		c.AbortWithStatusJSON(http.StatusNotAcceptable, resp)
		return
	}

	log.Println(string(userData))

	c.JSON(
		http.StatusOK,
		new(ErrorDTO).SetCode(http.StatusOK),
	)
}
