package provider

import (
	"context"
	"github.com/evscott/z3-e2c-api/models"
)

func (c *Config) CreateSubmission(ctx context.Context, submission *models.Submission) error {
	return c.db.Insert(submission)
}

func (c *Config) GetSubmissionByBranchAndRepo(ctx context.Context, submission *models.Submission) error {
	return c.db.Model(submission).
		Where("branch = ?", *submission.Branch).
		Select()
}
