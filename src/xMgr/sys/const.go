package sys

//websocket
const (
	CMD_SUCCESS         = 0
	CMD_LOGIN           = 1
	CMD_QUERY_GS        = 2
	CMD_QUERY_SQL       = 3
	CMD_CHAT            = 4
	CMD_CHECK_SESSION   = 5
	CMD_QUERY_USER      = 6
	CMD_DELETE_USER     = 7
	CMD_EDIT_USER       = 8
	CMD_ADD_USER        = 9
	CMD_ADD_GS          = 12
	CMD_DELETE_GS       = 13
	CMD_EDIT_GS         = 14
	CMD_QUERY_MANAGE_GS = 15

	CMD_ERROR_MSG = 10 //S->C
	///////////////////////////////////////////////
	// nCmd<1001 is websocket; nCmd >1001 is http
	CMD_HTTP_UPLOAD    = 1001 //upload
	CMD_HTTP_EXEC_FILE = 1002 //execFile
)

//http
const (
	HTTP_ERR        = 10 //err
	HTTP_SUCCEED    = 1  //succeed
	HTTP_SYSTEM_ERR = 0  //system err
)

//user privilige
const (
	USER_GTOUP_ID = 1 //user_group_id
)

///////////////////////////////////////////////
//id and name
const (
	/////////////////表的部分字段<1000
	TABLE_S_ITEM_TYPE = 1
	TABLE_S_GS        = 2
	TABLE_S_SDK_TYPE  = 3
	TABLE_S_GS_GROUP  = 4
	/////////////////整张表>=1000
	TABLE_All_S_GS = 1000
)

//ERR_CODE///////////////////////////////////
const (
	CODE_ERR                      = 1
	CODE_SUCCEED                  = 2
	CODE_LOGIN_SUC                = 1000
	CODE_NAME_NOT_EXIST           = 1001
	CODE_PWD_ERR                  = 1002
	CODE_NEED_PASS_PARAM          = 1003
	CODE_NEED_NAME_PARAM          = 1004
	CODE_SQL_NO_EXIST             = 1005
	CODE_DB_GROUP_ID_NO_EXIST     = 1006
	CODE_S_GS_ID_NO_EXIST         = 1007
	CODE_SESSION_NO_EXIST         = 1008
	CODE_SESSION_ERR              = 1009
	CODE_SQL_SYNTXT_ERR           = 1010
	CODE_MANAGE_NOT_AUTHORIZATION = 1011
	CODE_DELETE_USER_FAIL         = 1012
	CODE_ADD_USER_FAIL            = 1013
	CODE_NEED_ID_PARAM            = 1014
	CODE_NEED_TAG_PARAM           = 1015
	CODE_NO_PRIVILIGE             = 1016
	CODE_NEED_INFO_PARAM          = 1017
	CODE_PARAM_ERR                = 1018
	CODE_ROLE_NOT_EXIST           = 1019
	CODE_GS_NOT_EXIST             = 1020
	CODE_FORMAT_ERR               = 1021
	CODE_DATA_EXIST               = 1022
)

///////////////////////////////////////
