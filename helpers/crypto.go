package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sync"
	"time"
)

var unique__lock sync.Mutex

//MD5 hashing made easy
func MD5(k string) string {
	_s := md5.Sum([]byte(k))
	return hex.EncodeToString(_s[:])
}

//Unique_id generate
func Unique_id() string {
	unique__lock.Lock()
	defer unique__lock.Unlock()
	return MD5(fmt.Sprint(time.Now().UnixNano()))
}
