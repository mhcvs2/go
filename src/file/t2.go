package main

import (
	"path/filepath"
	"os"
	"fmt"
	"errors"
)

//GetSubFiles 列出文件夹下的所有文件，不包括子文件夹下的文件
func GetSubFiles(root string) ([]string, error) {
	res := []string{}
	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if path == root {
			if f.IsDir() {
				return nil
			} else {
				return errors.New(fmt.Sprintf("%s is not a directory", path))
			}
		}
		if f.IsDir(){
			return filepath.SkipDir
		}
		res = append(res, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func t21() {
	l, err := GetSubFiles("/root/.bash_history")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(l)
	}
}

func main() {
	t21()
}