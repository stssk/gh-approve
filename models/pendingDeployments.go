package models

import (
	"fmt"
	"time"
)

func PendingDeploymentsUrl(owner string, repo string, runId int) string {
	return fmt.Sprintf("repos/%s/%s/actions/runs/%d/pending_deployments", owner, repo, runId)
}

type PendingDeployments []struct {
	Environment           Environment `json:"environment"`
	WaitTimer             int         `json:"wait_timer"`
	WaitTimerStartedAt    time.Time   `json:"wait_timer_started_at"`
	CurrentUserCanApprove bool        `json:"current_user_can_approve"`
	Reviewers             []Reviewer  `json:"reviewers"`
}

type Environment struct {
	ID      int    `json:"id"`
	NodeID  string `json:"node_id"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}

type Reviewer struct {
	Type     string `json:"type"`
	Reviewer struct {
		ID      int    `json:"id"`
		NodeID  string `json:"node_id"`
		URL     string `json:"url"`
		HTMLURL string `json:"html_url"`
	} `json:"reviewer"`
}

type ReviewUser struct {
	Type     string `json:"type"`
	Reviewer struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"reviewer"`
}

type ReviewTeam struct {
	Type     string `json:"type"`
	Reviewer struct {
		ID              int    `json:"id"`
		NodeID          string `json:"node_id"`
		URL             string `json:"url"`
		HTMLURL         string `json:"html_url"`
		Name            string `json:"name"`
		Slug            string `json:"slug"`
		Description     string `json:"description"`
		Privacy         string `json:"privacy"`
		Permission      string `json:"permission"`
		MembersURL      string `json:"members_url"`
		RepositoriesURL string `json:"repositories_url"`
		Parent          any    `json:"parent"`
	} `json:"reviewer"`
}

type RequestApprovement struct {
	EnvironmentIds []int  `json:"environment_ids"`
	State          States `json:"state"`
	Comment        string `json:"comment"`
}

type States string

const (
	Approved States = "approved"
	Rejected States = "rejected"
)

type DeploymentResponse []struct {
	URL                   string    `json:"url"`
	ID                    int       `json:"id"`
	NodeID                string    `json:"node_id"`
	Sha                   string    `json:"sha"`
	Ref                   string    `json:"ref"`
	Task                  string    `json:"task"`
	Payload               Payload   `json:"payload"`
	OriginalEnvironment   string    `json:"original_environment"`
	Environment           string    `json:"environment"`
	Description           string    `json:"description"`
	Creator               Creator   `json:"creator"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	StatusesURL           string    `json:"statuses_url"`
	RepositoryURL         string    `json:"repository_url"`
	TransientEnvironment  bool      `json:"transient_environment"`
	ProductionEnvironment bool      `json:"production_environment"`
}
type Payload struct {
}
type Creator struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}
