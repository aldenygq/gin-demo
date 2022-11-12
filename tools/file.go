package tools

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

//目录是否存在
func IsDir(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

/**
创建文件夹
*/
func CreateDir(dirName string) bool {
	err := os.Mkdir(dirName, 755)
	if err != nil {
		return false
	}
	return true
}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
//golang判断文件或文件夹是否存在的方法为使用os.Stat()函数返回的错误值进行判断:
//1、如果返回的错误为nil,说明文件或文件夹存在
//2、如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
//3、如果返回的错误为其它类型,则不确定是否在存在,建议按照错误处理
func IsExist(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ToString(filePath string) (string, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

//remove file name trailing spaces
func ToTrimString(filePath string) (string, error) {
	str, err := ToString(filePath)
	if err != nil {
		return "", err
	}
	return str, nil
}

func ReadFileInfo(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(fd), nil
}

func ReadFileToStruct(file string, data interface{}) error {
	filecontent, err := ReadFileInfo(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(filecontent), &data)
	if err != nil {
		return err
	}

	return nil
}

func CopyFile(dstfile, srcfile string) error {
	src, err := os.Open(srcfile)
	if err != nil {
		return err
	}

	defer src.Close()

	dst, err := os.Create(dstfile)
	if err != nil {
		return err
	}
	defer dst.Close()
	_,err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	return nil
}
