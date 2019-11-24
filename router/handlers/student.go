package handlers

import (
	"context"
	"net/http"

	"github.com/evscott/z3-e2c-api/models"
	status "github.com/evscott/z3-e2c-api/shared/http-codes"
	"github.com/evscott/z3-e2c-api/shared/marsh"
)

// TODO
//
//
func (c *Config) CreateDropboxFile(w http.ResponseWriter, r *http.Request) {
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

	blobSHA, err := c.helpers.GH.GetMasterBlobSha(ctx, assignmentName)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.UpdateAssignmentBlob(ctx, assignmentName, *blobSHA); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.CreateFile(ctx, fileName, assignmentName, dropboxName, res.Commit.String()); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}

// TODO
//
func (c *Config) SubmitAssignment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqPullRequest{}
	if err := marsh.UnmarshalRequest(req, w, r); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	res, err := c.helpers.GH.CreatePullRequest(ctx, req.DropboxName, req.AssignmentName, req.DropboxName, req.Body)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.CreateSubmission(ctx, req.DropboxName, req.AssignmentName, *res.Number); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}
