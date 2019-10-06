package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

//<GO语言获取目录列表用 ioutil.ReadDir()，遍历目录用 filepath.Walk()>

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}

	return files, err
}

//遍历获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func AllListDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录

		if fi.IsDir() { // 忽略目录
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}

		return nil
	})

	return files, err
}

//分离目录文件名后缀文件名函数
/*
value	D:\yun\Go_项目生成\zz2\zz2_5\zz2_5.exe
n=0 D:\yun\Go_项目生成\zz2\zz2_5\
n=1 zz2_5.exe
n=2 D:\yun\Go_项目生成\zz2\zz2_5\zz2_5
n=3 .exe
*/
func Paths(value string, n int) string {
	var files string
	suffix := path.Ext(value)
	path_no_suffix := strings.TrimSuffix(value, suffix)
	path, _ := filepath.Split(value)
	filename := filepath.Base(value)
	if n == 0 {
		files = path
	} else if n == 1 {
		files = filename

	} else if n == 2 {
		files = path_no_suffix
	} else if n == 3 {
		files = suffix
	} else {
		files = ""
	}

	return files
}

func main() {
	allfiles, _ := AllListDir(`C:\Users\DELL\Desktop\翻译项目`, ".txt")
	for _, value := range allfiles {
		fmt.Println(value)
		fmt.Println(Paths(value, 0))
		fmt.Println(Paths(value, 1))
		fmt.Println(Paths(value, 2))
		fmt.Println(Paths(value, 3))
	}
}
