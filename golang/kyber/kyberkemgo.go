package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	kyberk2so "github.com/symbolicsoft/kyber-k2so"
)

func main() {
	// This program takes in command line arguments
	if len(os.Args) < 3 {	// Allow wrong numbers of argument to allow the program print more detailed error messages
        log.Fatalf("Usage: %s <mode(genkey/encap/decap)> [filename1] [filename2]\nExample: %s encap cipher.txt pub.bin", os.Args[0], os.Args[0])
    }

	// Get the mode and assign, init to 'mode'
	mode := os.Args[1]

	switch mode {
	case "genkey":
		if len(os.Args) != 4 {
            log.Fatalf("Usage: %s genkey <private_key_file> <public_key_file>", os.Args[0])
        }
        generateKeyPair(os.Args[2], os.Args[3])

	case "encap":
        if len(os.Args) != 4 {
            log.Fatalf("Usage: %s encap <plaintext_file> <public_key_file>", os.Args[0])
        }
        encapsulate(os.Args[2], os.Args[3])

    case "decap":
        if len(os.Args) != 4 {
            log.Fatalf("Usage: %s decap <ciphertext_file> <private_key_file>", os.Args[0])
        }
        decapsulate(os.Args[2], os.Args[3])

    default:
        log.Fatalf("Invalid mode!")
    }
	// End switch case =============
}

// Helper function to write data to a file
func writeToFile(filename, data string) {
	err := ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		log.Fatalf("Error writing to file %s: %v", filename, err)
	}
}

// Helper function to generate kyber key pairs:
func generateKeyPair(privateKeyFile, publicKeyFile string) {
    privateKey, publicKey, err := kyberk2so.KemKeypair768()
    if err != nil {
        log.Fatalf("Failed to generate key pair: %v", err)
    }
    if err := ioutil.WriteFile(privateKeyFile, privateKey[:], 0600); err != nil {
        log.Fatalf("Failed to write private key: %v", err)
    }
    if err := ioutil.WriteFile(publicKeyFile, publicKey[:], 0644); err != nil {
        log.Fatalf("Failed to write public key: %v", err)
    }
    fmt.Println("Keys generated successfully.")
}

// Helper function to encap key
func encapsulate(plaintextFile, publicKeyFile string) {
    publicKey, err := ioutil.ReadFile(publicKeyFile)
    if err != nil {
        log.Fatalf("Failed to read public key: %v", err)
    }
    plaintext, err := ioutil.ReadFile(plaintextFile)
    if err != nil {
        log.Fatalf("Failed to read plaintext: %v", err)
    }
    ciphertext, _, err := kyberk2so.KemEncrypt768(publicKey[:])
    if err != nil {
        log.Fatalf("Failed to encapsulate: %v", err)
    }
    // Assuming the ciphertext is what you need to save
    if err := ioutil.WriteFile("ciphertext.bin", ciphertext[:], 0644); err != nil {
        log.Fatalf("Failed to write ciphertext: %v", err)
    }
    fmt.Println("Message encapsulated successfully.")
}

// Helper function to decap key
func decapsulate(ciphertextFile, privateKeyFile string) {
    privateKey, err := ioutil.ReadFile(privateKeyFile)
    if err != nil {
        log.Fatalf("Failed to read private key: %v", err)
    }
    ciphertext, err := ioutil.ReadFile(ciphertextFile)
    if err != nil {
        log.Fatalf("Failed to read ciphertext: %v", err)
    }
    _, err = kyberk2so.KemDecrypt768(ciphertext[:], privateKey[:])
    if err != nil {
        log.Fatalf("Failed to decapsulate: %v", err)
    }
    fmt.Println("Message decapsulated successfully.")
}