package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func Md5encode(str string) string {

	d := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", d)
}
func Md5Encode(str string) string {
	return strings.ToUpper(Md5encode(str))
}
