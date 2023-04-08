package models

import "time"

type PendingDeployments []struct {
	Environment struct {
		ID      int    `json:"id"`
		NodeID  string `json:"node_id"`
		Name    string `json:"name"`
		URL     string `json:"url"`
		HTMLURL string `json:"html_url"`
	} `json:"environment"`
	WaitTimer             int        `json:"wait_timer"`
	WaitTimerStartedAt    time.Time  `json:"wait_timer_started_at"`
	CurrentUserCanApprove bool       `json:"current_user_can_approve"`
	Reviewers             []Reviewer `json:"reviewers"`
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
