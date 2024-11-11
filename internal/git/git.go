package git

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type RepositoryTagInfo struct {
	Name *string
}

type Git interface {
	GetRepo(owner string, name string) *github.Repository
	GetRepoTags(owner string, name string) []RepositoryTagInfo
}

type GitClient struct {
	api *github.Client
}

func NewClient(token string) *GitClient {
	sts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	oauthClient := oauth2.NewClient(context.Background(), sts)
	api := github.NewClient(oauthClient)

	return &GitClient{
		api: api,
	}
}

func (g GitClient) GetRepo(owner string, name string) *github.Repository {
	repo, _, err := g.api.Repositories.Get(context.Background(), owner, name)
	if err != nil {
		panic(err)
	}
	return repo
}

func (g GitClient) GetRepoTags(owner string, name string) []RepositoryTagInfo {
	result := []RepositoryTagInfo{}

	releases, _, err := g.api.Repositories.ListReleases(context.Background(), owner, name, nil)
	if err != nil {
		panic(err)
	}

	if len(releases) == 0 {
		tags, _, err := g.api.Repositories.ListTags(context.Background(), owner, name, nil)
		if err != nil {
			panic(err)
		}

		for _, v := range tags {
			result = append(result, RepositoryTagInfo{Name: v.Name})
		}
	} else {
		for _, v := range releases {
			result = append(result, RepositoryTagInfo{Name: v.Name})
		}
	}

	return result
}
