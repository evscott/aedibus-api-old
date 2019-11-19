package Routes

type Resource string

/***** API Resources *****/
const (
	Github      Resource = "/github"
	Repository  Resource = "/repository"
	Branch      Resource = "/branch"
	File        Resource = "/file"
	PullRequest Resource = "/pull-request"
	Comment     Resource = "/comment"
)

/***** HTTP Methods *****/
const (
	POST   = "POST"
	GET    = "GET"
	PUT    = "PUT"
	DELETE = "DELETE"
)
