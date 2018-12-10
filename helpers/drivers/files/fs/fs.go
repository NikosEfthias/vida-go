package fs

import (
	"fmt"
	"os"

	"gitlab.mugsoft.io/vida/go-api/config"
	"gitlab.mugsoft.io/vida/go-api/helpers"
)

var (
	E_INVALID_PATH = fmt.Errorf("invalid path")
	E_INVALID_DATA = fmt.Errorf("invalid data")
	E_WRITE        = fmt.Errorf("write error")
)

func Put(path string, data []byte) error {
	if "" == path {
		return E_INVALID_PATH
	}
	if nil == data {
		return E_INVALID_DATA
	}
	f, err := os.OpenFile(config.Get("PUBLIC_FILES_PATH")+path, os.O_CREATE|os.O_WRONLY, 0777)
	if nil != err {
		helpers.Log(helpers.ERR, "invalid path error:", err.Error())
		return E_INVALID_PATH
	}

	defer f.Close()
	n, err := f.Write(data)
	if n != len(data) || nil != err {
		helpers.Log(helpers.ERR, "mismatched len or err len=", len(data), "n=", n, "err=", err.Error())
		return E_WRITE
	}
	return err
}
