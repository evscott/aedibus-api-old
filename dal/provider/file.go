package provider

import (
	"context"
	"github.com/evscott/z3-e2c-api/models"
)

func (c *Config) CreateFile(ctx context.Context, file *models.File) error {

	return c.db.Insert(file)
}
