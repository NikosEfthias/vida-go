package helpers

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Multipart_to_data_url(img io.Reader, limit_filesize uint64, ALLOWED_MIMES []string) (string, error) {
	magic := make([]byte, 512)
	imgbuf := make([]byte, limit_filesize)
	var MIME string
	n, err := img.Read(magic)
	if nil != err {
		return "", err
	} else if n < 512 {
		return "", fmt.Errorf("too small")
	} else {
		//else is just for scoping variables here
		var valid_mime bool
		MIME = http.DetectContentType(magic)
		for i := range ALLOWED_MIMES {
			if strings.Contains(MIME, ALLOWED_MIMES[i]) {
				valid_mime = true
				break
			}
		}
		if !valid_mime {
			return "", fmt.Errorf("invalid image type")
		}
	}
	n, err = img.Read(imgbuf)
	if nil != err {
		return "", err
	} else if uint64(n) > limit_filesize-512 {
		//we already consumed the first 512 bytes so if the read amount is big img is bigger no matter what
		return "", fmt.Errorf("too big")
	}
	var __data_url = "data:" + MIME + ";base64," + base64.StdEncoding.EncodeToString(append(magic, imgbuf[:n]...))
	return __data_url, nil
}
