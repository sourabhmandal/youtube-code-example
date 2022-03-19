package config

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

type Config struct {
	GoogleLoginConfig   oauth2.Config
	FacebookLoginConfig oauth2.Config
}

var AppConfig Config

const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
const OauthFacebookUrlAPI = "https://graph.facebook.com/v13.0/me?fields=id,name,email,picture&access_token&access_token="

func LoadConfig() {
	// Oauth configuration for Google
	AppConfig.GoogleLoginConfig = oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:8080/google_callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}

	// Oauth configuration for Facebook
	AppConfig.FacebookLoginConfig = oauth2.Config{
		ClientID:     os.Getenv("FB_CLIENT_ID"),
		ClientSecret: os.Getenv("FB_CLIENT_SECRET"),
		Endpoint:     facebook.Endpoint,
		RedirectURL:  "http://localhost:8080/fb_callback",
		Scopes: []string{
			"email",
			"public_profile",
		},
	}
}
