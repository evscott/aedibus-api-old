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
		Where("name = ?", submission.Name).
		Where("submission_name = ?", submission.AssignmentName).
		Select()
}

func (c *Config) MarkSubmitted(ctx context.Context, submission *models.Submission) error {
	_, err := c.db.
		Model(submission).
		WherePK().
		Set("submitted = ?", true).
		Update()
	return err
}
