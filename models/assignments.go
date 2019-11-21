package models

type Assignment struct {
	Name     *string `pg:"name,pk"`
	Branch   *string `pg:"branch"`
	BlobShah *string `pg:"blob_shah"`
}
