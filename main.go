package main

import (
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
		if run.Status == "" {
			pendingRuns = append(pendingRuns, run)
			pendingRunsTexts = append(pendingRunsTexts, fmt.Sprintf("%s, %s (%s)", run.HeadCommit.Message, run.WorkflowURL, run.HeadBranch))
		} else {
			fmt.Printf("%s: %s", run.Status, run.HeadCommit.Message)
		}
	}
	// var selectedRun int
	prompt := &survey.Select{
		Message: "Choose a color:",
		Options: pendingRunsTexts,
	}

	var answer int
	survey.AskOne(prompt, &answer)

	fmt.Println(pendingRuns[answer])

	// survey.Ask(prompt, &selectedRun)

	// pendingDeployments := models.PendingDeployments{}
	// err = client.Get("repos/%s/%s/actions/runs/RUN_ID/pending_deployments", currentRepo.Owner(), currentRepo.Name())

}
