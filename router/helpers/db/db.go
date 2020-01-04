package db

import (
	"context"
	"github.com/evscott/aedibus-api/models"
)

func (c *Config) GetAssignments(ctx context.Context) (models.Assignments, error) {
	assignments, err := c.dal.Provider.GetAssignments(ctx)
	if err != nil {
		return nil, err
	}

	return assignments, nil
}

// TODO
//
func (c *Config) GetAssignmentByName(ctx context.Context, name string) (*models.Assignment, error) {
	assignment := &models.Assignment{
		Name: name,
	}
	if err := c.dal.Provider.GetAssignmentByName(ctx, assignment); err != nil {
		return nil, err
	}

	return assignment, nil
}

// TODO
//
func (c *Config) CreateAssignment(ctx context.Context, assignmentName string) (*models.Assignment, error) {
	assignment := &models.Assignment{
		Name: assignmentName,
	}
	if err := c.dal.Provider.CreateAssignment(ctx, assignment); err != nil {
		return nil, err
	}
	return assignment, nil
}

// TODO
//
func (c *Config) UpdateAssignmentBlob(ctx context.Context, aid, blobSHA string) error {
	assignment := &models.Assignment{
		ID:      aid,
		BlobSHA: blobSHA,
	}
	return c.dal.Provider.UpdateAssignment(ctx, assignment)
}

// TODO
//
func (c *Config) UpdateFile(ctx context.Context, assignmentName, dropboxName, fileName, commitID string) error {
	file := &models.File{
		Name:     fileName,
		AID:      assignmentName,
		DID:      dropboxName,
		CommitID: commitID,
	}
	return c.dal.Provider.UpdateFile(ctx, file)
}

// TODO
//
func (c *Config) CreateFile(ctx context.Context, fileName, aid, did, commitID string) error {
	file := &models.File{
		Name:     fileName,
		AID:      aid,
		DID:      did,
		CommitID: commitID,
	}
	return c.dal.Provider.CreateFile(ctx, file)
}

// TODO
//
func (c *Config) CreateDropbox(ctx context.Context, dropboxName, aid string) (*models.Dropbox, error) {
	dropbox := &models.Dropbox{
		Name: dropboxName,
		AID:  aid,
	}
	return dropbox, c.dal.Provider.CreateDropbox(ctx, dropbox)
}

// TODO
//
func (c *Config) GetDropbox(ctx context.Context, aid, dropboxName string) (*models.Dropbox, error) {
	dropbox := &models.Dropbox{
		Name: dropboxName,
		AID:  aid,
	}
	return dropbox, c.dal.Provider.GetDropboxByNameAndAssignment(ctx, dropbox)
}

// TODO
func (c *Config) GetDropboxes(ctx context.Context, aid string) (models.Dropboxes, error) {
	return c.dal.Provider.GetDropboxes(ctx, aid)
}

// TODO
//
func (c *Config) CreateSubmission(ctx context.Context, did, aid string, prNumber int) error {
	submission := &models.Submission{
		AID:      aid,
		DID:      did,
		PrNumber: prNumber,
	}
	return c.dal.Provider.CreateSubmission(ctx, submission)
}

// TODO
//
func (c *Config) GetSubmission(ctx context.Context, did, aid string) (*models.Submission, error) {
	submission := &models.Submission{
		AID: aid,
		DID: did,
	}
	return submission, c.dal.Provider.GetSubmission(ctx, submission)
}

// TODO
//
func (c *Config) GetFile(ctx context.Context, did, aid, fileName string) (*models.File, error) {
	file := &models.File{
		Name: fileName,
		AID:  aid,
		DID:  did,
	}
	return file, c.dal.Provider.GetFile(ctx, file)
}

// TODO
//
func (c *Config) DeleteAssignment(ctx context.Context, assignment *models.Assignment) error {
	return c.dal.Provider.DeleteAssignmentTx(ctx, assignment)
}
