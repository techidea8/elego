package utils

import (
	"encoding/xml"
)



func Xml2Map(xmlstr string) (result map[string]string, err error) {
	stringMap := make(map[string]string,0)
	err = xml.Unmarshal([]byte(xmlstr), (*map[string]string)(&stringMap))
	return stringMap, err
}



