package utils

import (
	"bytes"
	"fmt"
)

func Map2Xml(mapdata map[string]string) (result string, err error) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("<xml>")
	for k, v := range mapdata {
		buffer.WriteString(fmt.Sprintf("<%s><![CDATA[%v]]></%s>", k, v, k))
	}
	buffer.WriteString("</xml>")
	return buffer.String(),nil
}
