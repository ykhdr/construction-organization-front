package server

import (
	"construction-organization-system/construction-organization-front/internal/log"
	"construction-organization-system/construction-organization-front/internal/model"
	"construction-organization-system/construction-organization-front/internal/util"
	"construction-organization-system/construction-organization-front/internal/view"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("/templates/index.html"))
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

	tmpl := template.Must(template.ParseFiles("/templates/projects.html"))
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

	tmpl := template.Must(template.ParseFiles("/templates/project.html"))

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

	tmpl := template.Must(template.ParseFiles("/templates/schedules.html"))
	schedules, err := s.getSchedules(projectID)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting project schedules")
		http.Error(w, "Error on getting project schedules", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{"Schedules": schedules, "ProjectID": projectID})
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
	tmpl = template.Must(tmpl.ParseFiles("/templates/estimate.html"))

	estimate, err := s.getEstimate(projectID)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting project estimate")
		http.Error(w, "Error on getting project estimate", http.StatusInternalServerError)
		return
	}
	materials, err := s.getExceededUsageMaterials(projectID)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting exceeded usage material")
		http.Error(w, "Error on getting exceeded usage material", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{"Estimate": estimate, "ExceededUsageMaterials": materials})
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

	tmpl := template.Must(template.ParseFiles("/templates/construction_teams.html"))
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

	tmpl := template.Must(template.ParseFiles("/templates/machines.html"))
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

	tmpl := template.Must(template.ParseFiles("/templates/exceeded_deadlines.html"))
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

	tmpl := template.Must(template.ParseFiles("/templates/reports.html"))
	reports, err := s.getReports(projectID)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting project estimate")
		http.Error(w, "Error on getting project estimate", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{
		"Reports":   reports,
		"ProjectID": projectID,
	})
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

	tmpl := template.Must(template.ParseFiles("/templates/report_file.html"))
	report, err := s.getReport(reportId)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting report")
		http.Error(w, "Error on getting report", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{
		"ReportFile": report.ReportFile,
	})
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

	tmpl := template.Must(template.ParseFiles("/templates/engineers.html"))
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

	tmpl := template.Must(template.ParseFiles("/templates/engineer.html"))
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
	tmpl := template.Must(template.ParseFiles("/templates/managements.html"))
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

	tmpl := template.Must(template.ParseFiles("/templates/management.html"))
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

	tmpl := template.Must(template.ParseFiles("/templates/construction_team.html"))
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

	tmpl := template.Must(template.ParseFiles("/templates/work_types.html"))
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
	tmpl := template.Must(template.ParseFiles("/templates/building_organization.html"))
	err = tmpl.Execute(w, map[string]interface{}{"BuildingOrganization": organization})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing building organization template")
		http.Error(w, "Error on executing building organization template", http.StatusInternalServerError)
	}
}

func (s *Server) handleBuildingSites(w http.ResponseWriter, r *http.Request) {
	sites, err := s.getBuildingSites()
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting building sites")
		http.Error(w, "Error on getting building sites", http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("/templates/building_sites.html"))
	err = tmpl.Execute(w, map[string]interface{}{"BuildingSites": sites})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing building sites template")
		http.Error(w, "Error on executing building sites template", http.StatusInternalServerError)
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
	tmpl := template.Must(template.ParseFiles("/templates/building_site.html"))
	err = tmpl.Execute(w, map[string]interface{}{"BuildingSite": site})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing building site template")
		http.Error(w, "Error on executing building site template", http.StatusInternalServerError)
	}
}

func (s *Server) handleCreateReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting project id")
		http.Error(w, "Error on getting project id", http.StatusBadRequest)
		return
	}

	err = s.createReport(projectID)
	if err != nil {
		log.Logger.WithError(err).Error("Error on creating report")
		http.Error(w, "Error on creating report", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/project/"+strconv.Itoa(projectID)+"/report", http.StatusSeeOther)
}

func (s *Server) handleCreateSchedulePage(w http.ResponseWriter, r *http.Request) {
	projects, err := s.getProjects(0, 0)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting projects")
		http.Error(w, "Error on getting projects", http.StatusInternalServerError)
		return
	}

	workTypes, err := s.getWorkTypes(0, "", "")
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting work types")
		http.Error(w, "Error on getting work types", http.StatusInternalServerError)
		return
	}

	teams, err := s.getConstructionTeams(0, 0, "", "")
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting construction teams")
		http.Error(w, "Error on getting construction teams", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("/templates/schedules_create.html"))
	err = tmpl.Execute(w, map[string]interface{}{"Projects": projects, "WorkTypes": workTypes, "ConstructionTeams": teams})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing schedule create template")
		http.Error(w, "Error on executing schedule create template", http.StatusInternalServerError)
	}
}

