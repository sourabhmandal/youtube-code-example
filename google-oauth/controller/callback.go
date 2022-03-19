package controller

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/youtube/google-oauth/config"
)

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	// check is method is correct
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// get oauth state from cookie for this user
	oauthState, _ := r.Cookie("oauthstate")
	state := r.FormValue("state")
	code := r.FormValue("code")
	w.Header().Add("content-type", "application/json")
	
	// ERROR : Invalid OAuth State
	if state != oauthState.Value {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		fmt.Fprintf(w, "invalid oauth google state")
		return
	}

	// Exchange Auth Code for Tokens
	token, err := config.AppConfig.GoogleLoginConfig.Exchange(
		context.Background(), code)
	
	// ERROR : Auth Code Exchange Failed
	if err != nil {
		fmt.Fprintf(w, "falied code exchange: %s", err.Error())
		return
	}

	// Fetch User Data from google server
	response, err := http.Get(config.OauthGoogleUrlAPI + token.AccessToken)

	// ERROR : Unable to get user data from google
	if err != nil {
		fmt.Fprintf(w, "failed getting user info: %s", err.Error())
		return
	}

	// Parse user data JSON Object
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(w, "failed read response: %s", err.Error())
		return
	}

	// send back response to browser
	fmt.Fprintln(w, string(contents))
}

func FbCallback(w http.ResponseWriter, r *http.Request) {
	// check is method is correct
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// get oauth state from cookie for this user
	oauthState, _ := r.Cookie("oauthstate")
	state := r.FormValue("state")
	code := r.FormValue("code")
	w.Header().Add("content-type", "application/json")

	// ERROR : Invalid OAuth State
	if state != oauthState.Value {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		fmt.Fprintf(w, "invalid oauth google state")
		return
	}

	// Exchange Auth Code for Tokens
	token, err := config.AppConfig.FacebookLoginConfig.Exchange(
		context.Background(), code)

	// ERROR : Auth Code Exchange Failed
	if err != nil {
		fmt.Fprintf(w, "falied code exchange: %s", err.Error())
		return
	}

	// Fetch User Data from facebook server
	response, err := http.Get(config.OauthFacebookUrlAPI + token.AccessToken)
		
	// ERROR : Unable to get user data from google
	if err != nil {
		fmt.Fprintf(w, "failed getting user info: %s", err.Error())
		return
	}
	
	// Parse user data JSON Object
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(w, "failed read response: %s", err.Error())
		return
	}

	// send back response to browser
	fmt.Fprintln(w, string(contents))
}
