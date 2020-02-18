package github

import (
	"context"
	"os"

	"github.com/google/go-github/v29/github"
	"golang.org/x/oauth2"
)

var (
	GITHUB_API_TOKEN string
	client           *github.Client
)

func init() {
	val, ok := os.LookupEnv("GIT_GOING_GITHUB_TOKEN")
	if !ok {
		panic("Must set GIT_GOING_GITHUB_TOKEN with valid GitHub API token")
	}
	GITHUB_API_TOKEN = val
}

func CreateRepo(name string) (*github.Repository, error) {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GITHUB_API_TOKEN},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repo := &github.Repository{
		ID:                  nil,
		NodeID:              nil,
		Owner:               nil,
		Name:                github.String(name),
		FullName:            nil,
		Description:         github.String(name),
		Homepage:            nil,
		CodeOfConduct:       nil,
		DefaultBranch:       nil,
		MasterBranch:        nil,
		CreatedAt:           nil,
		PushedAt:            nil,
		UpdatedAt:           nil,
		HTMLURL:             nil,
		CloneURL:            nil,
		GitURL:              nil,
		MirrorURL:           nil,
		SSHURL:              nil,
		SVNURL:              nil,
		Language:            nil,
		Fork:                nil,
		ForksCount:          nil,
		NetworkCount:        nil,
		OpenIssuesCount:     nil,
		StargazersCount:     nil,
		SubscribersCount:    nil,
		WatchersCount:       nil,
		Size:                nil,
		AutoInit:            nil,
		Parent:              nil,
		Source:              nil,
		TemplateRepository:  nil,
		Organization:        nil,
		Permissions:         nil,
		AllowRebaseMerge:    nil,
		AllowSquashMerge:    nil,
		AllowMergeCommit:    nil,
		DeleteBranchOnMerge: nil,
		Topics:              nil,
		Archived:            nil,
		Disabled:            nil,
		License:             nil,
		Private:             nil,
		HasIssues:           nil,
		HasWiki:             nil,
		HasPages:            nil,
		HasProjects:         nil,
		HasDownloads:        nil,
		IsTemplate:          nil,
		LicenseTemplate:     nil,
		GitignoreTemplate:   nil,
		TeamID:              nil,
		URL:                 nil,
		ArchiveURL:          nil,
		AssigneesURL:        nil,
		BlobsURL:            nil,
		BranchesURL:         nil,
		CollaboratorsURL:    nil,
		CommentsURL:         nil,
		CommitsURL:          nil,
		CompareURL:          nil,
		ContentsURL:         nil,
		ContributorsURL:     nil,
		DeploymentsURL:      nil,
		DownloadsURL:        nil,
		EventsURL:           nil,
		ForksURL:            nil,
		GitCommitsURL:       nil,
		GitRefsURL:          nil,
		GitTagsURL:          nil,
		HooksURL:            nil,
		IssueCommentURL:     nil,
		IssueEventsURL:      nil,
		IssuesURL:           nil,
		KeysURL:             nil,
		LabelsURL:           nil,
		LanguagesURL:        nil,
		MergesURL:           nil,
		MilestonesURL:       nil,
		NotificationsURL:    nil,
		PullsURL:            nil,
		ReleasesURL:         nil,
		StargazersURL:       nil,
		StatusesURL:         nil,
		SubscribersURL:      nil,
		SubscriptionURL:     nil,
		TagsURL:             nil,
		TreesURL:            nil,
		TeamsURL:            nil,
		TextMatches:         nil,
	}

	repository, _, err := client.Repositories.Create(ctx, "", repo)

	if err != nil {
		return nil, err
	}

	return repository, nil
}
