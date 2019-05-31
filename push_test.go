package main

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	//Push("/tmp/tt/golang.tar")
	//Push("/tmp/test.tar")

	e := UnTar("/tmp/test.tar", "/tmp/test/")
	if e != nil {
		fmt.Println(e.Error())
	}
	tars := ListAllTars("/tmp/test/", "tar")
	for _, v := range tars {
		fmt.Println(v)
		Push(v)
	}
}
