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
	if err := marsh.UnmarshalRequest(req, w, r); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	README, err := c.helpers.GH.GetReadme(ctx, req.FileName, req.DropboxName)
	if err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := marsh.MarshalResponse(README, w); err != nil {
		c.logger.Error(err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}
