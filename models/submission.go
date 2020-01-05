package models

import "time"

type Submission struct {
	ID        string    `pg:"id,pk"`
	AID       string    `pg:"aid,fk"`
	DID       string    `pg:"did,fk"`
	PrNumber  int       `pg:"pr_number"`
	CreatedAt time.Time `pg:"created_at"`
}

type SubmissionResults struct {
	SID         string    `pg:"sid,pk"`
	TestsRan    int       `pg:"tests_ran"`
	TestsPassed int       `pg:"tests_passed"`
	Reviewed    bool      `pg:"reviewed"`
	CreatedAt   time.Time `pg:"created_at"`
}

type ReqGetSubmission struct {
	AssignmentName string `json:"assignmentName"`
	DropboxName    string `json:"dropboxName"`
}

type ResGetSubmission struct {
	AssignmentName string `json:"assignmentName"`
	DropboxName    string `json:"dropboxName"`
	Content        string `json:"content"`
}

type ReqSubmitAssignment struct {
	AssignmentName string `json:"assignmentName"`
	DropboxName    string `json:"dropboxName"`
	Body           string `json:"body"`
	FileName       string `json:"fileName"`
	Content        string `json:"content"`
}
