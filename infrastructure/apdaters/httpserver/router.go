package httpserver

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/thanhpp/go-oauth2-example/infrastructure/apdaters/ggoauth2"
)

func NewRouter() *gin.Engine {
	var router = gin.New()
	router.Use(gin.Logger())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	googleOAuthGr := router.Group("/auth/google")
	{
		ctrl := newGoogleOAuthCtrl("/home/thanhpp/go/src/github.com/thanhpp/go-oauth2-example/secrets/client_secret_538939070983-pj4a1puc36b1trjgkpsk4uciv8u0eth3.apps.googleusercontent.com.json")

		googleOAuthGr.GET("/login", ctrl.LoginHandler)

		log.Println("login ctrl ok")

		googleOAuthGr.GET("/callback", ctrl.CallbackHandler)
	}

	log.Println("router ok", router)

	return router
}

func newGoogleOAuthCtrl(secretFile string) GoogleOAuthCtrl {
	oauth2Cfg, err := ggoauth2.NewOAuth2ConfigFromFile(secretFile)
	if err != nil {
		panic(err)
	}

	return NewGoogleOAuthCtrl(oauth2Cfg)
}
