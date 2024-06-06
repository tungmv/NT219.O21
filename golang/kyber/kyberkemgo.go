package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"

	kyberk2so "github.com/symbolicsoft/kyber-k2so"
)

func main() {
	// Generate keypair
	// short variable declaration, initializes privateKey, publicKey, and err
	privateKey, publicKey, err := kyberk2so.KemKeypair768()
	if err != nil {
		log.Fatalf("Error generating keypair: %v", err)
	}

	// Encrypt
	ciphertext, ssA, err := kyberk2so.KemEncrypt768(publicKey)
	if err != nil {
		log.Fatalf("Error encrypting: %v", err)
	}

	// Decrypt
	ssB, err := kyberk2so.KemDecrypt768(ciphertext, privateKey)
	if err != nil {
		log.Fatalf("Error decrypting: %v", err)
	}

	// Use ssA and ssB by writing them to files and printing them
	writeToFile("ssA.hex", hex.EncodeToString(ssA[:]))
	writeToFile("ssB.hex", hex.EncodeToString(ssB[:]))
	writeToFile("ssA.base64", base64.StdEncoding.EncodeToString(ssA[:]))
	writeToFile("ssB.base64", base64.StdEncoding.EncodeToString(ssB[:]))

	fmt.Printf("Shared secret A (hex): %s\n", hex.EncodeToString(ssA[:]))
	fmt.Printf("Shared secret B (hex): %s\n", hex.EncodeToString(ssB[:]))
	fmt.Printf("Shared secret A (base64): %s\n", base64.StdEncoding.EncodeToString(ssA[:]))
	fmt.Printf("Shared secret B (base64): %s\n", base64.StdEncoding.EncodeToString(ssB[:]))

	// Write keys and ciphertext to files
	writeToFile("private_key.hex", hex.EncodeToString(privateKey[:]))
	writeToFile("public_key.hex", hex.EncodeToString(publicKey[:]))
	writeToFile("ciphertext.hex", hex.EncodeToString(ciphertext[:]))
	writeToFile("private_key.base64", base64.StdEncoding.EncodeToString(privateKey[:]))
	writeToFile("public_key.base64", base64.StdEncoding.EncodeToString(publicKey[:]))
	writeToFile("ciphertext.base64", base64.StdEncoding.EncodeToString(ciphertext[:]))

	fmt.Println("Keys, ciphertext, and shared secrets have been written to files.")
}

// Helper function to write data to a file
func writeToFile(filename, data string) {
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		log.Fatalf("Error writing to file %s: %v", filename, err)
	}
}
