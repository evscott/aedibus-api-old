package helpers

import (
	"context"
	"github.com/evscott/z3-e2c-api/router/helpers/db"
	"github.com/evscott/z3-e2c-api/router/helpers/gh"
	"net/http"

	"github.com/evscott/z3-e2c-api/dal"
	"github.com/evscott/z3-e2c-api/models"
	"github.com/google/go-github/github"
)

type Config struct {
	DB *db.Config
	GH *gh.Config
}

func Init(dal *dal.DAL, gal *github.Client) *Config {
	return &Config{
		DB: db.Init(dal),
		GH: gh.Init(gal),
	}
}

type PostgresHelpers interface {
	GetAssignmentByName(ctx context.Context, name string) (*models.Assignment, error)
	CreateAssignmentHelper(ctx context.Context, assignmentName string) error
	ReceiveFileContentsHelper(w http.ResponseWriter, r *http.Request, fileName string) ([]byte, error)
	UpdateAssignmentBlob(ctx context.Context, assignmentName, blobSHA string) error
	CreateFile(ctx context.Context, fileName, submissionName string) error
	CreateSubmission(ctx context.Context, submissionName, assignmentName string) error
	GetSubmissionByNameAndAssignment(ctx context.Context, submissionName, assignmentName string)
}

type GithubHelpers interface {
	CreateComment(ctx context.Context, fileName, assignmentName, commitID, body string, position int) (*github.PullRequestComment, error)
	CreatePullRequest(ctx context.Context, submissionName, assignmentName, title, body string) (*github.PullRequest, error)
	CreateRepository(ctx context.Context, assignmentName string) error
	CreateFile(ctx context.Context, assignmentName, submissionName, fileName string, contents []byte) error
	GetReadme(ctx context.Context, assignmentName, submissionName string) (*models.ResGetFile, error)
	GetFileContents(ctx context.Context, assignmentName, submissionName string) (*models.ResGetFile, error)
	UpdateFile(ctx context.Context, assignmentName, submissionName, fileName string, newContents []byte) error
	CreateSubmission(ctx context.Context, assignmentName, submissionName string) error
	GetMasterBlobSha(ctx context.Context, assignmentName string) (*string, error)
}
