package service

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"turingapp/config"

	log "github.com/cihub/seelog"
	"go.uber.org/zap"
)

type Local struct {
	config.Local
}



func (local *Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀

	// 拼接新文件名
	filename := Fikekey(file)
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(local.Path+"/"+filename[0:3], os.ModePerm)
	if mkdirErr != nil {
		log.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := local.Path + "/" + filename

	f, openError := file.Open() // 读取文件
	if openError != nil {
		log.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		log.Error("function os.Create() Filed", zap.Any("err", createErr.Error()))

		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		log.Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return local.Location + "/" + filename, p, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Local
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error

func (local *Local) DeleteFile(key string) error {

	if err := os.Remove(key); err != nil {
		return errors.New("本地文件删除失败, err:" + err.Error())
	}

	return nil
}
