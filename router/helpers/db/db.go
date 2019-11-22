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
func (c *Config) CreateFile(ctx context.Context, fileName, assignmentName, submissionName string) error {
	file := &models.File{
		Name:           fileName,
		AssignmentName: assignmentName,
		SubmissionName: submissionName,
	}
	return c.dal.Provider.CreateFile(ctx, file)
}

// TODO
//
func (c *Config) CreateSubmission(ctx context.Context, submissionName, assignmentName string) error {
	submission := &models.Submission{
		Name:           submissionName,
		AssignmentName: assignmentName,
	}
	return c.dal.Provider.CreateSubmission(ctx, submission)
}

// TODO
//
func (c *Config) GetSubmissionByNameAndAssignment(ctx context.Context, submissionName, assignmentName string) (*models.Submission, error) {
	submission := &models.Submission{
		Name:           submissionName,
		AssignmentName: assignmentName,
	}
	return submission, c.dal.Provider.GetSubmissionByBranchAndRepo(ctx, submission)
}
