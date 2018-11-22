__api="http://localhost:8080"
export tkn=""
export last_result=""
_register()
{ #{{{
	ep="/api/user/register"
	api=""
	test $API_ADDR && api=$API_ADDR || api=$__api
	last_result=`curl $api$ep -d "email=nikos@mugsoft.io" \
		-d "password=test" \
		-d "name=nikos" \
		-d "lastname=efthias" \
		-d "phone=0079600757769" 2>/dev/null`
	echo $last_result
} #}}}
_login()
{ #{{{
	ep="/api/user/login"
	api=""
	test $API_ADDR && api=$API_ADDR || api=$__api
	if test -z $mail
	then
		test -z $1 && mail="nikos@mugsoft.io" || mail=$1
	fi
	if test -z $pass
	then
		test -z $2 && pass="test" || pass=$2
	fi
	last_result=`curl $api$ep -d "email=$mail&password=$pass" 2>/dev/null`
	tkn=`echo $last_result|jq '.data'| cut -d'"' -f2`
	echo $last_result
} #}}}

_get()
{ #{{{
	_login &>/dev/null || (echo $last_result && return 1) 
	ep="$api/api/user/$tkn"
	last_result=`curl $ep 2>/dev/null`
	echo $last_result
} #}}}
_update()
{ #{{{
	test -z $1 && echo missing field && return 1
	test -z $2 && echo missing value && return 1
	_login &>/dev/null || (echo $last_result && return 1 )
	ep="$api/api/user/update/$tkn/$1"
	last_result=`curl $ep -d "value=$2" 2>/dev/null`
	echo $last_result
} #}}}
_pp()
{ #{{{
	test -z $1 && echo missing filename && return 1
	_login &>/dev/null || (echo $last_result && return 1 )
	ep="/api/user/pp/$tkn"
	last_result=`curl "$api$ep" -F "file=@$1" 2>/dev/null`
	echo $last_result
} #}}}
_invite_app()
{ #{{{
	_login &>/dev/null || (echo $last_result && return 1 )
	_ep="$api/api/app/invite/$tkn"
	echo $_ep 1>&2
	test -z $1 && mail_addr="nikos@mugsoft.io"||mail_addr=$1
	echo sending invitations to $mail_addr
	last_result=`curl $_ep -d "invitees=$mail_addr" 2>/dev/null`
	echo $last_result
} #}}}
_event_create()
{ #{{{
	test -z $1 && echo missing filename && return 1
	_login &>/dev/null || (echo $last_result && return 1 )
	ep="/api/event/create/$tkn"
	dt=$(($(date +%s)+32000))
	last_result=`curl $api$ep \
		-F "image=@$1" \
		-F "title=test_event" \
		-F "location=here" \
		-F "start_date=$dt" \
		-F "end_date="$((dt+5000))\
		-F "details=this is a test event" \
		-F "max_num_guest=10" \
		-F "min_num_guest=0" \
		-F "votable=false" \
		-F "cost=10.2" \
		2>/dev/null`
	echo $last_result
} #}}}
_event_delete()
{ #{{{
	if [ -z $2 ]
	then
		_event_create $1 &>/dev/null || return 1;
		id=`echo $last_result | jq '.data' | xargs echo`
	else
		_login &>/dev/null || (echo $last_result && return 1 )
		id=$2
	fi
	_ep="$api/api/event/delete/$id/$tkn"
	echo $_ep 1>&2
	last_result=`curl $_ep 2>/dev/null`
	echo $last_result
} #}}}
_event_get_by_id() 
{ #{{{
	if [ -z $2 ]
	then
		_event_create $1 &>/dev/null || return 1;
		id=`echo $last_result | jq '.data' | xargs echo`
	else
		_login &>/dev/null || (echo $last_result && return 1 )
		id=$2
	fi
	_ep="$api/api/event/byid/$id/$tkn"
	echo $_ep 1>&2
	last_result=`curl $_ep 2>/dev/null`
	echo $last_result
} #}}}
