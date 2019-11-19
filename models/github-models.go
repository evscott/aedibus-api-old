package models

type ReqCreateComment struct {
	RepoName *string `json:"repoName"`
	Path     *string `json:"path"`
	Body     *string `json:"body"`
	Position *int    `json:"position"`
	CommitID *string `json:"commitID"`
}

type ReqCreateBranch struct {
	RepoName   string `json:"repoName"`
	BranchName string `json:"branchName"`
}

type ReqCreateRepo struct {
	RepoName string `json:"repoName"`
}
