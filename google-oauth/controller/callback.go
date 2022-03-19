package controller

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/youtube/google-oauth/config"
)

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	oauthState, _ := r.Cookie("oauthstate")
	state := r.FormValue("state")
	code := r.FormValue("code")
	w.Header().Add("content-type", "application/json")

	if state != oauthState.Value {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		fmt.Fprintf(w, "invalid oauth google state")
		return
	}

	token, err := config.AppConfig.GoogleLoginConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Fprintf(w, "falied code exchange: %s", err.Error())
		return
	}
	response, err := http.Get(config.OauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		fmt.Fprintf(w, "failed getting user info: %s", err.Error())
		return
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(w, "failed read response: %s", err.Error())
		return
	}
	fmt.Fprintln(w, string(contents))
}

func FbCallback(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	oauthState, _ := r.Cookie("oauthstate")
	state := r.FormValue("state")
	code := r.FormValue("code")
	w.Header().Add("content-type", "application/json")

	if state != oauthState.Value {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		fmt.Fprintf(w, "invalid oauth google state")
		return
	}

	token, err := config.AppConfig.FacebookLoginConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Fprintf(w, "falied code exchange: %s", err.Error())
		return
	}

	response, err := http.Get(config.OauthFacebookUrlAPI + token.AccessToken)
	if err != nil {
		fmt.Fprintf(w, "failed getting user info: %s", err.Error())
		return
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(w, "failed read response: %s", err.Error())
		return
	}
	fmt.Fprintln(w, string(contents))
}
