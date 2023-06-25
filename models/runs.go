package models

import (
	"fmt"
	"time"
)

func RunsUrl(owner string, repo string) string {
	return fmt.Sprintf("repos/%s/%s/actions/runs", owner, repo)
}

type Runs struct {
	TotalCount   int           `json:"total_count"`
	WorkflowRuns []WorkflowRun `json:"workflow_runs"`
}

type WorkflowRun struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	DisplayTitle     string    `json:"display_title"`
	NodeID           string    `json:"node_id"`
	CheckSuiteID     int       `json:"check_suite_id"`
	CheckSuiteNodeID string    `json:"check_suite_node_id"`
	HeadBranch       string    `json:"head_branch"`
	HeadSha          string    `json:"head_sha"`
	RunNumber        int       `json:"run_number"`
	Event            string    `json:"event"`
	Status           string    `json:"status"`
	Conclusion       any       `json:"conclusion"`
	WorkflowID       int       `json:"workflow_id"`
	URL              string    `json:"url"`
	HTMLURL          string    `json:"html_url"`
	PullRequests     []any     `json:"pull_requests"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Actor            struct {
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
	} `json:"actor"`
	RunAttempt      int       `json:"run_attempt"`
	RunStartedAt    time.Time `json:"run_started_at"`
	TriggeringActor struct {
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
	} `json:"triggering_actor"`
	JobsURL       string `json:"jobs_url"`
	LogsURL       string `json:"logs_url"`
	CheckSuiteURL string `json:"check_suite_url"`
	ArtifactsURL  string `json:"artifacts_url"`
	CancelURL     string `json:"cancel_url"`
	RerunURL      string `json:"rerun_url"`
	WorkflowURL   string `json:"workflow_url"`
	HeadCommit    struct {
		ID        string    `json:"id"`
		TreeID    string    `json:"tree_id"`
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
		Committer struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"committer"`
	} `json:"head_commit"`
	Repository struct {
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Owner    struct {
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
		} `json:"owner"`
		Private          bool   `json:"private"`
		HTMLURL          string `json:"html_url"`
		Description      string `json:"description"`
		Fork             bool   `json:"fork"`
		URL              string `json:"url"`
		ArchiveURL       string `json:"archive_url"`
		AssigneesURL     string `json:"assignees_url"`
		BlobsURL         string `json:"blobs_url"`
		BranchesURL      string `json:"branches_url"`
		CollaboratorsURL string `json:"collaborators_url"`
		CommentsURL      string `json:"comments_url"`
		CommitsURL       string `json:"commits_url"`
		CompareURL       string `json:"compare_url"`
		ContentsURL      string `json:"contents_url"`
		ContributorsURL  string `json:"contributors_url"`
		DeploymentsURL   string `json:"deployments_url"`
		DownloadsURL     string `json:"downloads_url"`
		EventsURL        string `json:"events_url"`
		ForksURL         string `json:"forks_url"`
		GitCommitsURL    string `json:"git_commits_url"`
		GitRefsURL       string `json:"git_refs_url"`
		GitTagsURL       string `json:"git_tags_url"`
		GitURL           string `json:"git_url"`
		IssueCommentURL  string `json:"issue_comment_url"`
		IssueEventsURL   string `json:"issue_events_url"`
		IssuesURL        string `json:"issues_url"`
		KeysURL          string `json:"keys_url"`
		LabelsURL        string `json:"labels_url"`
		LanguagesURL     string `json:"languages_url"`
		MergesURL        string `json:"merges_url"`
		MilestonesURL    string `json:"milestones_url"`
		NotificationsURL string `json:"notifications_url"`
		PullsURL         string `json:"pulls_url"`
		ReleasesURL      string `json:"releases_url"`
		SSHURL           string `json:"ssh_url"`
		StargazersURL    string `json:"stargazers_url"`
		StatusesURL      string `json:"statuses_url"`
		SubscribersURL   string `json:"subscribers_url"`
		SubscriptionURL  string `json:"subscription_url"`
		TagsURL          string `json:"tags_url"`
		TeamsURL         string `json:"teams_url"`
		TreesURL         string `json:"trees_url"`
		HooksURL         string `json:"hooks_url"`
	} `json:"repository"`
	HeadRepository struct {
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
		Owner    struct {
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
		} `json:"owner"`
		HTMLURL          string `json:"html_url"`
		Description      any    `json:"description"`
		Fork             bool   `json:"fork"`
		URL              string `json:"url"`
		ForksURL         string `json:"forks_url"`
		KeysURL          string `json:"keys_url"`
		CollaboratorsURL string `json:"collaborators_url"`
		TeamsURL         string `json:"teams_url"`
		HooksURL         string `json:"hooks_url"`
		IssueEventsURL   string `json:"issue_events_url"`
		EventsURL        string `json:"events_url"`
		AssigneesURL     string `json:"assignees_url"`
		BranchesURL      string `json:"branches_url"`
		TagsURL          string `json:"tags_url"`
		BlobsURL         string `json:"blobs_url"`
		GitTagsURL       string `json:"git_tags_url"`
		GitRefsURL       string `json:"git_refs_url"`
		TreesURL         string `json:"trees_url"`
		StatusesURL      string `json:"statuses_url"`
		LanguagesURL     string `json:"languages_url"`
		StargazersURL    string `json:"stargazers_url"`
		ContributorsURL  string `json:"contributors_url"`
		SubscribersURL   string `json:"subscribers_url"`
		SubscriptionURL  string `json:"subscription_url"`
		CommitsURL       string `json:"commits_url"`
		GitCommitsURL    string `json:"git_commits_url"`
		CommentsURL      string `json:"comments_url"`
		IssueCommentURL  string `json:"issue_comment_url"`
		ContentsURL      string `json:"contents_url"`
		CompareURL       string `json:"compare_url"`
		MergesURL        string `json:"merges_url"`
		ArchiveURL       string `json:"archive_url"`
		DownloadsURL     string `json:"downloads_url"`
		IssuesURL        string `json:"issues_url"`
		PullsURL         string `json:"pulls_url"`
		MilestonesURL    string `json:"milestones_url"`
		NotificationsURL string `json:"notifications_url"`
		LabelsURL        string `json:"labels_url"`
		ReleasesURL      string `json:"releases_url"`
		DeploymentsURL   string `json:"deployments_url"`
	} `json:"head_repository"`
}
