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

type ReqGetSubmissionResults struct {
	AID string `json:"assignmentName"`
	DID string `pg:"dropboxName"`
}

type ResGetSubmissionResults struct {
	TestsRan    int `json:"numberOfTests"`
	TestsPassed int `json:"testsPassed"`
}
