package models

type File struct {
	Name           string `pg:"name,pk"`
	AssignmentName string `pg:"assignment_name,pk"`
	SubmissionName string `pg:"submission_name,pk"`
}

type ReqGetFile struct {
	FileName       string `json:"fileName"`
	AssignmentName string `json:"assignmentName"`
	SubmissionName string `json:"submissionName"`
}

type ResGetFile struct {
	FileName       string `json:"fileName"`
	AssignmentName string `json:"assignmentName"`
	SubmissionName string `json:"submissionName"`
	Content        string `json:"content"`
}
