package fileplus

import (
	"os"
	"path/filepath"
	"time"
)

func ErrorWriteFile(content string) error {
	mulu := "errLog/"
	filename := mulu + "" + time.Now().Format("2006-01-02") + ".txt"

	err := WriteToFileWithCreate(filename, content)
	if err != nil {
		return err
	}

	return nil
}

func WriteToFileWithCreate(path string, content string) error {
	var err error
	err = os.MkdirAll(filepath.Dir(path), 0755) // 创建目录
	if err != nil {
		return err
	}

	err = AppendToFile(path, content)
	if err != nil {
		return err
	}

	return nil
}
