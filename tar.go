package main

import (
	"fmt"
	"github.com/artyom/untar"
	"os"
	"path/filepath"
	"strings"
)

func UnTar(src, dst string) (err error) {
	if ExistDir(dst) {
		os.RemoveAll(dst)
	}
	file, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = untar.Untar(file, dst)
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func ListAllTars(dst, suffix string) []string {
	var paths []string
	err := filepath.Walk(dst, func(path string, info os.FileInfo, err error) error {

		if strings.HasSuffix(info.Name(), suffix) {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		println(err.Error())
	}

	return paths

}

// 判断目录是否存在
func ExistDir(dirname string) bool {
	fi, err := os.Stat(dirname)
	return (err == nil || os.IsExist(err)) && fi.IsDir()
}
