package ggoauth2_test

import (
	"fmt"
	"testing"

	"github.com/thanhpp/go-oauth2-example/infrastructure/apdaters/ggoauth2"
)

func TestNewConfigFromFile(t *testing.T) {
	var (
		path = "/home/thanhpp/go/src/github.com/thanhpp/go-oauth2-example/secrets/client_secret_538939070983-pj4a1puc36b1trjgkpsk4uciv8u0eth3.apps.googleusercontent.com.json"
	)
	cfg, err := ggoauth2.NewConfigFromFile(path)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(cfg)
}
