package models

type Assignment struct {
	Name    string `pg:"name,pk"`
	BlobSHA string `pg:"blob_sha"`
}

type ReqCreateAssignment struct {
	AssignmentName       string
	InstructionsContents []byte
	TestRunnerContents   []byte
}
