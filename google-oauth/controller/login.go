package controller

import (
	"net/http"

	"github.com/youtube/google-oauth/config"
	"github.com/youtube/google-oauth/utils"
)

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Create oauthState cookie
	oauthState := utils.GenerateStateOauthCookie(w)
	/*
		AuthCodeURL receive state that is a token to protect the user from CSRF attacks.
		You must always provide a non-empty string and validate that it matches the
		the state query parameter on your redirect callback.
	*/
	u := config.AppConfig.GoogleLoginConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func FbLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Create oauthState cookie
	oauthState := utils.GenerateStateOauthCookie(w)

	/*
		AuthCodeURL receive state that is a token to protect the user from CSRF attacks.
		You must always provide a non-empty string and validate that it matches the
		the state query parameter on your redirect callback.
	*/
	u := config.AppConfig.FacebookLoginConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}
