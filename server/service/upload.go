package service

import (
	"fmt"
	"math/rand"
	"mime/multipart"
	"path"
	"strings"
	"turingapp/global"

	uuid "github.com/satori/go.uuid"
)

type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

func NewOss() OSS {
	switch global.Conf.System.OssType {
	case "local":
		return &Local{
			global.Conf.Local,
		}
	case "aliyun":
		return &Alioss{
			global.Conf.Aliyun,
		}
	default:
		return &Local{
			global.Conf.Local,
		}
	}
}

func Fikekey(file *multipart.FileHeader) string {
	dirstr := "0123456789abcdef"
	suffix := path.Ext(file.Filename)
	str := strings.Replace(uuid.NewV4().String(), "-", "", -1)
	idx := rand.Intn(15)
	c1 := dirstr[idx : idx+1]
	idx = rand.Intn(15)
	c2 := dirstr[idx : idx+1]
	fileKey := fmt.Sprintf("%s/%s/%s%s%s%s", c1, c2, c1, c2, str, suffix) // 文件名格式 自己可以改 建议保证唯一性
	return fileKey
}
