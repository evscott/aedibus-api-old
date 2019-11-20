package models

type Assignments struct {
	Name         string `json:"name,pk"`
	Branch       string `json:"branch"`
	BlobShah     string `json:"blob_shah"`
	LatestCommit string `json:"latest_commit"`
}

type Submissions struct {
	Assignment   string `json:"assignment,pk"`
	Branch       string `json:"branch"`
	LatestCommit string `json:"latest_commit"`
	Submitted    bool   `json:"submitted"`
	Grade        bool   `json:"grade"`
}
