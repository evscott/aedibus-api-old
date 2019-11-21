package Routes

import "strings"

type Resource string

/***** API Resources *****/
const (
	Github      Resource = "/github"
	Repository  Resource = "/repository"
	Branch      Resource = "/branch"
	File        Resource = "/file"
	PullRequest Resource = "/pull-request"
	Comment     Resource = "/comment"

	Readme Resource = "/readme"
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
