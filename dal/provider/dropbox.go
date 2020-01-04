package provider

import (
	"context"
	"github.com/evscott/aedibus-api/models"
)

func (c *Config) CreateDropbox(ctx context.Context, dropbox *models.Dropbox) error {
	return c.db.Insert(dropbox)
}

func (c *Config) GetDropboxes(ctx context.Context, aid string) (models.Dropboxes, error) {
	var dropboxes models.Dropboxes
	return dropboxes, c.db.Model(&dropboxes).Where("aid = ?", aid).Select()
}

func (c *Config) GetDropboxByNameAndAssignment(ctx context.Context, dropbox *models.Dropbox) error {
	return c.db.Model(dropbox).
		Where("name = ?", dropbox.Name).
		Where("assignment_name = ?", dropbox.AID).
		Select()
}

func (c *Config) MarkSubmitted(ctx context.Context, dropbox *models.Dropbox) error {
	_, err := c.db.
		Model(dropbox).
		WherePK().
		Set("submitted = ?", true).
		UpdateNotZero()
	return err
}
