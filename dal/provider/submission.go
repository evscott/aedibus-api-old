package provider

import (
	"context"

	"github.com/evscott/aedibus-api/models"
)

func (c *Config) CreateSubmission(ctx context.Context, submission *models.Submission) error {
	return c.db.Insert(submission)
}

func (c *Config) GetSubmission(ctx context.Context, submission *models.Submission) error {
	return c.db.Model(submission).
		WherePK().
		Select()
}
