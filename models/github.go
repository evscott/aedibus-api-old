package models

type ReqCreateSubmission struct {
	Name           string `json:"name"`
	AssignmentName string `json:"assignmentName"`
}

type ReqCreateAssignment struct {
	Name string `json:"assignmentName"`
}
