package models

type Submission struct {
	AssignmentName string `pg:"assignment_name"`
	DropboxName    string `pg:"dropbox_name"`
	Grade          string `pg:"grade"`
	PrNumber       int    `pg:"pr_number"`
}
