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
func (c *Config) GetReadme(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqGetFile{}
	marsh.UnmarshalRequest(req, w, r)

	README, err := c.helpers.GH.GetReadmeHelper(ctx, req.Name, req.SubmissionName)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	marsh.MarshalResponse(README, w)
}
