package core

const (
	CODE_OK int = iota
	CODE_DB_CONNECT_ERROR
	CODE_SQL_ERROR
	CODE_INVALID_PARAMS
	CODE_SYSTEM_ERROR
	CODE_NO_USER
	CODE_ERROR_PASSWORD
	CODE_USER_EXIST
)

var errormap = map[int]string{
	CODE_OK:               "OK",
	CODE_DB_CONNECT_ERROR: "db connect error",
	CODE_SQL_ERROR:        "db sql error",
	CODE_INVALID_PARAMS:   "error params",
	CODE_SYSTEM_ERROR:     "system error",
	CODE_NO_USER:          "no user",
	CODE_ERROR_PASSWORD:   "pass error",
	CODE_USER_EXIST:       "user exist",
}

func GetMsg(errorcode int) string {
	return errormap[errorcode]
}
