package http_codes

type HttpStatus int

const (
	OK                  HttpStatus = 200
	InternalServerError HttpStatus = 500
)

func Status(status HttpStatus) int {
	return int(status)
}
