package event

import (
	//{{{
	"bytes"
	"fmt"
	"html/template"
	"io"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/mugsoft/tools/bytesize"
	"gitlab.mugsoft.io/vida/api/go-api/config"
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
	_f_cost, err := strconv.ParseFloat(cost, 64)
	if nil != err {
		return "", fmt.Errorf("invalid cost number option error: %s", err.Error())
	} else if math.IsNaN(_f_cost) {
		return "", fmt.Errorf("cost cannot be non")
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
		Cost:      _f_cost,
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
	event, err := models.Event_get_by_id(id)
	if nil != err {
		return "", err
	}
	if event.Owner != u.Id {
		return "", fmt.Errorf("event can only be deleted by its owner")
	}

	return event.Id, models.Event_delete(event.Id)
	//}}}
}

func Service_get_by_id(token string, qid string, filter_options interface{}) (interface{}, error) {
	//{{{
	u := storage.Get_user_by_token(token)
	if nil == u {
		return "", services.ERR_N_LOGIN
	}
	e, err := models.Event_get_by_id(qid)
	if nil != err {
		return nil, err
	}
	if !helpers.Can_user_see_event(u.Id, e.Guests, e.Owner) {
		return nil, fmt.Errorf("only the event owner and the guest can see the event details")
	}
	return e, nil //}}}
}
func Service_get_by_owner(token, page string, filter_options interface{}) (interface{}, error) {
	//{{{
	//err checks{{{
	u := storage.Get_user_by_token(token)
	if nil == u {
		return "", services.ERR_N_LOGIN
	}
	_i_page, err := strconv.Atoi(page)
	if nil != err {
		return "", fmt.Errorf("page is not a valid integer err:%s", err.Error())
	} //}}}
	return models.Event_get_by_owner(u.Id, _i_page) //}}}
}
func Service_get_by_participant(token, page string, filter_options interface{}) (interface{}, error) {
	//{{{
	//err checks {{{
	u := storage.Get_user_by_token(token)
	if nil == u {
		return "", services.ERR_N_LOGIN
	}
	_i_page, err := strconv.Atoi(page)
	if nil != err {
		return "", fmt.Errorf("page is not a valid integer err:%s", err.Error())
	} //}}}
	return models.Event_get_by_guest(u.Id, _i_page) //}}}
}
func Service_event_invite(token, event_id, invitees string) (string, error) {
	//{{{
	// error checks{{{
	u := storage.Get_user_by_token(token)
	if nil == u {
		return "", services.ERR_N_LOGIN
	}
	var _invitees = strings.Split(invitees, ":")
	if len(_invitees) < 1 {
		return "", fmt.Errorf("no email to invite")
	}
	event, err := models.Event_get_by_id(event_id)
	if nil != err {
		return "", fmt.Errorf("invalid event id :%s", err.Error())
	}
	if u.Id != event.Owner {
		return "", fmt.Errorf("Only event owners can invite users")
	}
	//}}}
	for _, invitee := range _invitees {
		usr, err := models.User_or_tmp(invitee)
		if nil != err {
			helpers.Log(helpers.ERR, err.Error())
			return "", fmt.Errorf("unexpected system error")
		}
		var buf = new(bytes.Buffer)
		err = template.Must(template.New("mail").Parse(config.Get("APP_INVITATION_TEMPLATE"))).
			Execute(buf, map[string]string{"Name": u.Name, "Link": config.Get("APP_BASE_URL") + "/#/accept_invitation?token=" + usr.Token + "&event_id=" + event_id})
			//TODO:  add custom message to the invitation
		inv, err := models.Invitation_create(models.INV_EVENT, []rune(event_id), u.Id, usr.Id, buf.String())
		if nil != err {
			helpers.Log(helpers.ERR, "invitation cannot be created err:", err)
			return "", fmt.Errorf("Cannot create invitation")
		}
		helpers.SendOneMailPreconfigured(invitee, "Event Invitation", inv.Message)
	}
	return "", fmt.Errorf("not implemented yet")
	//}}}
}
