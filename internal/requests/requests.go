package requests

import (
	"bytes"
	"construction-organization-system/construction-organization-front/internal/model"
	"encoding/json"
	"errors"
	"io"
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

func GetWorkTypes(query string) ([]model.WorkType, error) {
	response, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	var workTypes []model.WorkType
	err = json.NewDecoder(response.Body).Decode(&workTypes)
	if err != nil {
		return nil, err
	}

	return workTypes, nil
}

func GetReports(query string) ([]model.Report, error) {
	response, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	var reports []model.Report
	err = json.NewDecoder(response.Body).Decode(&reports)
	if err != nil {
		return nil, err
	}

	return reports, nil
}

func GetReport(query string) (model.Report, error) {
	response, err := http.Get(query)
	if err != nil {
		return model.Report{ID: 0}, err
	}

	var report model.Report

	err = json.NewDecoder(response.Body).Decode(&report)
	if err != nil {
		return model.Report{ID: 0}, err
	}

	return report, nil
}

func GetEngineers(query string) ([]model.EngineerWorker, error) {
	response, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	var engineers []model.EngineerWorker
	err = json.NewDecoder(response.Body).Decode(&engineers)
	if err != nil {
		return nil, err
	}

	return engineers, nil
}

func GetEngineer(query string) (model.EngineerWorker, error) {
	response, err := http.Get(query)
	if err != nil {
		return model.EngineerWorker{ID: 0}, err
	}

	var engineer model.EngineerWorker
	err = json.NewDecoder(response.Body).Decode(&engineer)
	if err != nil {
		return model.EngineerWorker{ID: 0}, err
	}

	return engineer, nil
}

func GetManagements(query string) ([]model.ConstructionManagement, error) {
	response, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	var managements []model.ConstructionManagement
	err = json.NewDecoder(response.Body).Decode(&managements)
	if err != nil {
		return nil, err
	}

	return managements, nil
}

func GetManagement(query string) (model.ConstructionManagement, error) {
	response, err := http.Get(query)
	if err != nil {
		return model.ConstructionManagement{ID: 0}, err
	}

	var management model.ConstructionManagement
	err = json.NewDecoder(response.Body).Decode(&management)
	if err != nil {
		return model.ConstructionManagement{ID: 0}, err
	}

	return management, nil
}

func GetConstructionTeam(query string) (model.ConstructionTeam, error) {
	response, err := http.Get(query)
	if err != nil {
		return model.ConstructionTeam{ID: 0}, err
	}

	var team model.ConstructionTeam
	err = json.NewDecoder(response.Body).Decode(&team)
	if err != nil {
		return model.ConstructionTeam{ID: 0}, err
	}

	return team, nil
}

func GetBuildingOrganization(query string) (model.BuildingOrganization, error) {
	response, err := http.Get(query)
	if err != nil {
		return model.BuildingOrganization{ID: 0}, err
	}

	var organization model.BuildingOrganization
	err = json.NewDecoder(response.Body).Decode(&organization)
	if err != nil {
		return model.BuildingOrganization{ID: 0}, err
	}

	return organization, nil
}

func GetBuildingSite(query string) (model.BuildingSite, error) {
	response, err := http.Get(query)
	if err != nil {
		return model.BuildingSite{ID: 0}, err
	}

	var site model.BuildingSite
	err = json.NewDecoder(response.Body).Decode(&site)
	if err != nil {
		return model.BuildingSite{ID: 0}, err
	}

	return site, nil
}

func GetMaterials(query string) ([]model.Material, error) {
	response, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	var materials []model.Material
	err = json.NewDecoder(response.Body).Decode(&materials)
	if err != nil {
		return nil, err
	}

	return materials, nil
}

func CreateReport(query string) error {
	_, err := http.Get(query)
	if err != nil {
		return err
	}

	return nil
}

func GetBuildingSites(query string) ([]model.BuildingSite, error) {
	response, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	var sites []model.BuildingSite
	err = json.NewDecoder(response.Body).Decode(&sites)
	if err != nil {
		return nil, err
	}

	return sites, nil
}

func SaveWorkSchedule(query string, schedule *model.WorkSchedule) error {

	jsn, err := json.Marshal(schedule)
	if err != nil {
		return err
	}

	response, err := http.Post(query, "application/json", bytes.NewBuffer(jsn))
	if err != nil {
		return err
	}

	if response.StatusCode == 200 || response.StatusCode == 201 {
		return nil
	}

	messageBytes, _ := io.ReadAll(response.Body)
	message := string(messageBytes)

	return errors.New("failed to save work schedule: " + message)
}

func SaveConstructionTeam(query string, team *model.ConstructionTeam) error {
	jsn, err := json.Marshal(team)
	if err != nil {
		return err
	}

	response, err := http.Post(query, "application/json", bytes.NewBuffer(jsn))
	if err != nil {
		return err
	}

	if response.StatusCode == 200 || response.StatusCode == 201 {
		return nil
	}

	messageBytes, _ := io.ReadAll(response.Body)
	message := string(messageBytes)

	return errors.New("failed to save construction team: " + message)
}

func GetSchedule(query string) (model.WorkSchedule, error) {
	response, err := http.Get(query)
	if err != nil {
		return model.WorkSchedule{ID: 0}, err
	}

	var schedule model.WorkSchedule
	err = json.NewDecoder(response.Body).Decode(&schedule)
	if err != nil {
		return model.WorkSchedule{ID: 0}, err
	}

	return schedule, nil
}

func DeleteSchedule(query string) error {
	req, err := http.NewRequest("DELETE", query, nil)
	if err != nil {
		return err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if response.StatusCode == http.StatusOK {
		return nil
	}

	messageBytes, _ := io.ReadAll(response.Body)
	message := string(messageBytes)

	return errors.New("failed to delete schedule: " + message)
}

func DeleteConstructionTeam(query string) error {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", query, nil)
	if err != nil {
		return err
	}

	response, err := client.Do(req)
	if err != nil {
		return err
	}

	if response.StatusCode == http.StatusOK {
		return nil
	}

	messageBytes, _ := io.ReadAll(response.Body)
	message := string(messageBytes)

	return errors.New("failed to delete construction team: " + message)
}

func UpdateSchedule(query string, schedule *model.WorkSchedule) error {
	client := &http.Client{}
	jsn, err := json.Marshal(schedule)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, query, bytes.NewBuffer(jsn))
	if err != nil {
		return err
	}

	response, err := client.Do(req)
	if err != nil {
		return err
	}

	if response.StatusCode == 200 || response.StatusCode == 201 {
		return nil
	}

	messageBytes, _ := io.ReadAll(response.Body)
	message := string(messageBytes)

	return errors.New("failed to update schedule: " + message)
}

func GetEngineerTeams(query string) ([]model.EngineerTeam, error) {

	response, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	var teams []model.EngineerTeam
	err = json.NewDecoder(response.Body).Decode(&teams)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func GetEngineerTeam(query string) (model.EngineerTeam, error) {

	response, err := http.Get(query)
	if err != nil {
		return model.EngineerTeam{ID: 0}, err
	}

	var team model.EngineerTeam
	err = json.NewDecoder(response.Body).Decode(&team)
	if err != nil {
		return model.EngineerTeam{ID: 0}, err
	}

	return team, nil
}

func UpdateConstructionTeam(query string, team *model.ConstructionTeam) error {

	client := &http.Client{}
	jsn, err := json.Marshal(team)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, query, bytes.NewBuffer(jsn))
	if err != nil {
		return err
	}

	response, err := client.Do(req)
	if err != nil {
		return err
	}

	if response.StatusCode == 200 || response.StatusCode == 201 {
		return nil
	}

	messageBytes, _ := io.ReadAll(response.Body)
	message := string(messageBytes)

	return errors.New("failed to update construction team: " + message)
}

func GetBuildingOrganizations(query string) ([]model.BuildingOrganization, error) {
	response, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	var organizations []model.BuildingOrganization
	err = json.NewDecoder(response.Body).Decode(&organizations)
	if err != nil {
		return nil, err
	}

	return organizations, nil
}
