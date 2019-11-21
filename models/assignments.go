package models

type Assignment struct {
	Name     *string `pg:"name,pk"`
	Branch   *string `pg:"branch"`
	BlobShah *string `pg:"blob_shah"`
}

type ReqGetAssignment struct {
	Name   *string `json:"name"`
	Branch *string `json:"branch"`
}

type ResGetAssignment struct {
	Name    *string `json:"name"`
	Branch  *string `json:"branch"`
	Content *string `json:"content"`
}
