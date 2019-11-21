package models

type Submission struct {
	ID         *string `pg:"id"`
	Assignment *string `pg:"assignment,pk"`
	Branch     *string `pg:"branch"`
	Submitted  *bool   `pg:"submitted"`
	Grade      *bool   `pg:"grade"`
}
