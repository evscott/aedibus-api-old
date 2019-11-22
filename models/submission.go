package models

type Submission struct {
	Name           string `pg:"name,pk"`
	AssignmentName string `pg:"assignment_name,pk"`
	Submitted      bool   `pg:"submitted"`
	Grade          bool   `pg:"grade"`
}

type ReqCreateSubmission struct {
	Name           string `json:"submissionName"`
	AssignmentName string `json:"assignmentName"`
}

type ReqPullRequest struct {
	SubmissionName string `json:"submissionName"`
	AssignmentName string `json:"assignmentName"`
	Body           string `json:"body"`
}
