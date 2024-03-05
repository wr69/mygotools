package fileplus

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func CreateFileIfNotExists(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 文件不存在，创建文件
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		if _, err := os.Create(filePath); err != nil {
			return err
		}

		//fmt.Printf("Created file: %s\n", filePath)
	}
	return nil
}

func GetFile(filePath string) ([]byte, error) {
	if err := CreateFileIfNotExists(filePath); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ReadFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteByteFile(filePath string, content []byte) error {
	if err := ioutil.WriteFile(filePath, content, 0644); err != nil {
		return err
	}
	//sfmt.Printf("Write to file: %s\n", filePath)
	return nil
}

func WriteFile(filePath string, content string) error {
	if err := ioutil.WriteFile(filePath, []byte(content), 0644); err != nil {
		return err
	}
	//sfmt.Printf("Write to file: %s\n", filePath)
	return nil
}

func AppendToFile(filePath string, content string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write([]byte(content + "\n")); err != nil {
		return err
	}

	//fmt.Println("Appended content to file:", content)
	return nil
}
