package cfencrypt

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
	encrypt(string(key), "plaintext")
	decrypt(string(key), "plaintext.aes")

}

