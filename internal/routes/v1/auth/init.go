package auth

import (
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

type ConfigGomniAuth struct {
	securityKey string
	googleClientId string
	googleClientSecret string
	redirectUrl string
}

func NewConfigGomniAuth(
	securityKey string,
	googleClientId string,
	googleClientSecret string,
	redirectUrl string,
) *ConfigGomniAuth {
	return &ConfigGomniAuth{
		securityKey: securityKey,
		googleClientId: googleClientId,
		googleClientSecret: googleClientSecret,
		redirectUrl: redirectUrl,
	}
}

func (c *ConfigGomniAuth) InitGomniauth() {
	gomniauth.SetSecurityKey(c.securityKey)
	gomniauth.WithProviders(
		google.New(
			c.googleClientId,
			c.googleClientSecret,
			c.redirectUrl,
		),
	)
}