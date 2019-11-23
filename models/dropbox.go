package models

type Dropbox struct {
	Name           string `pg:"name,pk"`
	AssignmentName string `pg:"assignment_name,pk"`
	Submitted      bool   `pg:"submitted"`
	Grade          bool   `pg:"grade"`
}

type ReqCreateDropbox struct {
	DropboxName    string `json:"dropboxName"`
	AssignmentName string `json:"assignmentName"`
}

type ReqPullRequest struct {
	DropboxName    string `json:"dropboxName"`
	AssignmentName string `json:"assignmentName"`
	Body           string `json:"body"`
}
