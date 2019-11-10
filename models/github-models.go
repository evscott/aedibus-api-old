package models

type ReqCreateRef struct {
	RepoName   string `json:"repoName"`
	BranchName string `json:"branchName"`
}

type ReqCreateRepo struct {
	RepoName string `json:"repoName"`
}
