package utils

// 业务错误码

const (
	SUCCESS = 0
	ERROR   = 9999
	// code = 1xxx users模块的错误
	ERROR_TOKEN_WRONG           = 1001
	ERROR_TOKEN_FORMAT_WRONG    = 1002
	ERROR_USERNAME_FORMAT_WRONG = 1003

	// code = 2xxx commits模块的错误

	// code = 3xxx objects模块的错误

	// code = 4xxx owners模块的错误

)

// 错误码映射表
var codeMap = map[int]string{
	SUCCESS:                     "OK",
	ERROR:                       "FAIL",
	ERROR_TOKEN_WRONG:           "Token 不正确",
	ERROR_TOKEN_FORMAT_WRONG:    "Token格式错误",
	ERROR_USERNAME_FORMAT_WRONG: "用户名格式错误",
}

func GetErrMsg(code int) string {
	return codeMap[code]
}
