package ctrl

import (
	"fmt"
	"github.com/techidea8/restgo/middleware"
	"testing"
)


func  TestRespUserInfo(t *testing.T) {
	token, err := middleware.CreateToken(3)
	if err!=nil{
		return
	}
	fmt.Println(token)
}
