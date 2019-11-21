package models

type File struct {
	SubmissionID *string `pg:"submission_id"`
	FileName     *string `json:"file_name"`
}

type ReqGetFile struct {
	Name   *string `json:"name"`
	Branch *string `json:"branch"`
	Path   *string `json:"path"`
}

type ResGetFile struct {
	Name    *string `json:"name"`
	Branch  *string `json:"branch"`
	Content *string `json:"content"`
}
