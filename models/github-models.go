package models

type ReqCreateRef struct {
	Repo   string `json:"repo"`
	Branch string `json:"branch"`
}

type ReqCreateRepo struct {
	Repo string `json:"repo"`
}
