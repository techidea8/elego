package utils

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"golang.org/x/crypto/pkcs12"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)


//HTTPGet get 请求
func HTTPGet(uri string) ([]byte, error) {
	response, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

//HTTPPost post 请求
func HTTPPost(uri string, data string) ([]byte, error) {
	body := bytes.NewBuffer([]byte(data))
	response, err := http.Post(uri, "", body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

//PostJSON post json 数据请求
func PostJSON(uri string, obj interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u003c"), []byte("<"))
	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u003e"), []byte(">"))
	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u0026"), []byte("&"))
	body := bytes.NewBuffer(jsonData)
	response, err := http.Post(uri, "application/json;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

// PostJSONWithRespContentType post json数据请求，且返回数据类型
func PostJSONWithRespContentType(uri string, obj interface{}) ([]byte, string, error) {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return nil, "", err
	}

	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u003c"), []byte("<"))
	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u003e"), []byte(">"))
	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u0026"), []byte("&"))

	body := bytes.NewBuffer(jsonData)
	response, err := http.Post(uri, "application/json;charset=utf-8", body)
	if err != nil {
		return nil, "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	contentType := response.Header.Get("Content-Type")
	return responseData, contentType, err
}

//PostFile 上传文件
func PostFile(fieldname, filename, uri string) ([]byte, error) {
	fields := []MultipartFormField{
		{
			IsFile:    true,
			Fieldname: fieldname,
			Filename:  filename,
		},
	}
	return PostMultipartForm(fields, uri)
}

//MultipartFormField 保存文件或其他字段信息
type MultipartFormField struct {
	IsFile    bool
	Fieldname string
	Value     []byte
	Filename  string
}

//PostMultipartForm 上传文件或其他多个字段
func PostMultipartForm(fields []MultipartFormField, uri string) (respBody []byte, err error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	for _, field := range fields {
		if field.IsFile {
			fileWriter, e := bodyWriter.CreateFormFile(field.Fieldname, field.Filename)
			if e != nil {
				err = fmt.Errorf("error writing to buffer , err=%v", e)
				return
			}

			fh, e := os.Open(field.Filename)
			if e != nil {
				err = fmt.Errorf("error opening file , err=%v", e)
				return
			}
			defer fh.Close()

			if _, err = io.Copy(fileWriter, fh); err != nil {
				return
			}
		} else {
			partWriter, e := bodyWriter.CreateFormField(field.Fieldname)
			if e != nil {
				err = e
				return
			}
			valueReader := bytes.NewReader(field.Value)
			if _, err = io.Copy(partWriter, valueReader); err != nil {
				return
			}
		}
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, e := http.Post(uri, contentType, bodyBuf)
	if e != nil {
		err = e
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	respBody, err = ioutil.ReadAll(resp.Body)
	return
}

//PostXML perform a HTTP/POST request with XML body
func PostXML(uri string, obj interface{}) ([]byte, error) {
	xmlData, err := xml.Marshal(obj)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(xmlData)
	response, err := http.Post(uri, "application/xml;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http code error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

//httpWithTLS CA证书
func httpWithTLS(rootCa, key string) (*http.Client, error) {
	var client *http.Client
	certData, err := ioutil.ReadFile(rootCa)
	if err != nil {
		return nil, fmt.Errorf("unable to find cert path=%s, error=%v", rootCa, err)
	}
	cert := pkcs12ToPem(certData, key)
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	tr := &http.Transport{
		TLSClientConfig:    config,
		DisableCompression: true,
	}
	client = &http.Client{Transport: tr}
	return client, nil
}

//pkcs12ToPem 将Pkcs12转成Pem
func pkcs12ToPem(p12 []byte, password string) tls.Certificate {
	blocks, err := pkcs12.ToPEM(p12, password)
	defer func() {
		if x := recover(); x != nil {
			log.Print(x)
		}
	}()
	if err != nil {
		panic(err)
	}
	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}
	cert, err := tls.X509KeyPair(pemData, pemData)
	if err != nil {
		panic(err)
	}
	return cert
}

//PostXMLWithTLS perform a HTTP/POST request with XML body and TLS
func PostXMLWithTLS(uri string, obj interface{}, ca, key string) ([]byte, error) {
	xmlData, err := xml.Marshal(obj)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(xmlData)
	client, err := httpWithTLS(ca, key)
	if err != nil {
		return nil, err
	}
	response, err := client.Post(uri, "application/xml;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http code error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}

var filedatamap sync.Map
func loadfiledata(filepath string)([]byte,error){
	data,ok := filedatamap.Load(filepath);
	if(ok){
		return data.([]byte),nil
	}else{
		databytes,err := ioutil.ReadFile(filepath)
		if err==nil{
			filedatamap.Store(filepath,databytes)
		}
		return databytes,err
	}

}
//从缓存中读取加密解密文件
func MakeCertHttpClientUseP12(p12file, passwd string) (client *http.Client,err error) {
	certdata,err := loadfiledata(p12file);
	if err!=nil {
		return
	}
	client, err = buildClient(certdata, passwd)
	return

}

//wxPayBuildClient 构建带证数的httpClient
func buildClient(data []byte, mchId string) (client *http.Client, err error) {
	// 将pkcs12证书转成pem
	cert, err := pkc12ToPerm(data, mchId)
	if err != nil {
		return
	}
	// tls配置
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	// 带证书的客户端
	client = &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout:     3 * time.Minute,
			TLSHandshakeTimeout: 10 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 10 * time.Minute,
				DualStack: true,
			}).DialContext,
			TLSClientConfig:    config,
			DisableCompression: true,
		},
	}
	return
}

//wxPayPkc12ToPerm 证数格式转化
func pkc12ToPerm(data []byte, mchId string) (cert tls.Certificate, err error) {
	blocks, err := pkcs12.ToPEM(data, mchId)
	if err != nil {
		return
	}
	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}
	cert, err = tls.X509KeyPair(pemData, pemData)
	return
}


func MakeCertHttpClientUseKey(certpath,keypath,capath string) (client *http.Client,err error) {

	certs, e := tls.LoadX509KeyPair(certpath, keypath)
	if e != nil {
		//fmt.Println(errorx.Wrap(e).Error())
		return nil,e
	}

	rootCa, e := ioutil.ReadFile(capath)
	if e != nil {
		//fmt.Println(errorx.Wrap(e))
		return nil,e
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(rootCa)
	c := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      pool,
				Certificates: []tls.Certificate{certs},
			},
		},
	}
	return c,nil
}