package models

type ReqCreatePR struct {
	RepoName string `json:"repoName"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Head     string `json:"head"`
}

type ReqCreateRef struct {
	RepoName   string `json:"repoName"`
	BranchName string `json:"branchName"`
}

type ReqCreateRepo struct {
	RepoName string `json:"repoName"`
}
