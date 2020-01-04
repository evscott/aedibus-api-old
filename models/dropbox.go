package models

type Dropbox struct {
	ID   string `pg:"id"`
	Name string `pg:"name"`
	AID  string `pg:"aid"`
}

type Dropboxes []Dropbox

type ReqCreateDropbox struct {
	AssignmentName string   `json:"assignmentName"`
	DropboxName    []string `json:"dropboxName"`
}

type ResGetDropboxes struct {
	Count     int       `json:"count"`
	Dropboxes Dropboxes `json:"dropboxes"`
}

type ReqPullRequest struct {
	DID  string `json:"did"`
	AID  string `json:"aid"`
	Body string `json:"body"`
}
