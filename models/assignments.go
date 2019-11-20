package models

type Assignment struct {
	Name     string `json:"name,pk"`
	Branch   string `json:"branch,notnull"`
	BlobShah string `json:"blob_shah"`
}

type Submission struct {
	Assignment string `json:"assignment,pk"`
	Branch     string `json:"branch,notnull"`
	Submitted  bool   `json:"submitted"`
	Grade      bool   `json:"grade"`
}
