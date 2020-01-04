package handlers

import (
	"context"
	consts "github.com/evscott/aedibus-api/shared/constants"
	status "github.com/evscott/aedibus-api/shared/http-codes"
	"net/http"

	"github.com/evscott/aedibus-api/models"
	"github.com/evscott/aedibus-api/shared/marsh"
)

// TODO
//
//
func (c *Config) CreateAssignment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqCreateAssignment{}
	if err := marsh.UnmarshalRequest(req, w, r); err != nil {
		c.logger.MarshError("unmarshalling request", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	//
	// Create record of assignment on Github and in database
	//

	if err := c.helpers.GH.CreateRepository(ctx, req.AssignmentName); err != nil {
		c.logger.GalError("creating assignment (repository) on Github", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	assignment, err := c.helpers.DB.CreateAssignment(ctx, req.AssignmentName)
	if err != nil {
		c.logger.DalError("creating assignment in database", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	//
	// Create assignment master branch
	//

	dropbox, err := c.helpers.DB.CreateDropbox(ctx, consts.MASTER, assignment.ID)
	if err != nil {
		c.logger.DalError("creating dropbox in database", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	//
	// Create record of README on Github and in database
	//

	res, err := c.helpers.GH.CreateFile(ctx, assignment.Name, consts.MASTER, consts.README, []byte(req.ReadmeContent))
	if err != nil {
		c.logger.DalError("creating readme file on Github", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.CreateFile(ctx, consts.README, assignment.ID, dropbox.ID, *res.Commit.SHA); err != nil {
		c.logger.DalError("creating readme file in database", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	//
	// Update stupid blob sha as Github requires
	//

	blobSHA, err := c.helpers.GH.GetMasterBlobSha(ctx, assignment.Name)
	if err != nil {
		c.logger.GalError("getting blob sha from Github", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.UpdateAssignmentBlob(ctx, assignment.ID, *blobSHA); err != nil {
		c.logger.DalError("updating blob sha in database", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	//
	// Create dropboxes if dropbox names are provided
	//

	for _, dropboxName := range req.DropboxNames {
		if err := c.helpers.GH.CreateDropbox(ctx, dropboxName, assignment.Name); err != nil {
			c.logger.GalError("creating branch", err)
			w.WriteHeader(status.Status(status.InternalServerError))
		}

		if _, err := c.helpers.DB.CreateDropbox(ctx, dropboxName, assignment.ID); err != nil {
			c.logger.DalError("creating dropbox", err)
			w.WriteHeader(status.Status(status.InternalServerError))
		}
	}

	w.WriteHeader(status.Status(status.OK))
}

func (c *Config) DeleteAssignment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqDeleteAssignment{}
	if err := marsh.UnmarshalRequest(req, w, r); err != nil {
		c.logger.MarshError("unmarshalling request", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.GH.DeleteRepository(ctx, req.AssignmentName); err != nil {
		c.logger.DalError("deleting repository", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	assignment, err := c.helpers.DB.GetAssignmentByName(ctx, req.AssignmentName)
	if err != nil {
		c.logger.DalError("Getting assignment by name", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := c.helpers.DB.DeleteAssignment(ctx, assignment); err != nil {
		c.logger.DalError("Deleting assignment", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	w.WriteHeader(status.Status(status.OK))
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
//
func (c *Config) GetFileContents(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqGetFile{}
	if err := marsh.UnmarshalRequest(req, w, r); err != nil {
		c.logger.MarshError("unmarshalling request", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	res, err := c.helpers.GH.GetFileContents(ctx, req.FileName, req.AssignmentName)
	if err != nil {
		c.logger.GalError("getting file contents", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := marsh.MarshalResponse(res, w); err != nil {
		c.logger.MarshError("marshalling response", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	w.WriteHeader(status.Status(status.OK))
}

// TODO
//
//
func (c *Config) CreateDropbox(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Unmarshal create reference request
	req := &models.ReqCreateDropbox{}
	if err := marsh.UnmarshalRequest(req, w, r); err != nil {
		c.logger.MarshError("unmarshalling response", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	assignment, err := c.helpers.DB.GetAssignmentByName(ctx, req.AssignmentName)
	if err != nil {
		c.logger.DalError("getting assignment", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	for _, dropboxName := range req.DropboxName {
		if err := c.helpers.GH.CreateDropbox(ctx, dropboxName, assignment.Name); err != nil {
			c.logger.GalError("creating branch", err)
			w.WriteHeader(status.Status(status.InternalServerError))
		}

		if _, err := c.helpers.DB.CreateDropbox(ctx, dropboxName, assignment.ID); err != nil {
			c.logger.DalError("creating dropbox", err)
			w.WriteHeader(status.Status(status.InternalServerError))
		}
	}

	w.WriteHeader(status.Status(status.OK))
}

// TODO
//
//
func (c *Config) GetDropboxes(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	keys := r.URL.Query()
	assignmentName := keys.Get("assignmentName")

	assignment, err := c.helpers.DB.GetAssignmentByName(ctx, assignmentName)
	if err != nil {
		c.logger.DalError("getting assignment", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	dropboxes, err := c.helpers.DB.GetDropboxes(ctx, assignment.ID)
	if err != nil {
		c.logger.DalError("getting dropboxes", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	res := &models.ResGetDropboxes{
		Count:     len(dropboxes),
		Dropboxes: dropboxes,
	}

	if err := marsh.MarshalResponse(res, w); err != nil {
		c.logger.MarshError("marshalling response", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}
