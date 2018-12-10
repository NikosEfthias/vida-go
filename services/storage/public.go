package storage

import (
	"net/http"

	"gitlab.mugsoft.io/vida/go-api/helpers/drivers/files/fs"
)

//Service_public_files serves public files by filename return values are file,mime_type and error
func Service_public_files(fname string) ([]byte, string, error) {
	d, err := fs.Get(fname)
	if nil != err {
		return nil, "", err
	}
	return d, http.DetectContentType(d), err
}
