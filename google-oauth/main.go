package main

import (
	"log"
	"net/http"

	"github.com/youtube/google-oauth/config"
	"github.com/youtube/google-oauth/controller"
)

func main() {
	mux := http.NewServeMux()
	config.LoadEnv()
	config.LoadConfig()
	mux.HandleFunc("/google_login", controller.GoogleLogin)
	mux.HandleFunc("/google_callback", controller.GoogleCallback)
	mux.HandleFunc("/fb_login", controller.FbLogin)
	mux.HandleFunc("/fb_callback", controller.FbCallback)

	log.Println("started server on :: http://localhost:8080/")
	if oops := http.ListenAndServe(":8080", mux); oops != nil {
		log.Fatal(oops)
	}
}