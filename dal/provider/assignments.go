package provider

import (
	"context"

	"github.com/evscott/z3-e2c-api/models"
)

func (c *Config) CreateAssignment(ctx context.Context, assignment *models.Assignment) error {
	return c.db.Insert(assignment)
}

func (c *Config) UpdateAssignment(ctx context.Context, assignment *models.Assignment) error {
	_, err := c.db.Model(assignment).WherePK().Update()
	return err
}
