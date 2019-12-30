package provider

import (
	"context"
	"github.com/evscott/aedibus-api/models"
)

func (c *Config) CreateAssignment(ctx context.Context, assignment *models.Assignment) error {
	return c.db.Insert(assignment)
}

func (c *Config) UpdateAssignmentByName(ctx context.Context, assignment *models.Assignment) error {
	_, err := c.db.Model(assignment).
		Where("name = ?", assignment.Name).
		UpdateNotZero()
	return err
}

func (c *Config) GetAnAssignment(ctx context.Context, assignment *models.Assignment) error {
	return c.db.Model(assignment).
		WherePK().
		Select()
}

func (c *Config) GetAssignments(ctx context.Context) (models.Assignments, error) {
	var assignments models.Assignments
	return assignments, c.db.Model(&assignments).
		Select()
}
