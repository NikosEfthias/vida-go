__api="http://localhost:8080"
export tkn=""
export last_result=""
_login()
{
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
	last_result=`curl $__api$ep -d "email=$mail&password=$pass"`
	tkn=`echo $last_result|jq '.data'| cut -d'"' -f2`
	echo $last_result
}

_get()
{
	_login &>/dev/null || {echo $last_result && return 1}
	ep="$api/api/user/$tkn"
	last_result=`curl $ep`
	echo $last_result
}

_update()
{
	test -z $1 && echo missing field && return 1
	test -z $2 && echo missing value && return 1
	_login &>/dev/null || {echo $last_result && return 1}
	ep="$api/api/user/update/$tkn/$1"
	last_result=`curl $ep -d "value=$2"`
	echo $last_result
}
_pp()
{
	test -z $1 && echo missing filename && return 1
	_login &>/dev/null || {echo $last_result && return 1}
	ep="/api/user/pp/$tkn"
	last_result=`curl "$api$ep" -F "file=@$1"`
	echo $last_result
}
