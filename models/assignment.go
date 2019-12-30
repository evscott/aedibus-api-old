package models

type Assignment struct {
	Name    string `pg:"name,pk"`
	BlobSHA string `pg:"blob_sha"`
}

type Assignments []Assignment

type ReqCreateAssignment struct {
	AssignmentName       string
	InstructionsContents []byte
	TestRunnerContents   []byte
}

type ResGetAssignment struct {
	Name string `json:"name"`
}

type ResGetAssignments []ResGetAssignment
