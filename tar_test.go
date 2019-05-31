package main

import (
	"fmt"
	"testing"
)

func TestDeCompress(t *testing.T) {
	e := UnTar("/tmp/test.tar", "/tmp/test/")
	if e != nil {
		fmt.Println(e.Error())
	}
}

func TestListAllTars(t *testing.T) {
	ListAllTars("/tmp/1", "zip")

}
