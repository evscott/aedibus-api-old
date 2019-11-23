package db

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/evscott/z3-e2c-api/dal"
	"github.com/evscott/z3-e2c-api/models"
)

type Config struct {
	dal *dal.DAL
}

func Init(dal *dal.DAL) *Config {
	return &Config{
		dal: dal,
	}
}

// TODO
//
func (c *Config) GetAssignment(ctx context.Context, name string) (*models.Assignment, error) {
	assignment := &models.Assignment{
		Name: name,
	}
	if err := c.dal.Provider.GetAssignment(ctx, assignment); err != nil {
		return nil, err
	}

	return assignment, nil
}

// TODO
//
func (c *Config) CreateAssignment(ctx context.Context, assignmentName string) error {
	assignment := &models.Assignment{
		Name: assignmentName,
	}
	if err := c.dal.Provider.CreateAssignment(ctx, assignment); err != nil {
		return err
	}
	return nil
}

// TODO
//
func (c *Config) GetFileFromForm(w http.ResponseWriter, r *http.Request, fileName string) ([]byte, error) {
	file, _, err := r.FormFile(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read file contents through buffer
	buffer := bytes.Buffer{}
	if _, err := io.Copy(&buffer, file); err != nil {
		return nil, err
	}
	defer buffer.Reset()
	return buffer.Bytes(), nil
}

// TODO
//
func (c *Config) UpdateAssignmentBlob(ctx context.Context, assignmentName, blobSHA string) error {
	assignment := &models.Assignment{
		Name:    assignmentName,
		BlobSHA: blobSHA,
	}
	return c.dal.Provider.UpdateAssignment(ctx, assignment)
}

// TODO
//
func (c *Config) CreateFile(ctx context.Context, fileName, assignmentName, dropboxName string) error {
	file := &models.File{
		Name:           fileName,
		AssignmentName: assignmentName,
		DropboxName:    dropboxName,
	}
	return c.dal.Provider.CreateFile(ctx, file)
}

// TODO
//
func (c *Config) CreateDropbox(ctx context.Context, dropboxName, assignmentName string) error {
	dropbox := &models.Dropbox{
		Name:           dropboxName,
		AssignmentName: assignmentName,
	}
	return c.dal.Provider.CreateDropbox(ctx, dropbox)
}

// TODO
//
func (c *Config) GetDropboxByNameAndAssignment(ctx context.Context, dropboxName, assignmentName string) (*models.Dropbox, error) {
	dropbox := &models.Dropbox{
		Name:           dropboxName,
		AssignmentName: assignmentName,
	}
	return dropbox, c.dal.Provider.GetDropboxByNameAndAssignment(ctx, dropbox)
}
