package main

import (
	"encryption"
	"fmt"
)

func main() {
	fmt.Println("Starting the application...")
	ciphertext := encryption.Encrypt([]byte("Hello World"), "password")
	fmt.Printf("Encrypted: %x\n", ciphertext)
	plaintext := encryption.Decrypt(ciphertext, "password")
	fmt.Printf("Decrypted: %s\n", plaintext)
	encryption.EncryptFile("sample.txt", []byte("Hello World"), "password1")
	fmt.Println(string(encryption.DecryptFile("sample.txt", "password1")))
}
