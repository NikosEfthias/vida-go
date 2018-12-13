package fs

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gitlab.mugsoft.io/vida/go-api/config"
	"gitlab.mugsoft.io/vida/go-api/helpers"
)

var (
	E_INVALID_PATH = fmt.Errorf("invalid path")
	E_INVALID_DATA = fmt.Errorf("invalid data")
	E_WRITE        = fmt.Errorf("write error")
	E_NOT_FOUND    = os.ErrNotExist
)

//internals{{{
//Put file to destination address
func Put(path string, data []byte) error { //{{{
	if "" == path {
		return E_INVALID_PATH
	}
	if nil == data {
		return E_INVALID_DATA
	}
	path = config.Get("PUBLIC_FILES_PATH") + path
	{
		_path := strings.Split(path, "/")
		if len(path) > 1 {
			path := strings.Join(_path[:len(_path)-1], "/") + "/"
			err := os.MkdirAll(path, 0777)
			if nil != err {
				return err
			}
		}
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0777)
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
} //}}}
//Get  file
func Get(path string) ([]byte, error) { //{{{
	if "" == path {
		return nil, E_INVALID_PATH
	}
	path = config.Get("PUBLIC_FILES_PATH") + path
	f, err := os.Open(path)
	if nil != err {
		return nil, E_NOT_FOUND
	}
	defer f.Close()
	return ioutil.ReadAll(f)
} //}}}
//Del  file
func Del(path string) error { //{{{
	return os.Remove(config.Get("PUBLIC_FILES_PATH") + path)
} //}}}
//}}}
func Put_user_data(user_id string, data []byte) (fname string, err error) { //{{{
	fname = user_id + "/" + helpers.Unique_id()
	err = Put(fname, data)
	return
} //}}}
func Put_event_data(user_id string, event_id string, data []byte) (fname string, err error) { //{{{
	fname = user_id + "/" + event_id + "/" + helpers.Unique_id()
	err = Put(fname, data)
	return
} //}}}
