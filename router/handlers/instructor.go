package handlers

import (
	"context"
	"github.com/evscott/z3-e2c-api/shared/constants"
	status "github.com/evscott/z3-e2c-api/shared/http-codes"
	"net/http"

	"github.com/evscott/z3-e2c-api/models"
	"github.com/evscott/z3-e2c-api/shared/marsh"
)

// TODO
//
//
func (c *Config) CreateAssignment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	assignmentName := r.FormValue("assignmentName")
	instructions := r.FormValue("instructions")
	testRunner := r.FormValue("testRunner")

	instructionsContents, err := c.helpers.DB.GetFileFromForm(w, r, instructions)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
	testRunnerContents, err := c.helpers.DB.GetFileFromForm(w, r, testRunner)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	req := &models.ReqCreateAssignment{
		AssignmentName:       assignmentName,
		InstructionsContents: instructionsContents,
		TestRunnerContents:   testRunnerContents,
	}

	if err := c.helpers.GH.CreateRepository(ctx, req.AssignmentName); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.CreateAssignment(ctx, req.AssignmentName); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.CreateDropbox(ctx, constants.MASTER, req.AssignmentName); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	resInstructions, err := c.helpers.GH.CreateFile(ctx, assignmentName, constants.MASTER, instructions, instructionsContents)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	resTestRunner, err := c.helpers.GH.CreateFile(ctx, assignmentName, constants.MASTER, testRunner, testRunnerContents)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.CreateFile(ctx, instructions, assignmentName, constants.MASTER, *resInstructions.Commit.SHA); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.CreateFile(ctx, testRunner, assignmentName, constants.MASTER, *resTestRunner.Commit.SHA); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	blobSHA, err := c.helpers.GH.GetMasterBlobSha(ctx, assignmentName)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.UpdateAssignmentBlob(ctx, assignmentName, *blobSHA); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}

// TODO
//
//
func (c *Config) CreateAssignmentFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	assignmentName := r.FormValue("assignmentName")
	dropboxName := r.FormValue("dropboxName")
	fileName := r.FormValue("fileName")

	contents, err := c.helpers.DB.GetFileFromForm(w, r, fileName)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	res, err := c.helpers.GH.CreateFile(ctx, assignmentName, dropboxName, fileName, contents)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.CreateFile(ctx, fileName, assignmentName, constants.MASTER, *res.Commit.SHA); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	blobSHA, err := c.helpers.GH.GetMasterBlobSha(ctx, assignmentName)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.UpdateAssignmentBlob(ctx, assignmentName, *blobSHA); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}

// TODO
//
//
func (c *Config) UpdateAssignmentFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	assignmentName := r.FormValue("assignmentName")
	dropboxName := r.FormValue("dropboxName")
	fileName := r.FormValue("fileName")

	contents, err := c.helpers.DB.GetFileFromForm(w, r, fileName)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	res, err := c.helpers.GH.UpdateFile(ctx, assignmentName, dropboxName, fileName, contents)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.UpdateFile(ctx, assignmentName, dropboxName, fileName, *res.Commit.SHA); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	blobSHA, err := c.helpers.GH.GetMasterBlobSha(ctx, assignmentName)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.UpdateAssignmentBlob(ctx, assignmentName, *blobSHA); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}

// TODO
//
//
func (c *Config) GetFileContents(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqGetFile{}
	if err := marsh.UnmarshalRequest(req, w, r); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	res, err := c.helpers.GH.GetFileContents(ctx, req.FileName, req.DropboxName)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := marsh.MarshalResponse(res, w); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}

// TODO
//
//
func (c *Config) CreateDropbox(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Unmarshal create reference request
	req := &models.ReqCreateDropbox{}
	if err := marsh.UnmarshalRequest(req, w, r); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.GH.CreateDropbox(ctx, req.DropboxName, req.AssignmentName); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.CreateDropbox(ctx, req.DropboxName, req.AssignmentName); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}

// TODO
//
//
func (c *Config) GetSubmissionResults(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	keys := r.URL.Query()
	assignmentName := keys.Get("assignmentName")
	dropboxName := keys.Get("dropboxName")

	req := &models.ReqGetSubmissionResults{
		AssignmentName: assignmentName,
		DropboxName:    dropboxName,
	}

	submission, err := c.helpers.DB.GetSubmission(ctx, req.DropboxName, req.AssignmentName)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	res := &models.ResGetSubmissionResults{
		TestsRan:    submission.TestsRan,
		TestsPassed: submission.TestsPassed,
	}

	if err := marsh.MarshalResponse(res, w); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}

// TODO
//
//
func (c *Config) LeaveFeedbackOnSubmission(w http.ResponseWriter, r *http.Request) {
	//ctx := context.Background()
}
