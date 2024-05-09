package server

import (
	"construction-organization-system/construction-organization-front/internal/log"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type Server struct {
	router     *mux.Router
	backendUrl string
}

func NewServer() *Server {
	backendUrl := os.Getenv("BACKEND_URL")
	router := mux.NewRouter()

	server := &Server{router: router, backendUrl: backendUrl}

	return server
}

func (s *Server) loadCSSStyles() {
	s.router.PathPrefix("/static/css/").Handler(http.StripPrefix("/static/css/", http.FileServer(http.Dir("static/css"))))
}

func (s *Server) RunServer() {
	s.initializeRoutes()
	s.loadCSSStyles()

	log.Logger.Infoln("Server start")
	log.Logger.Fatal(http.ListenAndServe(":80", s.router))
}
