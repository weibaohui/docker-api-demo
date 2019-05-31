package main

import (
	"flag"
	"fmt"
)

func main() {
	var tarPath *string
	var dst *string
	tarPath = flag.String("tar", "", "请输入tar文件位置")
	dst = flag.String("dst", "", "请输入解压文件夹位置")
	flag.Parse()
	fmt.Println("tar文件位置", *tarPath)
	fmt.Println("解压文件夹位置", *dst)

	if *tarPath == "" || *dst == "" {
		fmt.Println("参数不全")
		return
	}

	e := UnTar(*tarPath, *dst)
	if e != nil {
		fmt.Println(e.Error())
	}
	tars := ListAllTars(*dst, "tar")
	for _, v := range tars {
		fmt.Println(v)
		Push(v)
	}
}
