package lib

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
)

func Encrypt(message, password string, outfile string) (string, error) {
	key := keyDigest(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	plaintext := []byte(message)
	plaintext = PKCS7Padding(plaintext, aes.BlockSize)
	ciphertext := make([]byte, len(plaintext))

	mode := cipher.NewCBCEncrypter(block, key[:aes.BlockSize])
	mode.CryptBlocks(ciphertext, plaintext)

	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	if outfile != "" {
		err = os.WriteFile(outfile, []byte(encoded), 0644)
		if err != nil {
			return "", err
		}
	}
	return encoded, nil
}

// Decrypt the message using AES-256-CBC with the provided password
func Decrypt(encodedMessage, password string, outfile string) (string, error) {
	key := keyDigest(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encodedMessage)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", err //The ciphertext is too short
	}

	mode := cipher.NewCBCDecrypter(block, key[:aes.BlockSize])
	mode.CryptBlocks(ciphertext, ciphertext)

	ciphertext, err = PKCS7UnPadding(ciphertext, aes.BlockSize)
	if err != nil {
		return "", err
	}

	if outfile != "" {
		err = os.WriteFile(outfile, ciphertext, 0644)
		if err != nil {
			return "", nil
		}
	}

	return string(ciphertext), nil
}

// Create a SHA256 digest of the given password
func keyDigest(password string) []byte {
	digest := sha256.Sum256([]byte(password))
	return digest[:]
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(ciphertext []byte, blockSize int) ([]byte, error) {
	length := len(ciphertext)
	padding := int(ciphertext[length-1])

	if padding < 1 || padding > blockSize {
		return nil, fmt.Errorf("invalid padding")
	}

	return ciphertext[:(length - padding)], nil
}
