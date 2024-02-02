package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func Encrypt(key, text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	plaintext := []byte(text)

	// Pad the plaintext to be a multiple of the block size
	blockSize := block.BlockSize()
	plaintext = PKCS7Padding(plaintext, blockSize)

	// Generate a random IV (Initialization Vector)
	ciphertext := make([]byte, blockSize+len(plaintext))
	iv := ciphertext[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Encrypt the padded plaintext
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[blockSize:], plaintext)

	// Return the base64-encoded ciphertext
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(key, text string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Extract the IV from the ciphertext
	blockSize := block.BlockSize()
	iv := ciphertext[:blockSize]
	ciphertext = ciphertext[blockSize:]

	// Decrypt the ciphertext
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// Unpad the decrypted plaintext
	plaintext, err := PKCS7Unpadding(ciphertext)
	if err != nil {
		return "", err
	}

	// Return the plaintext
	return string(plaintext), nil
}

func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func PKCS7Unpadding(data []byte) ([]byte, error) {
	length := len(data)
	unpadding := int(data[length-1])
	if unpadding > length {
		return nil, fmt.Errorf("Invalid padding")
	}
	return data[:(length - unpadding)], nil
}

// func main() {
// 	key := "0123456789abcdef0123456789abcdef" // 32 bytes for AES-256
// 	sessionID := "SPSE_SESSION=f7e35221c548bda6181b311919f6b0a2794e49c0-___AT=1ee64210fef71b7cc011f0154b130060aef18b8a&___TS=1706885405685&___ID=e9e93ff0-d039-445a-8f38-4d70f3b362fe"

// 	// Encrypt
// 	encryptedText, err := encrypt(key, sessionID)
// 	if err != nil {
// 		fmt.Println("Encryption error:", err)
// 		return
// 	}
// 	fmt.Println("Encrypted:", encryptedText)
// }
