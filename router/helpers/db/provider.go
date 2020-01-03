package db

import (
	"context"

	"github.com/evscott/aedibus-api/dal"
	"github.com/evscott/aedibus-api/models"
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
	GetAssignments(ctx context.Context) (models.Assignments, error)
	GetAssignmentByName(ctx context.Context, name string) (*models.Assignment, error)
	CreateAssignment(ctx context.Context, assignmentName string) (*models.Assignment, error)
	UpdateAssignmentBlob(ctx context.Context, aid, blobSHA string) error
	UpdateFile(ctx context.Context, assignmentName, dropboxName, fileName, commitID string) error
	CreateFile(ctx context.Context, fileName, aid, did, commitID string) error
	CreateDropbox(ctx context.Context, dropboxName, aid string) (*models.Dropbox, error)
	GetDropboxByNameAndAssignment(ctx context.Context, dropboxName, assignmentName string) (*models.Dropbox, error)
	CreateSubmission(ctx context.Context, dropboxName, assignmentName string, prNumber int) error
	GetSubmission(ctx context.Context, dropboxName, assignmentName string) (*models.Submission, error)
	GetFile(ctx context.Context, dropboxName, assignmentName, fileName string) (*models.File, error)
	DeleteAssignment(ctx context.Context, assignment *models.Assignment) error
}
