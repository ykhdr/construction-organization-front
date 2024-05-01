package requests

import (
	"construction-organization-system/construction-organization-front/internal/model"
	"encoding/json"
	"net/http"
)

func GetProjects(query string) ([]model.ConstructionProject, error) {
	response, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	var projects []model.ConstructionProject
	err = json.NewDecoder(response.Body).Decode(&projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func GetProject(query string) (model.ConstructionProject, error) {
	response, err := http.Get(query)
	if err != nil {
		return model.ConstructionProject{ID: 0, Name: "UNKNOWN"}, err
	}

	var project model.ConstructionProject
	err = json.NewDecoder(response.Body).Decode(&project)
	if err != nil {
		return model.ConstructionProject{ID: 0}, err
	}

	return project, nil
}

func GetSchedules(query string) ([]model.WorkSchedule, error) {
	response, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	var schedules []model.WorkSchedule
	err = json.NewDecoder(response.Body).Decode(&schedules)
	if err != nil {
		return nil, err
	}

	return schedules, nil
}

func GetEstimate(query string) (model.Estimate, error) {
	response, err := http.Get(query)
	if err != nil {
		return model.Estimate{ID: 0}, err
	}

	var estimate model.Estimate
	err = json.NewDecoder(response.Body).Decode(&estimate)
	if err != nil {
		return model.Estimate{ID: 0}, err
	}

	return estimate, nil
}

func GetConstructionTeams(query string) ([]model.ConstructionTeam, error) {
	response, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	var teams []model.ConstructionTeam
	err = json.NewDecoder(response.Body).Decode(&teams)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func GetMachines(query string) ([]model.ConstructionMachinery, error) {
	response, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	var machines []model.ConstructionMachinery
	err = json.NewDecoder(response.Body).Decode(&machines)
	if err != nil {
		return nil, err
	}

	return machines, nil
}
