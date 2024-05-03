package server

import (
	"construction-organization-system/construction-organization-front/internal/log"
	"construction-organization-system/construction-organization-front/internal/util"
	"construction-organization-system/construction-organization-front/internal/view"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
	"time"
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
	var (
		managementID   = 0
		buildingSiteID = 0

		managementIDStr   = r.URL.Query().Get("management_id")
		buildingSiteIDStr = r.URL.Query().Get("building_site_id")
	)

	if managementIDStr != "" {
		id, err := strconv.Atoi(managementIDStr)
		if err != nil {
			log.Logger.WithError(err).Error("Error on getting management id")
			http.Error(w, "Error on getting management id", http.StatusBadRequest)
			return
		}
		managementID = id
	}

	if buildingSiteIDStr != "" {
		id, err := strconv.Atoi(buildingSiteIDStr)
		if err != nil {
			log.Logger.WithError(err).Error("Error on getting building site id")
			http.Error(w, "Error on getting building site id", http.StatusBadRequest)
			return
		}
		buildingSiteID = id
	}

	tmpl := template.Must(template.ParseFiles("templates/projects.html"))
	projects, err := s.getProjects(managementID, buildingSiteID)

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
	var workTypeID int

	projectIDStr := r.URL.Query().Get("project_id")
	workTypeIDStr := r.URL.Query().Get("work_type")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if workTypeIDStr != "" {
		id, err := strconv.Atoi(workTypeIDStr)
		if err != nil {
			log.Logger.WithError(err).Error("Error on getting work type id")
			http.Error(w, "Error on getting work type id", http.StatusBadRequest)
			return
		}
		workTypeID = id
	} else {
		workTypeID = 0
	}

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

	if _, err := time.Parse("2006-01-02", startDate); startDate != "" && err != nil {
		http.Error(w, "Invalid start date format. Use YYYY-MM-DD format.", http.StatusBadRequest)
		return
	}

	if _, err := time.Parse("2006-01-02", endDate); endDate != "" && err != nil {
		http.Error(w, "Invalid end date format. Use YYYY-MM-DD format.", http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/construction_teams.html"))
	teams, err := s.getConstructionTeams(projectID, workTypeID, startDate, endDate)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting project schedules")
		http.Error(w, "Error on getting project schedules", http.StatusInternalServerError)
		return
	}
	workTypes, err := s.getWorkTypes(0, "", "")
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting work types")
		http.Error(w, "Error on getting work types", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{
		"ConstructionTeams": teams,
		"WorkTypes":         workTypes,
	})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing project schedules template")
		http.Error(w, "Error on executing project schedules template", http.StatusInternalServerError)
	}
}

func (s *Server) handleMachines(w http.ResponseWriter, r *http.Request) {
	var (
		projectID    = 0
		managementID = 0

		projectIDStr    = r.URL.Query().Get("project_id")
		managementIDStr = r.URL.Query().Get("management_id")
		startDate       = r.URL.Query().Get("start_date")
		endDate         = r.URL.Query().Get("end_date")
	)

	if projectIDStr != "" {
		id, err := strconv.Atoi(projectIDStr)
		if err != nil {
			log.Logger.WithError(err).Error("Error on getting project id")
			http.Error(w, "Error on getting project id", http.StatusBadRequest)
			return
		}

		projectID = id
	}

	if managementIDStr != "" {
		id, err := strconv.Atoi(managementIDStr)
		if err != nil {
			log.Logger.WithError(err).Error("Error on getting management id")
			http.Error(w, "Error on getting management id", http.StatusBadRequest)
			return
		}
		managementID = id
	}

	tmpl := template.Must(template.ParseFiles("templates/machines.html"))
	machines, err := s.getMachines(projectID, managementID, startDate, endDate)
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
	var (
		managementID   = 0
		buildingSiteID = 0

		managementIDStr   = r.URL.Query().Get("management_id")
		buildingSiteIDStr = r.URL.Query().Get("building_site_id")
	)

	if managementIDStr != "" {
		id, err := strconv.Atoi(managementIDStr)
		if err != nil {
			log.Logger.WithError(err).Error("Error on getting management id")
			http.Error(w, "Error on getting management id", http.StatusBadRequest)
			return
		}
		managementID = id
	}

	if buildingSiteIDStr != "" {
		id, err := strconv.Atoi(buildingSiteIDStr)
		if err != nil {
			log.Logger.WithError(err).Error("Error on getting building site id")
			http.Error(w, "Error on getting building site id", http.StatusBadRequest)
			return
		}
		buildingSiteID = id
	}

	tmpl := template.Must(template.ParseFiles("templates/engineers.html"))
	engineers, err := s.getEngineers(managementID, buildingSiteID)
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

func (s *Server) handleConstructionTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	teamID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting team id")
		http.Error(w, "Error on getting team id", http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/construction_team.html"))
	constructionTeam, err := s.getConstructionTeam(teamID)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting construction team")
		http.Error(w, "Error on getting construction team", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{"ConstructionTeam": constructionTeam})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing construction team template")
		http.Error(w, "Error on executing construction team template", http.StatusInternalServerError)
	}
}

