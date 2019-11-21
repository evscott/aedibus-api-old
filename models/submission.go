package models

type Submission struct {
	Assignment *string `pg:"assignment,pk"`
	Branch     *string `pg:"branch"`
	Submitted  *bool   `pg:"submitted"`
	Grade      *bool   `pg:"grade"`
}
