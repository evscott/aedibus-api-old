package Routes

import "strings"

type Resource string

/***** API Resources *****/
const (
	Student     Resource = "/student"
	Instructor  Resource = "/instructor"
	Assignment  Resource = "/assignment"
	Assignments Resource = "/assignments"
	Dropbox     Resource = "/dropbox"
	Dropboxes   Resource = "/dropboxes"
	Submit      Resource = "/submit"
	Submission  Resource = "/submission"
	File        Resource = "/file"
	Contents    Resource = "/contents"
	Readme      Resource = "/readme"
	Feedback    Resource = "/feedback"
)

/***** HTTP Methods *****/
const (
	POST   = "POST"
	GET    = "GET"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func Path(resources ...Resource) string {
	var route strings.Builder
	for _, r := range resources {
		route.WriteString(string(r))
	}

	return route.String()
}
