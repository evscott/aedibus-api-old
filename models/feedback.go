package models

type ReqLeaveFeedback struct {
	AssignmentName string `json:"assignmentName"`
	DropboxName    string `json:"dropboxName"`
	FileName       string `json:"fileName"`
	LineNumber     int    `json:"lineNumber"`
	Feedback       string `json:"feedback"`
}
