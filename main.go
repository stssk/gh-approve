package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/stssk/gh-approve/models"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cli/go-gh"
	"github.com/cli/go-gh/pkg/api"
	"github.com/cli/go-gh/pkg/repository"
)

func main() {
	client, err := gh.RESTClient(nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	currentRepo, err := gh.CurrentRepository()
	if err != nil {
		fmt.Println(err)
		return
	}

	pendingRuns, answer, shouldReturn := selectPendingRuns(client, currentRepo)
	if shouldReturn {
		return
	}

	pendingDeployments, approvedEnvironments, shouldReturn := selectPendingDeployments(client, currentRepo, pendingRuns, answer)
	if shouldReturn {
		return
	}

	approveDeployments(approvedEnvironments, pendingDeployments, client, currentRepo, pendingRuns, answer)
}

// Fetch pending runs and prompt the user to select one of them
func selectPendingRuns(client api.RESTClient, currentRepo repository.Repository) ([]models.WorkflowRun, int, bool) {
	runs := models.Runs{}
	err := client.Get(models.RunsUrl(currentRepo.Owner(), currentRepo.Name()), &runs)
	if err != nil {
		fmt.Println(err)
		return nil, 0, true
	}
	if runs.TotalCount == 0 {
		fmt.Println("No runs detected")
		return nil, 0, true
	}
	pendingRuns := make([]models.WorkflowRun, 0)
	pendingRunsTexts := make([]string, 0)
	for _, run := range runs.WorkflowRuns {
		if run.Status == "waiting" {
			pendingRuns = append(pendingRuns, run)
			pendingRunsTexts = append(pendingRunsTexts, fmt.Sprintf("%s, %s (%s)", run.HeadCommit.Message, run.Name, run.HeadBranch))
		}
	}
	promptRuns := &survey.Select{
		Message: "Select a workflow run",
		Options: pendingRunsTexts,
	}

	answer := -1
	survey.AskOne(promptRuns, &answer, survey.WithValidator(survey.Required))
	if answer < 0 {
		return nil, 0, true
	}
	return pendingRuns, answer, false
}

// Fetch pending deployments for a single run and prompt the user to select 0, 1 or more of them
func selectPendingDeployments(client api.RESTClient, currentRepo repository.Repository, pendingRuns []models.WorkflowRun, answer int) (models.PendingDeployments, []int, bool) {
	pendingDeployments := models.PendingDeployments{}
	err := client.Get(models.PendingDeploymentsUrl(currentRepo.Owner(), currentRepo.Name(), pendingRuns[answer].ID), &pendingDeployments)
	if err != nil {
		fmt.Println(err)
		return nil, nil, true
	}

	pendingDeploymentTexts := make([]string, len(pendingDeployments))
	for i, d := range pendingDeployments {
		pendingDeploymentTexts[i] = d.Environment.Name
	}

	approvedEnvironments := []int{}
	promptDeployments := &survey.MultiSelect{
		Message: "Select environments to approve",
		Options: pendingDeploymentTexts,
	}
	survey.AskOne(promptDeployments, &approvedEnvironments)

	if len(approvedEnvironments) == 0 {
		return nil, nil, true
	}
	return pendingDeployments, approvedEnvironments, false
}

// Send the approval request to GitHub actions
func approveDeployments(approvedEnvironments []int, pendingDeployments models.PendingDeployments, client api.RESTClient, currentRepo repository.Repository, pendingRuns []models.WorkflowRun, answer int) {
	approveIds := make([]int, len(approvedEnvironments))
	for i, e := range approvedEnvironments {
		approveIds[i] = pendingDeployments[e].Environment.ID
	}

	approveRequest := models.RequestApprovement{
		EnvironmentIds: approveIds,
		State:          models.Approved,
	}
	req, err := json.Marshal(approveRequest)
	if err != nil {
		fmt.Println(err)
		return
	}
	approvedDeployments := models.DeploymentResponse{}
	err = client.Post(models.PendingDeploymentsUrl(currentRepo.Owner(), currentRepo.Name(), pendingRuns[answer].ID), bytes.NewReader(req), &approvedDeployments)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, e := range approvedDeployments {
		fmt.Printf(" • %s approved @%s\n", e.Environment, e.CreatedAt)
	}
}
