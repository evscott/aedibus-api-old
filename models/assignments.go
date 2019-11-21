package models

type Assignment struct {
	Name     *string `pg:"name,pk"`
	Branch   *string `pg:"branch"`
	BlobShah *string `pg:"blob_shah"`
}

type Submission struct {
	Assignment *string `pg:"assignment,pk"`
	Branch     *string `pg:"branch"`
	Submitted  *bool   `pg:"submitted"`
	Grade      *bool   `pg:"grade"`
}
