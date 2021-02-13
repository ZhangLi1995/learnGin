package common

/* 错误码从 10001 开始递增 */
var (
	DBError = &BizError{errorCode: 10001, errorMsg: "操作数据库出错"}
)

type BizError struct {
	errorCode int
	errorMsg  string
}

func (e *BizError) Error() string {
	return e.errorMsg
}

func (e *BizError) ErrMsg() string {
	return e.errorMsg
}

func (e *BizError) ErrCode() int {
	return e.errorCode
}
