package provider

import (
	"context"

	"github.com/evscott/aedibus-api/models"
)

func (c *Config) CreateAssignment(ctx context.Context, assignment *models.Assignment) error {
	return c.db.Insert(assignment)
}

func (c *Config) UpdateAssignment(ctx context.Context, assignment *models.Assignment) error {
	_, err := c.db.Model(assignment).
		WherePK().
		Update()
	return err
}

func (c *Config) GetAssignment(ctx context.Context, assignment *models.Assignment) error {
	return c.db.Model(assignment).
		WherePK().
		Select()
}
