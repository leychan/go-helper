package go_helper

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
)

//将指定路径的图片转换成base64 URLencode数据
func base64encodeImg(path string, t string) (string, error) {
	if !FileOrDirExists(path) {
		return "", errors.New("file is not exist")
	}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer file.Close()
	content, _ := ioutil.ReadAll(file)
	if t == "json" || t == "JSON" {
		//如果是json传递,则不需要进行url.QueryEscape, 否则需要url.QueryEscape(s)
		s := base64.StdEncoding.EncodeToString(content)
		return s, nil
	}
	s := base64.StdEncoding.EncodeToString(content)
	return url.QueryEscape(s), nil
}

//用作Json提交的格式, 不需要进行url.QueryEscape
func JsonBase64encodeImg(path string) (string, error) {
	return base64encodeImg(path, "json")
}

//用作表单提交的格式, 需要进行url.QueryEscape
func FormBase64encodeImg(path string) (string, error) {
	return base64encodeImg(path, "form")
}

//解码base64 URLencode数据成字符串
func Base64decodeImg(str string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(str)

	if err != nil {
		return "", err
	}
	return string(b), nil
}
