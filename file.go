package go_helper

import (
	"os"
)

//检查文件或者文件夹是否存在, 存在返回true, 否则返回false
func FileOrDirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func CreateDir(path string, perm os.FileMode) error {
	err := os.MkdirAll(path, perm)
	if err != nil {
		return err
	}
	return nil
}

func DelFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

func WriteFile(path string, content []byte) error {
	f, err := CreateFile(path)
	if err != nil {
		return err
	}
	_, err = f.Write(content)
	if err != nil {
		return err
	}
	f.Close()
	return nil
}
