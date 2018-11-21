package event

import (
	//{{{
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/mugsoft/tools/bytesize"
	"gitlab.mugsoft.io/vida/api/go-api/helpers"
	"gitlab.mugsoft.io/vida/api/go-api/models"
	"gitlab.mugsoft.io/vida/api/go-api/services"
	"gitlab.mugsoft.io/vida/api/go-api/services/storage"
	//}}}
)

func Service_create(token, title, loc, startdate, enddate, details, max_num_guest, min_num_guest, cost, votable string, img io.Reader) (string, error) {
	//{{{
	// err checks{{{
	const LIMIT_FILESIZE = bytesize.MB * 10
	var ALLOWED_MIMES = []string{"jpeg", "jpg", "png", "jpeg"}
	if img == nil {
		return "", fmt.Errorf("cannot read the img")
	}
	u := storage.Get_user_by_token(token)
	if nil == u {
		return "", services.ERR_N_LOGIN
	}
	err := helpers.Check_missing_fields(
		[]string{"title", "location", "start_date", "end_date", "details", "max_num_guest", "min_num_guest", "cost", "votable"},
		[]string{title, loc, startdate, enddate, details, max_num_guest, min_num_guest, cost, votable})
	if nil != err {
		return "", err
	}
	__i_start_date, err := strconv.ParseInt(startdate, 10, 64)
	if nil != err {
		return "", fmt.Errorf("invalid date error: %s", err.Error())
	}
	__i_end_date, err := strconv.ParseInt(enddate, 10, 64)
	if nil != err {
		return "", fmt.Errorf("invalid date error: %s", err.Error())
	}
	__i_min_num_guests, err := strconv.Atoi(min_num_guest)
	if nil != err {
		return "", fmt.Errorf("invalid min-max guest number option error: %s", err.Error())
	}
	__i_max_num_guests, err := strconv.Atoi(max_num_guest)
	if nil != err {
		return "", fmt.Errorf("invalid min-max guest number option error: %s", err.Error())
	}
	__f_cost, err := strconv.ParseFloat(cost, 64)
	if nil != err {
		return "", fmt.Errorf("invalid min-max guest number option error: %s", err.Error())
	}
	__data_url, err := helpers.Multipart_to_data_url(img, LIMIT_FILESIZE, ALLOWED_MIMES)
	if nil != err {
		return "", fmt.Errorf("cannot read event photo error:%s", err.Error())
	}
	__b_votable, err := strconv.ParseBool(votable)
	if nil != err {
		return "", fmt.Errorf("votable is not a valid bool error:%s", err.Error())
	}
	//}}}
	event := &models.Event{
		Owner:     u.Id,
		Title:     title,
		Loc:       loc,
		Detail:    details,
		MaxGuest:  __i_max_num_guests,
		MinGuest:  __i_min_num_guests,
		Cost:      __f_cost,
		Img:       __data_url,
		Votable:   __b_votable,
		StartDate: time.Unix(__i_start_date, 0),
		EndDate:   time.Unix(__i_end_date, 0),
	}
	err = models.Event_new(event)
	//TODO:  check  if dates are on the future
	return event.Id, err //}}}
}
func Service_delete(token, id string) (string, error) {
	//{{{
	u := storage.Get_user_by_token(token)
	if nil == u {
		return "", services.ERR_N_LOGIN
	}
	return "", fmt.Errorf("not implemented yet")
	//}}}
}

func Service_get_by_id(token string, qid string, filter_options interface{}) (interface{}, error) {
	//{{{
	u := storage.Get_user_by_token(token)
	if nil == u {
		return "", services.ERR_N_LOGIN
	}
	return nil, fmt.Errorf("not implemented") //}}}
}
func Service_get_by_owner(token, start, end string, filter_options interface{}) (interface{}, error) {
	//{{{
	u := storage.Get_user_by_token(token)
	if nil == u {
		return "", services.ERR_N_LOGIN
	}
	return nil, fmt.Errorf("not implemented") //}}}
}
func Service_get_by_participant(token, start, end string, filter_options interface{}) (interface{}, error) {
	//{{{
	u := storage.Get_user_by_token(token)
	if nil == u {
		return "", services.ERR_N_LOGIN
	}
	return nil, fmt.Errorf("not implemented") //}}}
}
