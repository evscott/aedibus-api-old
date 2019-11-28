package db

import (
	"context"
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
func (c *Config) UpdateAssignmentBlob(ctx context.Context, assignmentName, blobSHA string) error {
	assignment := &models.Assignment{
		Name:    assignmentName,
		BlobSHA: blobSHA,
	}
	return c.dal.Provider.UpdateAssignment(ctx, assignment)
}

// TODO
//
func (c *Config) UpdateFile(ctx context.Context, assignmentName, dropboxName, fileName, commitID string) error {
	file := &models.File{
		Name:           fileName,
		AssignmentName: assignmentName,
		DropboxName:    dropboxName,
		CommitID:       commitID,
	}
	return c.dal.Provider.UpdateFile(ctx, file)
}

// TODO
//
func (c *Config) CreateFile(ctx context.Context, fileName, assignmentName, dropboxName, commitID string) error {
	file := &models.File{
		Name:           fileName,
		AssignmentName: assignmentName,
		DropboxName:    dropboxName,
		CommitID:       commitID,
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

// TODO
//
func (c *Config) CreateSubmission(ctx context.Context, dropboxName, assignmentName string, prNumber int) error {
	submission := &models.Submission{
		AssignmentName: assignmentName,
		DropboxName:    dropboxName,
		PrNumber:       prNumber,
	}
	return c.dal.Provider.CreateSubmission(ctx, submission)
}

// TODO
//
func (c *Config) GetSubmission(ctx context.Context, dropboxName, assignmentName string) (*models.Submission, error) {
	submission := &models.Submission{
		AssignmentName: assignmentName,
		DropboxName:    dropboxName,
	}
	return submission, c.dal.Provider.GetSubmission(ctx, submission)
}

// TODO
//
func (c *Config) GetFile(ctx context.Context, dropboxName, assignmentName, fileName string) (*models.File, error) {
	file := &models.File{
		Name:           fileName,
		AssignmentName: assignmentName,
		DropboxName:    dropboxName,
	}
	return file, c.dal.Provider.GetFile(ctx, file)
}
