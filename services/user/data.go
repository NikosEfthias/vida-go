package user

import (
	"fmt"

	"github.com/mugsoft/vida/services/storage"
)

func Service_get(token string) (interface{}, error) {
	u := storage.Get_user_by_token(token)
	if nil == u {
		return nil, fmt.Errorf("not logged in")
	}
	return u, nil
}
