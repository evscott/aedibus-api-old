package models

type File struct {
	Name           string `pg:"name,pk"`
	AssignmentName string `pg:"assignment_name,pk"`
	DropboxName    string `pg:"dropbox_name,pk"`
}

type ReqGetFile struct {
	FileName       string `json:"fileName"`
	AssignmentName string `json:"assignmentName"`
	DropboxName    string `json:"dropboxName"`
}

type ResGetFile struct {
	FileName       string `json:"fileName"`
	AssignmentName string `json:"assignmentName"`
	DropboxName    string `json:"dropboxName"`
	Content        string `json:"content"`
}