func (s *Server) handleCreateSchedule(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing the form", http.StatusBadRequest)
		return
	}

	constructionTeamID, err := strconv.Atoi(r.FormValue("construction_team_id"))
	if err != nil {
		http.Error(w, "Invalid construction team ID", http.StatusBadRequest)
		return
	}

	workTypeID, err := strconv.Atoi(r.FormValue("work_type_id"))
	if err != nil {
		http.Error(w, "Invalid work type ID", http.StatusBadRequest)
		return
	}

	projectID, err := strconv.Atoi(r.FormValue("project_id"))
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	planStartDate, err := time.Parse("2006-01-02", r.FormValue("plan_start_date"))
	if err != nil {
		http.Error(w, "Invalid plan start date", http.StatusBadRequest)
		return
	}

	planEndDate, err := time.Parse("2006-01-02", r.FormValue("plan_end_date"))
	if err != nil {
		http.Error(w, "Invalid plan end date", http.StatusBadRequest)
		return
	}

	factStartDate, _ := time.Parse("2006-01-02", r.FormValue("fact_start_date"))
	factEndDate, _ := time.Parse("2006-01-02", r.FormValue("fact_end_date"))

	planOrder, _ := strconv.Atoi(r.FormValue("plan_order"))
	factOrder, _ := strconv.Atoi(r.FormValue("fact_order"))

	newSchedule := &model.WorkSchedule{
		ConstructionTeamID: constructionTeamID,
		WorkType:           model.WorkType{ID: workTypeID},
		ProjectID:          projectID,
		PlanStartDate:      planStartDate,
		PlanEndDate:        planEndDate,
		FactStartDate:      factStartDate,
		FactEndDate:        factEndDate,
		PlanOrder:          planOrder,
		FactOrder:          factOrder,
	}

	err = s.saveWorkSchedule(newSchedule)
	if err != nil {
		log.Logger.WithError(err).Error("Error on saving work schedule")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/schedule?project_id="+strconv.Itoa(projectID), http.StatusSeeOther)
}

func (s *Server) handleCreateConstructionTeamPage(w http.ResponseWriter, r *http.Request) {
	projects, err := s.getProjects(0, 0)
	if err != nil {
		log.Logger.WithError(err).Error("Error on getting projects")
		http.Error(w, "Error on getting projects", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("/templates/construction_team_create.html"))
	err = tmpl.Execute(w, map[string]interface{}{"Projects": projects})
	if err != nil {
		log.Logger.WithError(err).Error("Error on executing schedule create template")
		http.Error(w, "Error on executing schedule create template", http.StatusInternalServerError)
	}
}

func (s *Server) handleCreateConstructionTeam(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing the form", http.StatusBadRequest)
		return
	}

	projectID, err := strconv.Atoi(r.FormValue("project_id"))
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")

	newTeam := &model.ConstructionTeam{
		Name:      name,
		ProjectID: projectID,
	}

	err = s.saveConstructionTeam(newTeam)
	if err != nil {
		log.Logger.WithError(err).Error("Error on saving construction team")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/construction_team?project_id="+strconv.Itoa(projectID), http.StatusSeeOther)
}

func (s *Server) handleSchedule(w http.ResponseWriter, r *http.Request) {

	scheduleID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid schedule ID", http.StatusBadRequest)
		return
	}

	schedule, err := s.getSchedule(scheduleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("/templates/schedule.html"))
	err = tmpl.Execute(w, map[string]interface{}{"Schedule": schedule})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) handleDeleteSchedule(w http.ResponseWriter, r *http.Request) {

	scheduleID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid schedule ID", http.StatusBadRequest)
		return
	}

	err = s.deleteSchedule(scheduleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/schedule", http.StatusSeeOther)
}

func (s *Server) handleDeleteConstructionTeam(w http.ResponseWriter, r *http.Request) {

	teamID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid construction team ID", http.StatusBadRequest)
		return
	}

	err = s.deleteConstructionTeam(teamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/construction_team", http.StatusSeeOther)
}

