package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/stssk/gh-approve/models"

	"charm.land/huh/v2"
	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/cli/go-gh/v2/pkg/repository"
)

func main() {
	client, err := api.DefaultRESTClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	currentRepo, err := repository.Current()
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
func selectPendingRuns(client *api.RESTClient, currentRepo repository.Repository) ([]models.WorkflowRun, int, bool) {
	runs := models.Runs{}
	err := client.Get(models.RunsUrl(currentRepo.Owner, currentRepo.Name)+"?status=waiting", &runs)
	if err != nil {
		fmt.Println(err)
		return nil, 0, true
	}
	if runs.TotalCount == 0 {
		fmt.Println("No runs detected")
		return nil, 0, true
	}
	pendingRuns := make([]models.WorkflowRun, 0)
	runOptions := make([]huh.Option[int], 0)
	now := time.Now()
	for i, run := range runs.WorkflowRuns {
		pendingRuns = append(pendingRuns, run)
		runOptions = append(runOptions, huh.NewOption(fmt.Sprintf("%s, %s (%s) %s ago", run.DisplayTitle, run.Name, run.HeadBranch, now.Sub(run.RunStartedAt).Round(time.Second)), i))
	}

	answer := -1
	err = huh.NewSelect[int]().
		Title("Select a workflow run").
		Options(runOptions...).
		Value(&answer).
		Run()
	if err != nil || answer < 0 {
		return nil, 0, true
	}
	return pendingRuns, answer, false
}

// Fetch pending deployments for a single run and prompt the user to select 0, 1 or more of them
func selectPendingDeployments(client *api.RESTClient, currentRepo repository.Repository, pendingRuns []models.WorkflowRun, answer int) (models.PendingDeployments, []int, bool) {
	pendingDeployments := models.PendingDeployments{}
	err := client.Get(models.PendingDeploymentsUrl(currentRepo.Owner, currentRepo.Name, pendingRuns[answer].ID), &pendingDeployments)
	if err != nil {
		fmt.Println(err)
		return nil, nil, true
	}

	deploymentOptions := make([]huh.Option[int], len(pendingDeployments))
	for i, d := range pendingDeployments {
		deploymentOptions[i] = huh.NewOption(d.Environment.Name, i)
	}

	approvedEnvironments := []int{}
	err = huh.NewMultiSelect[int]().
		Title("Select environments to approve").
		Options(deploymentOptions...).
		Value(&approvedEnvironments).
		Run()
	if err != nil || len(approvedEnvironments) == 0 {
		return nil, nil, true
	}
	return pendingDeployments, approvedEnvironments, false
}

// Send the approval request to GitHub actions
func approveDeployments(approvedEnvironments []int, pendingDeployments models.PendingDeployments, client *api.RESTClient, currentRepo repository.Repository, pendingRuns []models.WorkflowRun, answer int) {
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
	err = client.Post(models.PendingDeploymentsUrl(currentRepo.Owner, currentRepo.Name, pendingRuns[answer].ID), bytes.NewReader(req), &approvedDeployments)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, e := range approvedDeployments {
		fmt.Printf(" • %s approved @%s\n", e.Environment, e.CreatedAt)
	}
}
