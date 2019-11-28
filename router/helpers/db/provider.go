package db

import (
	"context"

	"github.com/evscott/z3-e2c-api/dal"
	"github.com/evscott/z3-e2c-api/models"
)

type Config struct {
	dal *dal.DAL
}

func Init(dal *dal.DAL) *Config {
	return &Config{
		dal: dal,
	}
}

type Provider interface {
	GetAssignmentByNameAssignment(ctx context.Context, name string) (*models.Assignment, error)
	CreateAssignment(ctx context.Context, assignmentName string) error
	UpdateAssignmentBlob(ctx context.Context, assignmentName, blobSHA string) error
	UpdateFile(ctx context.Context, assignmentName, dropboxName, fileName, commitID string) error
	CreateFile(ctx context.Context, fileName, assignmentName, dropboxName, commitID string) error
	CreateDropbox(ctx context.Context, dropboxName, assignmentName string) error
	GetDropboxByNameAndAssignment(ctx context.Context, dropboxName, assignmentName string) (*models.Dropbox, error)
	CreateSubmission(ctx context.Context, dropboxName, assignmentName string, prNumber int) error
	GetSubmission(ctx context.Context, dropboxName, assignmentName string) (*models.Submission, error)
	GetFile(ctx context.Context, dropboxName, assignmentName, fileName string) (*models.File, error)
}
