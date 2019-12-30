package models

import (
	"github.com/go-pg/pg/v9"
)

type Assignment struct {
	ID        string       `pg:"id,pk"`
	Name      string       `pg:"name,pk"`
	BlobSHA   string       `pg:"blob_sha"`
	CreatedAt *pg.NullTime `pg:"created_at"`
}

type Assignments []Assignment

type ReqCreateAssignment struct {
	AssignmentName       string
	InstructionsContents []byte
	TestRunnerContents   []byte
}

type ResGetAssignment struct {
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	CreatedAt *pg.NullTime `json:"createdAt"`
}

type ResGetAssignments []ResGetAssignment
