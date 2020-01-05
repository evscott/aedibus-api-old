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
	AssignmentName string    `json:"assignmentName"`
	List           Dropboxes `json:"list"`
}
