package provider

import (
	"context"
	"github.com/evscott/aedibus-api/models"
)

func (c *Config) GetFile(ctx context.Context, file *models.File) error {
	return c.db.
		Model(file).
		WherePK().
		Select()
}

func (c *Config) CreateFile(ctx context.Context, file *models.File) error {
	return c.db.Insert(file)
}

func (c *Config) UpdateFile(ctx context.Context, file *models.File) error {
	_, err := c.db.Model(file).
		WherePK().
		Set("commit_id = ?", file.CommitID).
		Update()
	return err
}
