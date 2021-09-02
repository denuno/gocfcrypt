package controller

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestAdd(t *testing.T){

	fmt.Println("Starting the crypt...")
	key, err := ioutil.ReadFile("./keyfile")
	if err != nil {
		panic(err.Error())
	}
	encrypted:=OpensslEncrypt(string(key),  "hello world")
	fmt.Println(encrypted)
	fmt.Println(OpensslDecrypt(string(key),encrypted))

}

