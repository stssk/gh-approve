package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/stssk/gh-approve/models"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cli/go-gh"
)

func main() {
	client, err := gh.RESTClient(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// currentUser := models.User{}
	// err = client.Get(models.UserUrl(), &currentUser)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	currentRepo, err := gh.CurrentRepository()
	if err != nil {
		fmt.Println(err)
		return
	}

	runs := models.Runs{}
	err = client.Get(models.RunsUrl(currentRepo.Owner(), currentRepo.Name()), &runs)
	if err != nil {
		fmt.Println(err)
		return
	}
	if runs.TotalCount == 0 {
		fmt.Println("No runs detected")
		return
	}
	pendingRuns := make([]models.WorkflowRun, 0)
	pendingRunsTexts := make([]string, 0)
	for _, run := range runs.WorkflowRuns {
		if run.Status == "waiting" {
			pendingRuns = append(pendingRuns, run)
			pendingRunsTexts = append(pendingRunsTexts, fmt.Sprintf("%s, %s (%s)", run.HeadCommit.Message, run.Name, run.HeadBranch))
			// } else {
			// 	fmt.Printf("%s: %s", run.Status, run.HeadCommit.Message)
			// fmt.Println("🎉 No waiting workflow runs found")
			// return
		}
	}
	// var selectedRun int
	promptRuns := &survey.Select{
		Message: "Select a workflow run",
		Options: pendingRunsTexts,
	}

	answer := -1
	survey.AskOne(promptRuns, &answer, survey.WithValidator(survey.Required))
	if answer < 0 {
		return
	}

	pendingDeployments := models.PendingDeployments{}
	err = client.Get(models.PendingDeploymentsUrl(currentRepo.Owner(), currentRepo.Name(), pendingRuns[answer].ID), &pendingDeployments)
	if err != nil {
		fmt.Println(err)
		return
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
		return
	}

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
