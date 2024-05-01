package server

import (
	"construction-organization-system/construction-organization-front/internal/model"
	"construction-organization-system/construction-organization-front/internal/requests"
	"strconv"
)

func (s *Server) getProjects() ([]model.ConstructionProject, error) {
	query := "http://" + s.backendUrl + "/api/v1/construction_project"
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
		query = "http://" + s.backendUrl + "/api/v1/construction_teams"
	}

	teams, err := requests.GetConstructionTeams(query)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func (s *Server) getMachines(projectID int, startDate, endDate string) ([]model.ConstructionMachinery, error) {
	var query = ""
	if projectID != 0 {
		query = "http://" + s.backendUrl + "/api/v1/construction_project/" + strconv.Itoa(projectID) + "/machines"
		if startDate != "" || endDate != "" {
			query = query + "?start_date=" + startDate + "&end_date=" + endDate
		}
	} else {
		query = "http://" + s.backendUrl + "/api/v1/construction_machinery"
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
	if err != nil {
		return model.Report{ID: 0}, err
	}

	return report, nil
}
