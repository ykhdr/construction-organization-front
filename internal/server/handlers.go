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
	var managementID int
	managementIDStr := r.URL.Query().Get("management_id")

	if managementIDStr != "" {
		id, err := strconv.Atoi(managementIDStr)
		if err != nil {
			log.Logger.WithError(err).Error("Error on getting management id")
			http.Error(w, "Error on getting management id", http.StatusBadRequest)
			return
		}

		managementID = id
	} else {
		managementID = 0
	}

	tmpl := template.Must(template.ParseFiles("templates/projects.html"))
	projects, err := s.getProjects(managementID)

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
	var managerID int
	managerIDStr := r.URL.Query().Get("manager_id")
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

	if managerIDStr != "" {
		id, err := strconv.Atoi(managerIDStr)
		if err != nil {
			log.Logger.WithError(err).Error("Error on getting manager id")
			http.Error(w, "Error on getting manager id", http.StatusBadRequest)
			return
		}

		managerID = id
	} else {
		managerID = 0
	}

	tmpl := template.Must(template.ParseFiles("templates/machines.html"))
	machines, err := s.getMachines(projectID, managerID, startDate, endDate)
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

func (s *Server) handleReports(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting project id")
		http.Error(w, "Error on getting project id", http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/reports.html"))
	reports, err := s.getReports(projectID)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting project estimate")
		http.Error(w, "Error on getting project estimate", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{"Reports": reports})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing project report template")
		http.Error(w, "Error on executing project report template", http.StatusInternalServerError)
	}
}

func (s *Server) handleReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reportId, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Logger.WithError(err).Error("Error on getting report id")
		http.Error(w, "Error on getting report id", http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/report_file.html"))
	report, err := s.getReport(reportId)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting report")
		http.Error(w, "Error on getting report", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{"ReportFile": report.ReportFile})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing report template")
		http.Error(w, "Error on executing report template", http.StatusInternalServerError)
	}
}

func (s *Server) handleEngineers(w http.ResponseWriter, r *http.Request) {
	var managementID int
	managementIDStr := r.URL.Query().Get("management_id")

	if managementIDStr != "" {
		id, err := strconv.Atoi(managementIDStr)
		if err != nil {
			log.Logger.WithError(err).Error("Error on getting management id")
			http.Error(w, "Error on getting management id", http.StatusBadRequest)
			return
		}

		managementID = id
	} else {
		managementID = 0
	}

	tmpl := template.Must(template.ParseFiles("templates/engineers.html"))
	engineers, err := s.getEngineers(managementID)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting engineers")
		http.Error(w, "Error on getting engineers", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{"Engineers": engineers})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing engineers template")
		http.Error(w, "Error on executing engineers template", http.StatusInternalServerError)
	}
}

func (s *Server) handleEngineer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	engineerID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting engineer id")
		http.Error(w, "Error on getting engineer id", http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/engineer.html"))
	engineer, err := s.getEngineer(engineerID)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting engineer")
		http.Error(w, "Error on getting engineer", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{"Engineer": engineer})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing engineer template")
		http.Error(w, "Error on executing engineer template", http.StatusInternalServerError)
	}
}

func (s *Server) handleManagements(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/managements.html"))
	managements, err := s.getManagements()
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting managements")
		http.Error(w, "Error on getting managements", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{"Managements": managements})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing managements template")
		http.Error(w, "Error on executing managements template", http.StatusInternalServerError)
	}
}

func (s *Server) handleManagement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	managementID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting management id")
		http.Error(w, "Error on getting management id", http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/management.html"))
	management, err := s.getManagement(managementID)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting management")
		http.Error(w, "Error on getting management", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{"Management": management})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing management template")
		http.Error(w, "Error on executing management template", http.StatusInternalServerError)
	}
}

func (s *Server) initializeRoutes() {
	s.router.HandleFunc("/", s.handleIndex).Methods("GET")
	s.router.HandleFunc("/project", s.handleProjects).Methods("GET")
	s.router.HandleFunc("/project/{id:[0-9]+}", s.handleProject).Methods("GET")
	s.router.HandleFunc("/project/{id:[0-9]+}/estimate", s.handleEstimate).Methods("GET")
	s.router.HandleFunc("/project/{id:[0-9]+}/exceeded_deadlines_works", s.handleExceededDeadlinesWorks).Methods("GET")
	s.router.HandleFunc("/project/{id:[0-9]+}/report", s.handleReports).Methods("GET")
	s.router.HandleFunc("/report/{id:[0-9]+}", s.handleReport).Methods("GET")
	s.router.HandleFunc("/schedule", s.handleSchedules).Methods("GET")
	s.router.HandleFunc("/construction_team", s.handleConstructionTeams).Methods("GET")
	s.router.HandleFunc("/machinery", s.handleMachines).Methods("GET")
	s.router.HandleFunc("/engineer", s.handleEngineers).Methods("GET")
	s.router.HandleFunc("/engineer/{id:[0-9]+}", s.handleEngineer).Methods("GET")
	s.router.HandleFunc("/management", s.handleManagements).Methods("GET")
	s.router.HandleFunc("/management/{id:[0-9]+}", s.handleManagement).Methods("GET")
}
