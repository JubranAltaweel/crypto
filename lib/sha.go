package lib

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func Hash(data string, raw bool) string {
	sum := sha256.Sum256([]byte(data))

	if raw {
		return hex.EncodeToString(sum[:])
	}

	return base64.StdEncoding.EncodeToString(sum[:])
}

func Checksum(filepath string, raw bool) (string, error) {
	data, err := os.ReadFile(filepath)

	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("Error: File not found at %s", filepath)
		}

		return "", err
	}

	sum := sha256.Sum256(data)
	if raw {
		return hex.EncodeToString(sum[:]), nil
	}
	return base64.StdEncoding.EncodeToString(sum[:]), nil
}
