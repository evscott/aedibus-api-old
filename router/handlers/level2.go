package handlers

import (
	"context"
	"fmt"
	"github.com/evscott/aedibus-api/shared/utils"
	"net/http"

	"github.com/evscott/aedibus-api/models"
	status "github.com/evscott/aedibus-api/shared/http-codes"
	"github.com/evscott/aedibus-api/shared/marsh"
)

func (c *Config) GetAssignments(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	assignments, err := c.helpers.DB.GetAssignments(ctx)
	if err != nil {
		c.logger.DalError("getting assignments from DB", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	res := make(models.ResGetAssignments, len(assignments))
	for i, a := range assignments {
		fmt.Printf("%v\n", a)
		res[i].ID = a.ID
		res[i].Name = a.Name
		res[i].CreatedAt = a.CreatedAt
	}

	if err := marsh.MarshalResponse(res, w); err != nil {
		c.logger.MarshError("marshalling response", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}

// TODO
//
//
func (c *Config) CreateDropboxFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	assignmentName := r.FormValue("assignmentName")
	dropboxName := r.FormValue("dropboxName")
	fileName := r.FormValue("fileName")

	contents, err := utils.GetFileFromForm(r, fileName)
	if err != nil {
		c.logger.UtilsError("getting file from form", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	res, err := c.helpers.GH.CreateFile(ctx, assignmentName, dropboxName, fileName, contents)
	if err != nil {
		c.logger.GalError("creating file", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	blobSHA, err := c.helpers.GH.GetMasterBlobSha(ctx, assignmentName)
	if err != nil {
		c.logger.GalError("getting blob sha", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.UpdateAssignmentBlob(ctx, assignmentName, *blobSHA); err != nil {
		c.logger.DalError("updating blob sha", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.CreateFile(ctx, fileName, assignmentName, dropboxName, *res.Commit.SHA); err != nil {
		c.logger.DalError("creating file", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}

// TODO
//
func (c *Config) SubmitAssignment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqPullRequest{}
	if err := marsh.UnmarshalRequest(req, w, r); err != nil {
		c.logger.MarshError("unmarshalling request", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	res, err := c.helpers.GH.CreatePullRequest(ctx, req.DropboxName, req.AssignmentName, req.DropboxName, req.Body)
	if err != nil {
		c.logger.GalError("creating pull request", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.CreateSubmission(ctx, req.DropboxName, req.AssignmentName, *res.Number); err != nil {
		c.logger.DalError("creating submission", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}
