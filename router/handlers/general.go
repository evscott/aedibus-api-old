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

	req := &models.ReqGetFile{}
	if err := marsh.UnmarshalRequest(req, w, r); err != nil {
		c.logger.MarshError("unmarshalling request", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	README, err := c.helpers.GH.GetReadme(ctx, req.FileName, req.DropboxName)
	if err != nil {
		c.logger.GalError("getting README", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}

	if err := marsh.MarshalResponse(README, w); err != nil {
		c.logger.MarshError("marshalling response", err)
		w.WriteHeader(status.Status(status.InternalServerError))
	}
}
