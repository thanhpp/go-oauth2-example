package httpserver

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/thanhpp/go-oauth2-example/infrastructure/apdaters/ggoauth2"
)

func NewRouter() *gin.Engine {
	var router = new(gin.Engine)

	log.Println("router ok", router)

	googleOAuth := router.Group("/auth/google")
	{
		ctrl := newGoogleOAuthCtrl("/home/thanhpp/go/src/github.com/thanhpp/go-oauth2-example/secrets/client_secret_538939070983-pj4a1puc36b1trjgkpsk4uciv8u0eth3.apps.googleusercontent.com.json")

		log.Println("google ctrl ok", ctrl)

		googleOAuth.GET("/login", ctrl.LoginHandler)
		googleOAuth.GET("/callback", ctrl.CallbackHandler)
	}

	return router
}

func newGoogleOAuthCtrl(secretFile string) GoogleOAuthCtrl {
	oauth2Cfg, err := ggoauth2.NewOAuth2ConfigFromFile(secretFile)
	if err != nil {
		panic(err)
	}

	return NewGoogleOAuthCtrl(oauth2Cfg)
}
