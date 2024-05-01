package server

import (
	"construction-organization-system/construction-organization-front/internal/log"
	"construction-organization-system/construction-organization-front/internal/util"
	"construction-organization-system/construction-organization-front/internal/view"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) handleProjects(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/projects.html"))

	projects, err := s.getProjects()

	if err != nil {
		log.Logger.WithError(err).Error("Error on getting projects")
		http.Error(w, "Error on getting projects", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{"Projects": projects})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing projects template")
		http.Error(w, "Error on executing projects template", http.StatusInternalServerError)
	}

}

func (s *Server) handleProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting project id")
		http.Error(w, "Error on getting project id", http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/project.html"))

	project, err := s.getProject(id)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting project")
		http.Error(w, "Error on getting project", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{"Project": project})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing project template")
		http.Error(w, "Error on executing project template", http.StatusInternalServerError)
	}
}

func (s *Server) handleSchedules(w http.ResponseWriter, r *http.Request) {
	var projectID int

	projectIDStr := r.URL.Query().Get("project_id")

	if projectIDStr != "" {
		id, err := strconv.Atoi(projectIDStr)
		if err != nil {
			log.Logger.WithError(err).Error("Error on getting project id")
			http.Error(w, "Error on getting project id", http.StatusBadRequest)
			return
		}
		projectID = id
	} else {
		projectID = 0
	}

	tmpl := template.Must(template.ParseFiles("templates/schedules.html"))
	schedules, err := s.getSchedules(projectID)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting project schedules")
		http.Error(w, "Error on getting project schedules", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{"Schedules": schedules})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing project schedules template")
		http.Error(w, "Error on executing project schedules template", http.StatusInternalServerError)
	}
}

func (s *Server) handleEstimate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting project id")
		http.Error(w, "Error on getting project id", http.StatusBadRequest)
		return
	}

	tmpl := template.New("estimate.html").Funcs(template.FuncMap{"mul": util.Mul})
	tmpl = template.Must(tmpl.ParseFiles("templates/estimate.html"))

	estimate, err := s.getEstimate(projectID)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting project estimate")
		http.Error(w, "Error on getting project estimate", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{"Estimate": estimate})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing project estimate template")
		http.Error(w, "Error on executing project estimate template", http.StatusInternalServerError)
	}
}

func (s *Server) handleConstructionTeams(w http.ResponseWriter, r *http.Request) {
	var projectID int
	projectIDStr := r.URL.Query().Get("project_id")

	if projectIDStr != "" {
		id, err := strconv.Atoi(projectIDStr)
		if err != nil {
			log.Logger.WithError(err).Error("Error on getting project id")
			http.Error(w, "Error on getting project id", http.StatusBadRequest)
			return
		}

		projectID = id
	} else {
		projectID = 0
	}

	tmpl := template.Must(template.ParseFiles("templates/construction_teams.html"))
	teams, err := s.getConstructionTeams(projectID)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting project schedules")
		http.Error(w, "Error on getting project schedules", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{"ConstructionTeams": teams})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing project schedules template")
		http.Error(w, "Error on executing project schedules template", http.StatusInternalServerError)
	}
}

func (s *Server) handleMachines(w http.ResponseWriter, r *http.Request) {
	var projectID int
	projectIDStr := r.URL.Query().Get("project_id")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if projectIDStr != "" {
		id, err := strconv.Atoi(projectIDStr)
		if err != nil {
			log.Logger.WithError(err).Error("Error on getting project id")
			http.Error(w, "Error on getting project id", http.StatusBadRequest)
			return
		}

		projectID = id
	} else {
		projectID = 0
	}

	tmpl := template.Must(template.ParseFiles("templates/machines.html"))
	machines, err := s.getMachines(projectID, startDate, endDate)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting project machines")
		http.Error(w, "Error on getting project machines", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{"Machines": machines, "ProjectID": projectID})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing project machines template")
		http.Error(w, "Error on executing project machines template", http.StatusInternalServerError)
	}
}

func (s *Server) handleExceededDeadlinesWorks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting project id")
		http.Error(w, "Error on getting project id", http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/exceeded_deadlines.html"))
	workTypes, err := s.getExceededDeadlinesWorks(projectID)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting project exceeded deadlines works")
		http.Error(w, "Error on getting project exceeded deadlines works", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{"WorkTypes": workTypes})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing project exceeded deadlines works template")
		http.Error(w, "Error on executing project exceeded deadlines works template", http.StatusInternalServerError)
	}
}

func (s *Server) handleReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting project id")
		http.Error(w, "Error on getting project id", http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/report.html"))
	report, err := s.getReports(projectID)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting project estimate")
		http.Error(w, "Error on getting project estimate", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{"Report": report})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing project report template")
		http.Error(w, "Error on executing project report template", http.StatusInternalServerError)
	}
}

func (s *Server) initializeRoutes() {
	s.router.HandleFunc("/", s.handleIndex).Methods("GET")
	s.router.HandleFunc("/project", s.handleProjects).Methods("GET")
	s.router.HandleFunc("/project/{id:[0-9]+}", s.handleProject).Methods("GET")
	s.router.HandleFunc("/project/{id:[0-9]+}/estimate", s.handleEstimate).Methods("GET")
	s.router.HandleFunc("/project/{id:[0-9]+}/exceeded_deadlines_works", s.handleExceededDeadlinesWorks).Methods("GET")
	s.router.HandleFunc("/project/{id:[0-9]+}/report", s.handleReport).Methods("GET")
	s.router.HandleFunc("/schedule", s.handleSchedules).Methods("GET")
	s.router.HandleFunc("/construction_team", s.handleConstructionTeams).Methods("GET")
	s.router.HandleFunc("/machinery", s.handleMachines).Methods("GET")
}
