package service

import (
	"mime/multipart"
	"turingapp/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Alioss struct {
	config.Aliyun
}

//@author: [winlion](https://github.com/winlion)

func (aliyun *Alioss) UploadFile(file *multipart.FileHeader) (string, string, error) {

	fileKey := Fikekey(file)
	url := aliyun.Location + "/" + fileKey
	//aliyun.AccessKey, aliyun.SecretKey
	client, err := oss.New(aliyun.Endpoint, aliyun.AccessKey, aliyun.SecretKey)
	if err != nil {
		return url, fileKey, err
	}

	bucket, err := client.Bucket(aliyun.Bucket)
	if err != nil {
		return url, fileKey, err
	}

	f, err := file.Open()
	if err != nil {
		return url, fileKey, err

	}
	err = bucket.PutObject(fileKey, f)
	if err != nil {
		return url, fileKey, err

	}

	return url, fileKey, nil
}

func (aliyun *Alioss) DeleteFile(key string) error {
	client, err := oss.New(aliyun.Endpoint, aliyun.AccessKey, aliyun.SecretKey)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(aliyun.Bucket)
	if err != nil {
		return err
	}
	err = bucket.DeleteObject(key)
	if err != nil {
		return err
	}
	return nil
}
