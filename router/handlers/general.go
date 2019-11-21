package handlers

import (
	"context"
	"net/http"

	"github.com/evscott/z3-e2c-api/models"
	"github.com/evscott/z3-e2c-api/shared/marsh"
)

// TODO
//
//
func (c *Config) GetReadme(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	req := &models.ReqGetFile{}
	marsh.UnmarshalRequest(req, w, r)

	c.helpers.GetReadmeHelper(ctx, w, *req.Name, *req.Branch)
}
