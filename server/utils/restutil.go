package utils

import (
	"encoding/json"
	"encoding/xml"

	"math/rand"

	"net/url"

	"strings"
	"time"
)

//jsonString 生成模型字符串
func JsonString(m interface{}) string {
	bs, _ := json.Marshal(m)
	return string(bs)
}

//jsonString 生成模型字符串
func XmlString(m interface{}) string {
	bs, _ := xml.Marshal(m)
	return string(bs)
}

//formatDateTime 格式化时间,按照yyyyMMddHHmmss格式
func FormatDateTime(t time.Time, lay string) string {
	return t.Format(lay)
}

//encodePath 对URL进行Encode编码
func EncodePath(u string) (path string, err error) {
	uriObj, err := url.Parse(u)
	if err != nil {
		return
	}
	path = uriObj.EscapedPath()
	return
}

//generateRandomString 获取随机字符串
func FuncGenerateRandomString(length int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	b := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return strings.ToLower(string(result))
}

//pkcs7UnPadding 解密填充模式(去除补全码) pkcs7UnPadding 解密时,需要在最后面去掉加密时添加的填充byte
func Pkcs7UnPadding(plainText []byte) []byte {
	length := len(plainText)
	unpadding := int(plainText[length-1])   // 找到Byte数组最后的填充byte
	return plainText[:(length - unpadding)] // 只截取返回有效数字内的byte数组
}
