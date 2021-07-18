package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

//http请求的client
var client *http.Client

//初始化 http连接信息
func init() {
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
		},
	}
}

//httpGet 发起get请求
func httpGet(apiurl string) ([]byte, error) {
	resp, errGet := client.Get(apiurl)

	if errGet != nil {
		return nil, errGet
	}
	defer resp.Body.Close()

	body, errRead := ioutil.ReadAll(resp.Body)
	return body, errRead
}

//httpPost post请求,返回原始的字节数组
func httpPost(apiurl string, params map[string]interface{}) ([]byte, error) {
	//data := make(url.Values)
	//for k, v := range params {
	//	data.Add(k, v)
	//}
	byteparams, errparams := json.Marshal(params)
	if errparams != nil {
		return nil, errparams
	}
	resp, errPost := client.Post(apiurl, "application/json", bytes.NewReader(byteparams))
	if errPost != nil {
		return nil, errPost
	}
	defer resp.Body.Close()

	body, errRead := ioutil.ReadAll(resp.Body)

	return body, errRead
}

//httpPostXml 发送Post请求,参数是XML格式的字符串
func httpPostXml(url string, xmlBody string) (body []byte, err error) {
	resp, err := client.Post(url, "application/xml", strings.NewReader(xmlBody))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

//httpPostXmlWithCert 发送带证书的Post请求,参数是XML格式的字符串
func httpPostXmlWithCert(url string, xmlBody string, client *http.Client) (body []byte, err error) {
	resp, err := client.Post(url, "application/xml", strings.NewReader(xmlBody))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

//httpPostXmlWithCert 发送带证书的Post请求,参数是XML格式的字符串
func HttpPostXmlWithCert(url string, xmlBody string, client *http.Client) (body []byte, err error) {
	return httpPostXmlWithCert(url, xmlBody, client)
}



func WxPayGetCertHttpClientUsep12(p12file,mchId string) (*http.Client, error) {
	return MakeCertHttpClientUseP12(p12file,mchId)
}
func WxPayGetCertHttpClientUseKey(certpath,keypath,rootcapath string) (*http.Client, error) {
	return MakeCertHttpClientUseKey(certpath,keypath,rootcapath)
}
