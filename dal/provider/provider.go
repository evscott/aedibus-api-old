package provider

import (
	"context"
	"github.com/evscott/aedibus-api/models"
	"github.com/evscott/aedibus-api/shared/logger"
	"github.com/go-pg/pg/v9"
)

type Config struct {
	logger *logger.StandardLogger
	db     *pg.DB
}

func Init(logger *logger.StandardLogger, db *pg.DB) *Config {
	return &Config{
		logger: logger,
		db:     db,
	}
}

type Provider interface {
	// Assignments
	CreateAssignment(ctx context.Context, assignment *models.Assignment) error
	UpdateAssignment(ctx context.Context, assignment *models.Assignment) error
	GetAssignmentByName(ctx context.Context, assignment *models.Assignment) error
	GetAssignments(ctx context.Context) (models.Assignments, error)
	DeleteAssignmentTx(ctx context.Context, assignment *models.Assignment) error
	// Files
	CreateFile(ctx context.Context, file *models.File) error
	UpdateFile(ctx context.Context, file *models.File) error
	// Dropboxes
	CreateDropbox(ctx context.Context, dropbox *models.Dropbox) error
	GetDropboxByNameAndAssignment(ctx context.Context, submission *models.Dropbox) error
	// Submissions
	CreateSubmission(ctx context.Context, submission *models.Submission) error
	GetSubmission(ctx context.Context, submission *models.Submission) error
	GetFile(ctx context.Context, file *models.File) error
}
