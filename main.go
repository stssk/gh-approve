package main

import (
	"fmt"

	"github.com/stssk/gh-approve/models"

	"github.com/cli/go-gh"
)

func main() {
	fmt.Println("Hi world, this is the gh-approve extension!")
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
	fmt.Println(runs)
	// pendingDeployments := models.PendingDeployments{}
	// err = client.Get("repos/%s/%s/actions/runs/RUN_ID/pending_deployments", currentRepo.Owner(), currentRepo.Name())

}
