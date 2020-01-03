package gh

import (
	"context"
	"github.com/evscott/aedibus-api/models"
	"github.com/google/go-github/github"
)

type Config struct {
	gal *github.Client
}

func Init(gal *github.Client) *Config {
	return &Config{
		gal: gal,
	}
}

type Provider interface {
	CreateComment(ctx context.Context, fileName, assignmentName, commitID, body string, pullRequestNumber, position int) (*github.PullRequestComment, error)
	CreatePullRequest(ctx context.Context, dropboxName, assignmentName, title, body string) (*github.PullRequest, error)
	CreateRepository(ctx context.Context, assignmentName string) error
	DeleteRepository(ctx context.Context, assignmentName string) error
	CreateFile(ctx context.Context, assignmentName, dropboxName, fileName string, contents []byte) (*github.RepositoryContentResponse, error)
	UpdateFile(ctx context.Context, assignmentName, dropboxName, fileName string, newContents []byte) (*github.RepositoryContentResponse, error)
	GetReadme(ctx context.Context, assignmentName string) (*models.ResGetFile, error)
	GetFileContents(ctx context.Context, fileName, assignmentName string) (*models.ResGetFile, error)
	CreateDropbox(ctx context.Context, dropboxName, assignmentName string) error
	GetMasterBlobSha(ctx context.Context, assignmentName string) (*string, error)
	GetPrComments(ctx context.Context, assignmentName string, prNumber int) ([]*github.PullRequestComment, error)
}
