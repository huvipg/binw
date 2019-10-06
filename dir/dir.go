package dir

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Get(s string) string {

	resp, err := http.Get(s)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
	return string(body)
}
func A_file(filepath, str_content string) {
	//创建或打开一个文件,追加写入
	fd, _ := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	buf := []byte(str_content + "\n")
	fd.Write(buf)
	fd.Close()
}

//重新创建一个文件,并写入
func W_file(userFile, str_content string) {
	fd, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fd.Close()
	fd.WriteString(str_content + "\n")
}

//读文件打印成字符串
func R_file(filepath string) string {

	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("read file fail", err)
		return ""
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return ""
	}

	return string(fd)
}

//是否目录或文件
func Ispath(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//创建多级目录
func Mkdir(path string) {
	if !Ispath(path) {
		//创建多级目录 os.Mkdir 单目录
		os.MkdirAll(path, os.ModePerm)
	}
}

//windows 目录\替换成 /  linux
func Ren(path string) string {
	return strings.Replace(path, "\\", "/", -1)
}