func (s *Server) handleUpdateSchedule(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(r.FormValue("id"))
	constructionTeamID, _ := strconv.Atoi(r.FormValue("construction_team_id"))
	workTypeID, _ := strconv.Atoi(r.FormValue("work_type_id"))
	planStartDate, _ := time.Parse("2006-01-02", r.FormValue("plan_start_date"))
	planEndDate, _ := time.Parse("2006-01-02", r.FormValue("plan_end_date"))
	factStartDate, _ := time.Parse("2006-01-02", r.FormValue("fact_start_date"))
	factEndDate, _ := time.Parse("2006-01-02", r.FormValue("fact_end_date"))
	planOrder, _ := strconv.Atoi(r.FormValue("plan_order"))
	factOrder, _ := strconv.Atoi(r.FormValue("fact_order"))
	projectID, _ := strconv.Atoi(r.FormValue("project_id"))

	schedule := model.WorkSchedule{
		ID:                 id,
		ConstructionTeamID: constructionTeamID,
		WorkType:           model.WorkType{ID: workTypeID},
		PlanStartDate:      planStartDate,
		PlanEndDate:        planEndDate,
		FactStartDate:      factStartDate,
		FactEndDate:        factEndDate,
		PlanOrder:          planOrder,
		FactOrder:          factOrder,
		ProjectID:          projectID,
	}

	err = s.updateSchedule(&schedule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/schedule/"+strconv.Itoa(id), http.StatusSeeOther)
}

func (s *Server) handleUpdateSchedulePage(w http.ResponseWriter, r *http.Request) {

	scheduleID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid schedule ID", http.StatusBadRequest)
		return
	}

	schedule, err := s.getSchedule(scheduleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("/templates/schedule_edit.html"))
	err = tmpl.Execute(w, map[string]interface{}{"Schedule": schedule})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) handleEngineerTeams(w http.ResponseWriter, r *http.Request) {
	engineerTeams, err := s.getEngineerTeams()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("/templates/engineer_teams.html"))
	err = tmpl.Execute(w, map[string]interface{}{"EngineerTeams": engineerTeams})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (s *Server) handleEngineerTeam(w http.ResponseWriter, r *http.Request) {

	engineerTeamID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid engineer team ID", http.StatusBadRequest)
		return
	}

	engineerTeam, err := s.getEngineerTeam(engineerTeamID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("/templates/engineer_team.html"))
	err = tmpl.Execute(w, map[string]interface{}{"EngineerTeam": engineerTeam})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
	s.router.HandleFunc("/report/{id:[0-9]+}/create", s.handleCreateReport).Methods("GET")

	s.router.HandleFunc("/schedule", s.handleSchedules).Methods("GET")
	s.router.HandleFunc("/schedule/{id:[0-9]+}", s.handleSchedule).Methods("GET")
	s.router.HandleFunc("/schedule/{id:[0-9]+}", s.handleDeleteSchedule).Methods("DELETE")
	s.router.HandleFunc("/schedule/{id:[0-9]+}/update", s.handleUpdateSchedule).Methods("POST")
	s.router.HandleFunc("/schedule/{id:[0-9]+}/update", s.handleUpdateSchedulePage).Methods("GET")

	s.router.HandleFunc("/schedule/create", s.handleCreateSchedulePage).Methods("GET")
	s.router.HandleFunc("/schedule/create", s.handleCreateSchedule).Methods("POST")

	s.router.HandleFunc("/construction_team", s.handleConstructionTeams).Methods("GET")
	s.router.HandleFunc("/construction_team/create", s.handleCreateConstructionTeamPage).Methods("GET")
	s.router.HandleFunc("/construction_team/create", s.handleCreateConstructionTeam).Methods("POST")
	s.router.HandleFunc("/construction_team/{id:[0-9]+}", s.handleConstructionTeam).Methods("GET")
	s.router.HandleFunc("/construction_team/{id:[0-9]+}", s.handleDeleteConstructionTeam).Methods("DELETE")
	s.router.HandleFunc("/construction_team/{id:[0-9]+}/work_types", s.handleConstructionTeamWorkTypes).Methods("GET")

	s.router.HandleFunc("/machinery", s.handleMachines).Methods("GET")

	s.router.HandleFunc("/engineer", s.handleEngineers).Methods("GET")
	s.router.HandleFunc("/engineer/{id:[0-9]+}", s.handleEngineer).Methods("GET")

	s.router.HandleFunc("/engineer_team", s.handleEngineerTeams).Methods("GET")
	s.router.HandleFunc("/engineer_team/{id:[0-9]+}", s.handleEngineerTeam).Methods("GET")

	s.router.HandleFunc("/management", s.handleManagements).Methods("GET")
	s.router.HandleFunc("/management/{id:[0-9]+}", s.handleManagement).Methods("GET")

	s.router.HandleFunc("/building_organization/{id:[0-9]+}", s.handleBuildingOrganization).Methods("GET")

	s.router.HandleFunc("/building_site", s.handleBuildingSites).Methods("GET")
	s.router.HandleFunc("/building_site/{id:[0-9]+}", s.handleBuildingSite).Methods("GET")
}
