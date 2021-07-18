package utils

import (
	"fmt"
	"testing"
)


var xmlstr1 = `<xml><data1>test1</data1></xml>`
func TestXml2Map(t *testing.T) {
	r,_ := Xml2Map(xmlstr1)
	fmt.Println(r)
}



