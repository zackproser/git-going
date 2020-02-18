package github

import (
	"context"
	"os"

	"github.com/google/go-github/v29/github"
	"golang.org/x/oauth2"
)

var (
	// GithubAPIToken must be a valid GitHub Personal access token
	GithubAPIToken string
)

func init() {
	val, ok := os.LookupEnv("GIT_GOING_GITHUB_TOKEN")
	if !ok {
		panic("Must set GIT_GOING_GITHUB_TOKEN with valid GitHub API token")
	}
	GithubAPIToken = val
}

// CreateRepo makes an API into GitHub to create a new
// repository using the project slug
func CreateRepo(name string) (*github.Repository, error) {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GithubAPIToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repo := &github.Repository{
		Name:        github.String(name),
		Description: github.String(name),
	}

	repository, _, err := client.Repositories.Create(ctx, "", repo)

	if err != nil {
		return nil, err
	}

	return repository, nil
}
