package models

type Dropbox struct {
	ID   string `pg:"id"`
	Name string `pg:"name"`
	AID  string `pg:"aid"`
}

type ReqCreateDropbox struct {
	DropboxName string `json:"dropboxName"`
	AID         string `json:"aid"`
}

type ReqPullRequest struct {
	DID  string `json:"did"`
	AID  string `json:"aid"`
	Body string `json:"body"`
}
