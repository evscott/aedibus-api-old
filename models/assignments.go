package models

type Assignment struct {
	Name     *string `pg:"name,pk"`
	Branch   *string `pg:"branch"`
	BlobShah *string `pg:"blob_shah"`
}

type ReqGetFile struct {
	Name   *string `json:"name"`
	Branch *string `json:"branch"`
	Path   *string `json:"path"`
}

type ResGetFile struct {
	Name    *string `json:"name"`
	Branch  *string `json:"branch"`
	Content *string `json:"content"`
}
