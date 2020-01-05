package models

import "github.com/go-pg/pg/v9"

type File struct {
	ID        string       `pg:"id,pk"`
	Name      string       `pg:"name,fk"`
	AID       string       `pg:"aid,fk"`
	DID       string       `pg:"did,fk"`
	CommitID  string       `pg:"commit_id"`
	CreatedAt *pg.NullTime `pg:"created_at"`
}

type ReqGetFile struct {
	FileName       string `json:"fileName"`
	AssignmentName string `json:"assignmentName"`
}

type ResGetFile struct {
	FileName       string `json:"fileName"`
	AssignmentName string `json:"assignmentName"`
	DropboxName    string `json:"dropboxName"`
	Content        string `json:"content"`
}

type ReqCreateFile struct {
	AssignmentName string `json:"assignmentName"`
	DropboxName    string `json:"dropboxName"`
	FileName       string `json:"fileName"`
	Content        string `json:"content"`
}
