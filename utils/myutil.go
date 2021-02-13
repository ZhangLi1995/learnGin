package utils

import (
	"crypto/md5"
	"fmt"
	"learnGin/common"
	"strings"
)

/* MD5 加密 */
func MD5(str string) string {
	md5Str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5Str
}

/* 字符串判空 */
func StrIsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

func ConvertBizErr(err error) *common.BizError {
	return err.(*common.BizError)
}
