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
	"gitlab.mugsoft.io/vida/go-api/config"
	"gitlab.mugsoft.io/vida/go-api/helpers"
	"gitlab.mugsoft.io/vida/go-api/helpers/drivers/files/fs"
	"gitlab.mugsoft.io/vida/go-api/models"
	"gitlab.mugsoft.io/vida/go-api/services"
	"gitlab.mugsoft.io/vida/go-api/services/storage"
	//}}}
)

var (
	ERR_EVENT_NOT_FOUND    = fmt.Errorf("invalid event id")
	ERR_INVALID_EVENT_ID   = fmt.Errorf("invalid event id format")
	ERR_INVALID_TIME       = fmt.Errorf("invalid time format")
	ERR_NOT_INVITED        = fmt.Errorf("you are not invited or accepted your invitation")
	ERR_INVALID_TIME_RANGE = fmt.Errorf("time cannot be earlier then the event start date or later than the event end date")
	ERR_EVENT_NOT_VOTABLE  = fmt.Errorf("event is not votable")
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
	_, data, err := helpers.Multipart_to_byte_slice(img, LIMIT_FILESIZE, ALLOWED_MIMES)
	if nil != err {
		return "", fmt.Errorf("cannot read event photo error:%s", err.Error())
	}
	__b_votable, err := strconv.ParseBool(votable)
	if nil != err {
		return "", fmt.Errorf("votable is not a valid bool error:%s", err.Error())
	}
	var event_id = helpers.Unique_id()
	fname, err := fs.Put_event_data(u.Id, event_id, data)
	if nil != err {
		return "", err
	}
	//}}}
	event := &models.Event{
		Id:        event_id,
		Owner:     u.Id,
		Title:     title,
		Loc:       loc,
		Detail:    details,
		MaxGuest:  __i_max_num_guests,
		MinGuest:  __i_min_num_guests,
		Cost:      _f_cost,
		Img:       "/static/public/" + fname,
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
func Service_update(token, event_id, field string, value interface{}) (string, error) {
	//{{{
	//{{{
	u := storage.Get_user_by_token(token)
	if nil == u {
		return "", services.ERR_N_LOGIN
	}
	event, err := models.Event_get_by_id(event_id)
	if nil != err {
		return "", err
	}
	if event.Owner != u.Id {
		return "", fmt.Errorf("event can only be updated by its owner")
	} //}}}
	switch field { //type checks{{{
	case "title":
	case "location":
	case "start_date": //{{{
		_i_val, err := strconv.Atoi(value.(string))
		if nil != err {
			return "", fmt.Errorf("start date value (%s) is not a valid integer", value.(string))
		}
		value = time.Unix(int64(_i_val), 0) //}}}
	case "end_date": //{{{
		_i_val, err := strconv.Atoi(value.(string))
		if nil != err {
			return "", fmt.Errorf("end value (%s) is not a valid integer", value.(string))
		}
		value = time.Unix(int64(_i_val), 0) //}}}
	case "details":
	case "max_num_guest": //{{{
		__i_max_num_guests, err := strconv.Atoi(value.(string))
		if nil != err {
			return "", fmt.Errorf("max num guest value (%s) is not a valid integer", value.(string))
		}
		value = __i_max_num_guests //}}}
	case "min_num_guest": //{{{
		__i_min_num_guests, err := strconv.Atoi(value.(string))
		if nil != err {
			return "", fmt.Errorf("min num guest value (%s) is not a valid integer", value.(string))
		}
		value = __i_min_num_guests //}}}
	case "cost": //{{{
		_f_cost, err := strconv.ParseFloat(value.(string), 64)
		if nil != err {
			return "", fmt.Errorf("cost value (%s) is not a valid floating point number", value.(string))
		}
		if math.IsNaN(_f_cost) {
			return "", fmt.Errorf("cost cannot be NaN")
		}
		value = _f_cost //}}}
	case "votable": //{{{
		__b_votable, err := strconv.ParseBool(value.(string))
		if nil != err {
			return "", fmt.Errorf("votable value (%s) is not a valid boolean", value.(string))
		}
		value = __b_votable //}}}
	} //}}}
	return "success", models.Event_update(event_id, field, value) //}}}
}
func Service_update_img(token, event_id string, file io.Reader) (string, error) {
	//{{{
	//{{{
	const LIMIT_FILESIZE = bytesize.MB * 10
	var ALLOWED_MIMES = []string{"jpg", "png", "jpeg"}
	if file == nil {
		return "", fmt.Errorf("cannot read the file")
	}
	u := storage.Get_user_by_token(token)
	if nil == u {
		return "", services.ERR_N_LOGIN
	}
	event, err := models.Event_get_by_id(event_id)
	if nil != err {
		return "", err
	}
	if event.Owner != u.Id {
		return "", fmt.Errorf("event can only be updated by its owner")
	} //}}}
	_, data, err := helpers.Multipart_to_byte_slice(file, LIMIT_FILESIZE, ALLOWED_MIMES)
	if nil != err {
		return "", fmt.Errorf("cannot process the file error : %s", err.Error())
	}
	fname, err := fs.Put_event_data(u.Id, event_id, data)
	if nil != err {
		return "", err
	}
	return "success", models.Event_update(event_id, "img", "/static/public/"+fname) //}}}
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
	if !helpers.Can_user_see_event(u.Id, e.GetGuestIds(), e.Owner) {
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
func Service_get_by_participant(token, page string, filter_options map[string]interface{}) (interface{}, error) {
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
	return models.Event_get_by_guest(u.Id, _i_page, filter_options) //}}}
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
		//users cannot invite themselves
		if invitee == u.Email {
			continue
		}
		usr, err := models.User_or_tmp(invitee)
		if nil != err {
			helpers.Log(helpers.ERR, err.Error())
			return "", fmt.Errorf("unexpected system error")
		}
		var buf = new(bytes.Buffer)
		err = template.Must(template.New("mail").Parse(config.Get("APP_INVITATION_TEMPLATE"))).
			Execute(buf, map[string]string{"Name": u.Name + " " + u.Lastname, "Link": config.Get("APP_BASE_URL") + "/#/dashboard?token=" + usr.Token + "&event_id=" + event_id})
			//TODO:  add custom message to the invitation
		inv, err := models.Invitation_create(models.INV_EVENT, []rune(event_id), u.Id, usr.Id, buf.String())
		if nil != err {
			helpers.Log(helpers.ERR, "invitation cannot be created err:", err.Error())
			return "", fmt.Errorf("Cannot create invitation reason: %v", err.Error())
		}
		helpers.SendOneMailPreconfigured(invitee, "Event Invitation", inv.Message)
		//if the user is already logged in then use the old token instead of creating a new one and logging out the user
		if _usr := storage.Get_user_by_id(usr.Id); nil != _usr && _usr.Token != "" {
			usr.Token = _usr.Token
		}
		storage.Add_or_update_user(usr)
	}
	return "success", nil
	//}}}
}
func Service_event_accept(token, event_id string) (string, error) {
	//{{{
	//error checks {{{
	u := storage.Get_user_by_token(token)
	if nil == u {
		return "", services.ERR_N_LOGIN
	}
	inv, err := models.Invitation_get_by_invitee(models.INV_EVENT, u.Id, event_id)
	if len(inv) < 1 || err != nil {
		return "", fmt.Errorf("invitation does not exist")
	} //}}}

	return "success", models.Invitation_accept(event_id, u.Id) //}}}
}
func Service_event_decline(token, event_id string) (string, error) {
	//{{{
	//error checks {{{
	u := storage.Get_user_by_token(token)
	if nil == u {
		return "", services.ERR_N_LOGIN
	}
	inv, err := models.Invitation_get_by_invitee(models.INV_EVENT, u.Id, event_id)
	if len(inv) < 1 || err != nil {
		return "", fmt.Errorf("invitation does not exist")
	} //}}}
	return "success", models.Invitation_decline(event_id, u.Id) //}}}
}
func Service_vote(token, event_id, time string) (string, error) {
	//{{{
	//error checks{{{
	u := storage.Get_user_by_token(token)
	if nil == u {
		return "", services.ERR_N_LOGIN
	}
	_i_time, err := strconv.Atoi(time)
	if nil != err {
		return "", ERR_INVALID_TIME
	}
	if nil != helpers.Check_id_format(event_id) {
		return "", ERR_INVALID_EVENT_ID
	}
	event, err := models.Event_get_by_id(event_id)
	if nil != err {
		return "", ERR_EVENT_NOT_FOUND
	}
	if !event.Votable {
		return "", ERR_EVENT_NOT_VOTABLE
	}
	if int64(_i_time) < event.StartDate.Unix() || int64(_i_time) > event.EndDate.Unix() {
		return "", ERR_INVALID_TIME_RANGE
	}
	{ //check if user is invited{{{

		var found bool
		for _, i := range event.Invitations {
			if i.InviteeId == u.Id && i.Status == models.INV_STATUS_ACCEPTED {
				found = true
				break
			}
		}
		if !found {
			return "", ERR_NOT_INVITED
		}
	} //}}}
	//}}}
	return "success", models.Vote_event(event_id, u.Id, int64(_i_time)) //}}}
}
