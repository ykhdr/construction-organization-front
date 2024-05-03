package server

import (
	"construction-organization-system/construction-organization-front/internal/model"
	"construction-organization-system/construction-organization-front/internal/requests"
	"strconv"
)

func (s *Server) getProjects(managementID, buildingSiteID int) ([]model.ConstructionProject, error) {
	query := "http://" + s.backendUrl + "/api/v1"
	if managementID != 0 {
		query = query + "/construction_management/" + strconv.Itoa(managementID) + "/projects"
	} else if buildingSiteID != 0 {
		query = query + "/building_site/" + strconv.Itoa(buildingSiteID) + "/projects"
	} else {
		query = query + "/construction_project"
	}
	projects, err := requests.GetProjects(query)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (s *Server) getProject(id int) (model.ConstructionProject, error) {
	query := "http://" + s.backendUrl + "/api/v1/construction_project/" + strconv.Itoa(id)
	project, err := requests.GetProject(query)
	if err != nil {
		return model.ConstructionProject{ID: 0}, err
	}

	return project, nil
}

func (s *Server) getSchedules(projectID int) ([]model.WorkSchedule, error) {
	query := ""

	if projectID != 0 {
		query = "http://" + s.backendUrl + "/api/v1/construction_project/" + strconv.Itoa(projectID) + "/schedules"
	} else {
		query = "http://" + s.backendUrl + "/api/v1/work_schedule"
	}

	schedules, err := requests.GetSchedules(query)

	if err != nil {
		return nil, err
	}

	return schedules, nil
}

func (s *Server) getEstimate(projectID int) (model.Estimate, error) {
	query := "http://" + s.backendUrl + "/api/v1/construction_project/" + strconv.Itoa(projectID) + "/estimate"
	estimate, err := requests.GetEstimate(query)

	if err != nil {
		return model.Estimate{ID: 0}, err
	}

	return estimate, nil
}

func (s *Server) getConstructionTeams(projectID, workTypeID int, startDate, endDate string) ([]model.ConstructionTeam, error) {
	var query = "http://" + s.backendUrl + "/api/v1"
	if projectID != 0 {
		query = query + "/construction_project/" + strconv.Itoa(projectID) + "/construction_teams"
	} else if workTypeID != 0 {
		query = query + "/construction_team?work_type=" + strconv.Itoa(workTypeID) + "&start_date=" + startDate + "&end_date=" + endDate
	} else {
		query = query + "/construction_team"
	}

	teams, err := requests.GetConstructionTeams(query)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (s *Server) getMachines(projectID, managementID int, startDate, endDate string) ([]model.ConstructionMachinery, error) {
	var query = "http://" + s.backendUrl + "/api/v1"
	if projectID != 0 {
		query = query + "/construction_project/" + strconv.Itoa(projectID) + "/machines"
		if startDate != "" || endDate != "" {
			query = query + "?start_date=" + startDate + "&end_date=" + endDate
		}
	} else if managementID != 0 {
		query = query + "/construction_management/" + strconv.Itoa(managementID) + "/machines"
	} else {
		query = query + "/construction_machinery"
	}

	machines, err := requests.GetMachines(query)
	if err != nil {
		return nil, err
	}

	return machines, nil
}

func (s *Server) getExceededDeadlinesWorks(projectID int) ([]model.WorkType, error) {
	query := "http://" + s.backendUrl + "/api/v1/construction_project/" + strconv.Itoa(projectID) + "/exceeded_deadlines_works"

	workTypes, err := requests.GetWorkTypes(query)

	if err != nil {
		return nil, err
	}

	return workTypes, nil
}

func (s *Server) getReports(projectID int) ([]model.Report, error) {
	query := "http://localhost:8081/api/v1/report?project_id=" + strconv.Itoa(projectID)

	reports, err := requests.GetReports(query)
	if err != nil {
		return nil, err
	}

	return reports, nil
}

func (s *Server) getReport(reportID int) (model.Report, error) {
	query := "http://localhost:8081/api/v1/report/" + strconv.Itoa(reportID)

	report, err := requests.GetReport(query)

	//model.ReportFile = json.Unmarshal(, &report.ReportFile)

	if err != nil {
		return model.Report{ID: 0}, err
	}

	return report, nil
}

func (s *Server) getEngineers(managementID, buildingSiteID int) ([]model.EngineerWorker, error) {
	query := "http://" + s.backendUrl + "/api/v1"

	if managementID != 0 {
		query = query + "/construction_management/" + strconv.Itoa(managementID) + "/engineers"
	} else if buildingSiteID != 0 {
		query = query + "/building_site/" + strconv.Itoa(buildingSiteID) + "/engineers"
	} else {
		query = query + "/engineer_worker"
	}

	engineers, err := requests.GetEngineers(query)
	if err != nil {
		return nil, err
	}

	return engineers, nil
}

func (s *Server) getEngineer(engineerID int) (model.EngineerWorker, error) {
	query := "http://" + s.backendUrl + "/api/v1/engineer_worker/" + strconv.Itoa(engineerID)

	engineer, err := requests.GetEngineer(query)
	if err != nil {
		return model.EngineerWorker{ID: 0}, err
	}

	return engineer, nil
}

func (s *Server) getManagements() ([]model.ConstructionManagement, error) {
	query := "http://" + s.backendUrl + "/api/v1/construction_management"

	managements, err := requests.GetManagements(query)
	if err != nil {
		return nil, err
	}

	return managements, nil
}

func (s *Server) getManagement(managementID int) (model.ConstructionManagement, error) {
	query := "http://" + s.backendUrl + "/api/v1/construction_management/" + strconv.Itoa(managementID)

	management, err := requests.GetManagement(query)
	if err != nil {
		return model.ConstructionManagement{ID: 0}, err
	}

	return management, nil
}

func (s *Server) getWorkTypes(teamID int, startDate, endDate string) ([]model.WorkType, error) {
	query := "http://" + s.backendUrl + "/api/v1"
	if teamID != 0 {
		query += "/construction_team/" + strconv.Itoa(teamID) + "/work_types" + "?start_date=" + startDate + "&end_date=" + endDate
	} else {
		query += "/work_types"
	}

	workTypes, err := requests.GetWorkTypes(query)
	if err != nil {
		return nil, err
	}

	return workTypes, nil
}

func (s *Server) getConstructionTeam(teamID int) (model.ConstructionTeam, error) {
	query := "http://" + s.backendUrl + "/api/v1/construction_team/" + strconv.Itoa(teamID)

	team, err := requests.GetConstructionTeam(query)
	if err != nil {
		return model.ConstructionTeam{ID: 0}, err
	}

	return team, nil
}

func (s *Server) getBuildingOrganization(organizationID int) (model.BuildingOrganization, error) {
	query := "http://" + s.backendUrl + "/api/v1/building_organization/" + strconv.Itoa(organizationID)

	organization, err := requests.GetBuildingOrganization(query)
	if err != nil {
		return model.BuildingOrganization{ID: 0}, err
	}

	return organization, nil
}

func (s *Server) getBuildingSite(siteID int) (model.BuildingSite, error) {
	query := "http://" + s.backendUrl + "/api/v1/building_site/" + strconv.Itoa(siteID)

	site, err := requests.GetBuildingSite(query)
	if err != nil {
		return model.BuildingSite{ID: 0}, err
	}

	return site, nil
}

func (s *Server) getExceededUsageMaterials(estimateID int) ([]model.Material, error) {
	query := "http://" + s.backendUrl + "/api/v1/estimate/" + strconv.Itoa(estimateID) + "/exceeded_usage_material"

	materials, err := requests.GetMaterials(query)
	if err != nil {
		return nil, err
	}

	return materials, nil
}

func (s *Server) createReport(projectID int) error {
	query := "http://localhost:8081/api/v1/report/create?project_id=" + strconv.Itoa(projectID)

	err := requests.CreateReport(query)
	return err
}

func (s *Server) getBuildingSites() ([]model.BuildingSite, error) {
	query := "http://" + s.backendUrl + "/api/v1/building_site"

	sites, err := requests.GetBuildingSites(query)
	if err != nil {
		return nil, err
	}

	return sites, nil
}
