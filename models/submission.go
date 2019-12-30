package models

import "time"

type Submission struct {
	AssignmentName string    `pg:"assignment_name,pk"`
	DropboxName    string    `pg:"dropbox_name,pk"`
	PrNumber       int       `pg:"pr_number"`
	TestsRan       int       `pg:"tests_ran"`
	TestsPassed    int       `pg:"tests_passed"`
	CreatedAt      time.Time `pg:"created_at"`
}

type ReqGetSubmissionResults struct {
	AssignmentName string `json:"assignmentName"`
	DropboxName    string `pg:"dropboxName"`
}

type ResGetSubmissionResults struct {
	TestsRan    int `json:"numberOfTests"`
	TestsPassed int `json:"testsPassed"`
}
