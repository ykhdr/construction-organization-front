package server

import (
	"construction-organization-system/construction-organization-front/internal/model"
	"construction-organization-system/construction-organization-front/internal/requests"
	"strconv"
)

func (s *Server) getProjects(managementID int) ([]model.ConstructionProject, error) {
	query := "http://" + s.backendUrl + "/api/v1/"
	if managementID != 0 {
		query = query + "construction_management/" + strconv.Itoa(managementID) + "/projects"
	} else {
		query = query + "construction_project"
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

func (s *Server) getConstructionTeams(projectID int) ([]model.ConstructionTeam, error) {
	var query = ""
	if projectID != 0 {
		query = "http://" + s.backendUrl + "/api/v1/construction_project/" + strconv.Itoa(projectID) + "/construction_teams"
	} else {
		query = "http://" + s.backendUrl + "/api/v1/construction_team"
	}

	teams, err := requests.GetConstructionTeams(query)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (s *Server) getMachines(projectID, managerID int, startDate, endDate string) ([]model.ConstructionMachinery, error) {
	var query = "http://" + s.backendUrl + "/api/v1"
	if projectID != 0 {
		query = query + "/construction_project/" + strconv.Itoa(projectID) + "/machines"
		if startDate != "" || endDate != "" {
			query = query + "?start_date=" + startDate + "&end_date=" + endDate
		}
	} else if managerID != 0 {
		query = query + "/construction_management/" + strconv.Itoa(managerID) + "/machines"
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

func (s *Server) getEngineers(managementID int) ([]model.EngineerWorker, error) {
	var query string

	if managementID != 0 {
		query = "http://" + s.backendUrl + "/api/v1/construction_management/" + strconv.Itoa(managementID) + "/engineers"
	} else {
		query = "http://" + s.backendUrl + "/api/v1/engineer_worker"
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

func (s *Server) getWorkTypes() ([]model.WorkType, error) {
	query := "http://" + s.backendUrl + "/api/v1/work_type"

	workTypes, err := requests.GetWorkTypes(query)
	if err != nil {
		return nil, err
	}

	return workTypes, nil
}
