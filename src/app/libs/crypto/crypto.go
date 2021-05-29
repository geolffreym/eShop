package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func ComputeHmac256(data string, secret string) string {
	key := []byte(secret)
	message := []byte(data)

	hash := hmac.New(sha256.New, key)
	hash.Write(message)

	//hex.EncodeToString(hash.Sum(nil))

	return base64.StdEncoding.EncodeToString(
		hash.Sum(nil))
}
