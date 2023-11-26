package auth

import (
	authentication "floral/pkg/auth"
	"time"

	"floral/config"
	"fmt"
	"log"
)

var Token authentication.Maker

var BaseTokenDuration time.Duration = 60 * 24 * time.Hour

func init() {
	var err error
	Token, err = authentication.NewPasetoMaker(config.App.Token.Secret)
	if err != nil {
		log.Fatal(fmt.Errorf("error creating token maker: %w", err))
	}
}
