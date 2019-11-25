package models

type ResGetComment struct {
	Body       string `json:"body"`
	LineNumber int    `json:"lineNumber"`
	FileName   string `json:"fileName"`
	CommitID   string `json:"commitID"`
}

type ReqLeaveFeedback struct {
	AssignmentName string `json:"assignmentName"`
	DropboxName    string `json:"dropboxName"`
	FileName       string `json:"fileName"`
	LineNumber     int    `json:"lineNumber"`
	Feedback       string `json:"feedback"`
}

type ReqGetFeedback struct {
	AssignmentName string `json:"assignmentName"`
	DropboxName    string `json:"dropboxName"`
	FileName       string `json:"fileName"`
}

type ResGetFeedback []ResGetComment
