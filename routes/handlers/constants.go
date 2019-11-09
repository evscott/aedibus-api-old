package handlers

/***** Misc *****/
const (
	Z3E2C  = "z3-e2c"
	MASTER = "master"
)

/***** Commit Messages *****/
const (
	UploadingFile = "Uploading file"
	UpdatingFile  = "UpdatingFile"
)

/***** HTTP Status Codes *****/

type HttpStatus int

const (
	OK                  HttpStatus = 200
	InternalServerError HttpStatus = 500
)
