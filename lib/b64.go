package lib

import (
	"encoding/base64"
)

func Encode(data string, urlsafe bool) string {
	dataBytes := []byte(data)
	if urlsafe {
		return base64.URLEncoding.EncodeToString(dataBytes)
	}
	return base64.StdEncoding.EncodeToString(dataBytes)
}

func Decode(data string, urlsafe bool) (string, error) {
	var decodedByte []byte
	var err error
	if urlsafe {
		decodedByte, err = base64.URLEncoding.DecodeString(data)

	} else {
		decodedByte, err = base64.StdEncoding.DecodeString(data)
	}

	if err != nil {
		return "", err
	}

	return string(decodedByte), nil
}