func (s *Server) handleConstructionTeamWorkTypes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	teamID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting team id")
		http.Error(w, "Error on getting team id", http.StatusBadRequest)
		return
	}

	if _, err := time.Parse("2006-01-02", startDate); startDate != "" && err != nil {
		http.Error(w, "Invalid start date format. Use YYYY-MM-DD format.", http.StatusBadRequest)
		return
	}

	if _, err := time.Parse("2006-01-02", endDate); endDate != "" && err != nil {
		http.Error(w, "Invalid end date format. Use YYYY-MM-DD format.", http.StatusBadRequest)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/work_types.html"))
	workTypes, err := s.getWorkTypes(teamID, startDate, endDate)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting work types")
		http.Error(w, "Error on getting work types", http.StatusInternalServerError)
		return
	}
	constructionTeam, err := s.getConstructionTeam(teamID)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting construction team")
		http.Error(w, "Error on getting construction team", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{
		"WorkTypes":        workTypes,
		"ConstructionTeam": constructionTeam,
	})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing work types template")
		http.Error(w, "Error on executing work types template", http.StatusInternalServerError)
	}
}

func (s *Server) handleBuildingOrganization(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	organizationID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting organization id")
		http.Error(w, "Error on getting organization id", http.StatusBadRequest)
		return
	}

	organization, err := s.getBuildingOrganization(organizationID)
	tmpl := template.Must(template.ParseFiles("templates/building_organization.html"))
	err = tmpl.Execute(w, map[string]interface{}{"BuildingOrganization": organization})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing building organization template")
		http.Error(w, "Error on executing building organization template", http.StatusInternalServerError)
	}
}

func (s *Server) handleBuildingSite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	siteID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting site id")
		http.Error(w, "Error on getting site id", http.StatusBadRequest)
		return
	}

	site, err := s.getBuildingSite(siteID)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting site")
		http.Error(w, "Error on getting site", http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/building_site.html"))
	err = tmpl.Execute(w, map[string]interface{}{"BuildingSite": site})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing building site template")
		http.Error(w, "Error on executing building site template", http.StatusInternalServerError)
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
	s.router.HandleFunc("/construction_team/{id:[0-9]+}", s.handleConstructionTeam).Methods("GET")
	s.router.HandleFunc("/construction_team/{id:[0-9]+}/work_types", s.handleConstructionTeamWorkTypes).Methods("GET")

	s.router.HandleFunc("/machinery", s.handleMachines).Methods("GET")

	s.router.HandleFunc("/engineer", s.handleEngineers).Methods("GET")
	s.router.HandleFunc("/engineer/{id:[0-9]+}", s.handleEngineer).Methods("GET")

	s.router.HandleFunc("/management", s.handleManagements).Methods("GET")
	s.router.HandleFunc("/management/{id:[0-9]+}", s.handleManagement).Methods("GET")

	s.router.HandleFunc("/building_organization/{id:[0-9]+}", s.handleBuildingOrganization).Methods("GET")

	s.router.HandleFunc("/building_site/{id:[0-9]+}", s.handleBuildingSite).Methods("GET")
}
