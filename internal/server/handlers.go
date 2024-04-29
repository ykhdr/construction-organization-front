package server

import (
	"construction-organization-system/construction-organization-front/internal/log"
	"construction-organization-system/construction-organization-front/internal/view"
	"html/template"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := view.IndexPageData{
		PageData: view.PageData{
			PageTitle: "Construction System",
			Title:     "Construction System",
		},
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		log.Logger.Error("Error on executing Index template")
	}
}

func initializeRoutes() {
	http.HandleFunc("/", handleIndex)
}
