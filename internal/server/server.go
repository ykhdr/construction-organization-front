package server

import (
	"construction-organization-system/construction-organization-front/internal/log"
	"net/http"
)

func loadCSSStyles() {
	http.Handle("/static/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("static/css"))))
}

func RunServer() {
	initializeRoutes()
	loadCSSStyles()
	log.Logger.Infoln("Server start")
	log.Logger.Fatal(http.ListenAndServe(":80", nil))
}
