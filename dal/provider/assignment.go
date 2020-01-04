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
		UpdateNotZero()
	return err
}

func (c *Config) GetAssignmentByName(ctx context.Context, assignment *models.Assignment) error {
	return c.db.Model(assignment).
		Where("name = ?", assignment.Name).
		Select()
}

func (c *Config) GetAssignments(ctx context.Context) (models.Assignments, error) {
	var assignments models.Assignments
	return assignments, c.db.Model(&assignments).
		Select()
}

func (c *Config) DeleteAssignmentTx(ctx context.Context, assignment *models.Assignment) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	// delete all File records
	_, err = tx.Model(&models.File{}).
		Where("aid = ?", assignment.ID).
		Delete()
	if err != nil {
		return err
	}

	// delete all Dropbox records
	_, err = tx.Model(&models.Dropbox{}).
		Where("aid = ?", assignment.ID).
		Delete()
	if err != nil {
		return err
	}

	// delete all Submission records
	_, err = tx.Model(&models.Submission{}).
		Where("aid = ?", assignment.ID).
		Delete()
	if err != nil {
		return err
	}

	// delete assignment record
	_, err = tx.Model(assignment).
		WherePK().
		Delete()

	if err := tx.Commit(); err != nil {
		return tx.Rollback()
	}

	return nil
}
