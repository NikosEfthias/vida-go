package helpers

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Multipart_to_data_url(img io.Reader, limit_filesize uint64, ALLOWED_MIMES []string) (string, error) { //{{{
	MIME, data, err := Multipart_to_byte_slice(img, limit_filesize, ALLOWED_MIMES)
	if nil != err {
		return "", err
	}
	var __data_url = "data:" + MIME + ";base64," + base64.StdEncoding.EncodeToString(data)
	return __data_url, nil
} //}}}

func Multipart_to_byte_slice(img io.Reader, limit_filesize uint64, ALLOWED_MIMES []string) (string, []byte, error) { //{{{
	magic := make([]byte, 512)
	imgbuf := make([]byte, limit_filesize)
	n, err := img.Read(magic)
	if nil != err {
		return "", nil, err
	} else if n < 512 {
		return "", nil, fmt.Errorf("too small")
	}
	MIME, valid := Check_mime(magic, ALLOWED_MIMES)
	if !valid {
		return "", nil, fmt.Errorf("invalid image type")
	}
	n, err = img.Read(imgbuf)
	if nil != err {
		return "", nil, err
	} else if uint64(n) > limit_filesize-512 {
		//we already consumed the first 512 bytes so if the read amount is big img is bigger no matter what
		return "", nil, fmt.Errorf("too big")
	}
	return MIME, append(magic, imgbuf...), nil
} //}}}

func Check_mime(magic []byte, ALLOWED_MIMES []string) (MIME string, valid_mime bool) { //{{{
	MIME = http.DetectContentType(magic)
	for i := range ALLOWED_MIMES {
		if strings.Contains(MIME, ALLOWED_MIMES[i]) {
			valid_mime = true
			break
		}
	}
	return
} //}}}
