package utils

import (
	"bytes"
	"io"
	"net/http"
)

func String(s string) *string {
	return &s
}

func Int(i int) *int {
	return &i
}

func Bool(b bool) *bool {
	return &b
}

// TODO
//
func GetFileFromForm(r *http.Request, fileName string) ([]byte, error) {
	file, _, err := r.FormFile(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read file contents through buffer
	buffer := bytes.Buffer{}
	if _, err := io.Copy(&buffer, file); err != nil {
		return nil, err
	}
	defer buffer.Reset()
	return buffer.Bytes(), nil
}
