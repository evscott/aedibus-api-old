package provider

import (
	"context"
	"github.com/evscott/z3-e2c-api/models"
	"github.com/evscott/z3-e2c-api/shared/logger"
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
	GetAssignment(ctx context.Context, assignment *models.Assignment) error
	// Files
	CreateFile(ctx context.Context, file *models.File) error
	// Submissions
	CreateSubmission(ctx context.Context, submission *models.Submission) error
	GetSubmissionByBranchAndRepo(ctx context.Context, submission *models.Submission) error
}
