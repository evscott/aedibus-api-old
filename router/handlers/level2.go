package handlers

import (
	"context"
	"net/http"

	"github.com/evscott/aedibus-api/models"
	status "github.com/evscott/aedibus-api/shared/http-codes"
	"github.com/evscott/aedibus-api/shared/marsh"
)

// TODO
//
//
func (c *Config) GetReadme(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	keys := r.URL.Query()
	assignmentName := keys.Get("assignmentName")

	README, err := c.helpers.GH.GetReadme(ctx, assignmentName)
	if err != nil {
		c.logger.GalError("getting README.md from Github", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := marsh.MarshalResponse(README, w); err != nil {
		c.logger.MarshError("marshalling response", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}

func (c *Config) GetAssignments(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	assignments, err := c.helpers.DB.GetAssignments(ctx)
	if err != nil {
		c.logger.DalError("getting assignments from DB", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	res := make(models.ResGetAssignments, len(assignments))

	for i, a := range assignments {
		res[i].ID = a.ID
		res[i].Name = a.Name
		res[i].CreatedAt = a.CreatedAt

		// Get README.md content for each assignment
		readme, err := c.helpers.GH.GetReadme(ctx, res[i].Name)
		if err != nil {
			c.logger.GalError("getting assignments from DB", err)
			w.WriteHeader(status.Status(status.InternalServerError))
		}

		res[i].ReadmeContent = readme.Content
	}

	if err := marsh.MarshalResponse(res, w); err != nil {
		c.logger.MarshError("marshalling response", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}

// TODO
//
//
func (c *Config) CreateFile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqCreateFile{}
	if err := marsh.UnmarshalRequest(req, w, r); err != nil {
		c.logger.MarshError("unmarshalling request", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	//
	// Create record of file on Github
	//

	gitFile, err := c.helpers.GH.CreateFile(ctx, req.AssignmentName, req.DropboxName, req.FileName, []byte(req.Content))
	if err != nil {
		c.logger.GalError("creating file", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	//
	// Get assignment + dropbox for their IDs
	//

	assignment, err := c.helpers.DB.GetAssignmentByName(ctx, req.AssignmentName)
	if err != nil {
		c.logger.DalError("getting assignment", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	dropbox, err := c.helpers.DB.GetDropbox(ctx, assignment.ID, req.DropboxName)
	if err != nil {
		c.logger.DalError("getting dropboxes", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	//
	// Create record of file in database
	//

	if err := c.helpers.DB.CreateFile(ctx, req.FileName, assignment.ID, dropbox.ID, *gitFile.Commit.SHA); err != nil {
		c.logger.DalError("creating file", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	//
	// Update stupid blob sha as Github requires
	//

	blobSHA, err := c.helpers.GH.GetMasterBlobSha(ctx, req.AssignmentName)
	if err != nil {
		c.logger.GalError("getting blob sha", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.UpdateAssignmentBlob(ctx, assignment.ID, *blobSHA); err != nil {
		c.logger.GalError("updating blob sha", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	w.WriteHeader(status.Status(status.OK))
}

// TODO
//
func (c *Config) SubmitAssignment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqSubmitAssignment{}
	if err := marsh.UnmarshalRequest(req, w, r); err != nil {
		c.logger.MarshError("unmarshalling request", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	//
	// Create record of file on Github
	//

	gitFile, err := c.helpers.GH.CreateFile(ctx, req.AssignmentName, req.DropboxName, req.FileName, []byte(req.Content))
	if err != nil {
		c.logger.GalError("creating file", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	//
	// Get assignment + dropbox for their IDs
	//

	assignment, err := c.helpers.DB.GetAssignmentByName(ctx, req.AssignmentName)
	if err != nil {
		c.logger.DalError("getting assignment", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
	dropbox, err := c.helpers.DB.GetDropbox(ctx, assignment.ID, req.DropboxName)
	if err != nil {
		c.logger.DalError("getting dropboxes", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	//
	// Create record of file in database
	//

	if err := c.helpers.DB.CreateFile(ctx, req.FileName, assignment.ID, dropbox.ID, *gitFile.Commit.SHA); err != nil {
		c.logger.DalError("creating file", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	//
	// Update stupid blob sha as Github requires
	//

	blobSHA, err := c.helpers.GH.GetMasterBlobSha(ctx, req.AssignmentName)
	if err != nil {
		c.logger.GalError("getting blob sha", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
	if err := c.helpers.DB.UpdateAssignmentBlob(ctx, assignment.ID, *blobSHA); err != nil {
		c.logger.GalError("updating blob sha", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	//
	// Create pull request and record of submission in database
	//

	res, err := c.helpers.GH.CreatePullRequest(ctx, req.DropboxName, req.AssignmentName, req.DropboxName, req.Body)
	if err != nil {
		c.logger.GalError("creating pull request", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.CreateSubmission(ctx, dropbox.ID, assignment.ID, *res.Number); err != nil {
		c.logger.DalError("creating submission", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}

// TODO
//
func (c *Config) GetSubmission(w http.ResponseWriter, r *http.Request) {}
