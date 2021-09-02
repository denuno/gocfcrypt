package crypt

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestEncrypt(t *testing.T){
	fmt.Println("Test Encrypt ***")
	key, err := ioutil.ReadFile("./keyfile")
	secret, err := base64.StdEncoding.DecodeString(strings.TrimSpace(string(key)))
	data := []byte("this is a test from Go")
	encResult, err := Encrypt(data, secret)
	fmt.Println("encResult:", key)
	fmt.Println("encResult:", encResult)
	if err != nil {
		fmt.Println("enc err", err)
	}

	fmt.Println("writing ./goencrypted.txt")
	writeErr := ioutil.WriteFile("./goencrypted.txt", []byte(encResult), 0644)
	if writeErr != nil {
		return
	}

	result, err := Decrypt(encResult, secret)
	fmt.Println("decrypted result:", result)
	fmt.Println("decryption err:", err)
	fmt.Println("")
}

func TestDncrypt(t *testing.T){
	fmt.Println("Test Decrypt ***")
	key, err := ioutil.ReadFile("./keyfile")
	secret, err := base64.StdEncoding.DecodeString(strings.TrimSpace(string(key)))

	data, err := ioutil.ReadFile("./encrypted.txt")
	result, err := Decrypt(string(data), secret)
	fmt.Println("decrypted result", result)
	fmt.Println("decryption err", err)
	fmt.Println("")

}

func TestEDncrypt(t *testing.T){
	fmt.Println("Test EN Decrypt ***")
	key, err := ioutil.ReadFile("./keyfile")
	secret, err := base64.StdEncoding.DecodeString(strings.TrimSpace(string(key)))

	plaintext, err := ioutil.ReadFile("./plaintext.txt")
	encResult, err := Encrypt(plaintext, secret)
	fmt.Println("encResult", encResult)

	data, err := ioutil.ReadFile("./encrypted.txt")
	result, err := Decrypt(string(data), secret)
	fmt.Println("decrypted result", result)
	fmt.Println("decryption err", err)

	dataGo, err := ioutil.ReadFile("./goencrypted.txt")
	resultGo, err := Decrypt(string(dataGo), secret)
	fmt.Println("decrypted result", resultGo)
	fmt.Println("decryption err", err)
	fmt.Println("")

}
