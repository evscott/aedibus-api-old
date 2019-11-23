package provider

import (
	"context"
	"github.com/evscott/z3-e2c-api/models"
)

func (c *Config) CreateDropbox(ctx context.Context, dropbox *models.Dropbox) error {
	return c.db.Insert(dropbox)
}

func (c *Config) GetDropboxByNameAndAssignment(ctx context.Context, dropbox *models.Dropbox) error {
	return c.db.Model(dropbox).
		Where("name = ?", dropbox.Name).
		Where("assignment_name = ?", dropbox.AssignmentName).
		Select()
}

func (c *Config) MarkSubmitted(ctx context.Context, dropbox *models.Dropbox) error {
	_, err := c.db.
		Model(dropbox).
		WherePK().
		Set("submitted = ?", true).
		Update()
	return err
}
