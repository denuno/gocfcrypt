package gocfcrypt

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestEncrypt(t *testing.T){
	fmt.Println("Test Encrypt ***")
	key := ReadFile("./keyfile")
	secret := Base64Decode(key)
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

func TestDecrypt(t *testing.T){
	fmt.Println("Test Decrypt ***")
	key := ReadFile("./keyfile")
	secret := Base64Decode(key)

	data := ReadFile("./encrypted.txt")
	result, err := Decrypt(string(data), secret)
	fmt.Println("decrypted result", result)
	fmt.Println("decryption err", err)
	fmt.Println("")
}

func TestAgain(t *testing.T){
	fmt.Println("Test EN Decrypt ***")
	key := ReadFile("./keyfile")
	secret := Base64Decode(key)

	plaintext  := ReadFile("./plaintext.txt")
	encResult, err := Encrypt(plaintext, secret)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	fmt.Println("encResult", encResult)

	data := ReadFile("./encrypted.txt")
	result, err := Decrypt(string(data), secret)
	fmt.Println("decrypted result", result)
	fmt.Println("decryption err", err)

	dataGo := ReadFile("./goencrypted.txt")
	resultGo, err := Decrypt(string(dataGo), secret)
	fmt.Println("decrypted result", resultGo)
	fmt.Println("decryption err", err)
	fmt.Println("")
}
