package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"math"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	// "fmt"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "01/02/06"
)

func RoundToEven(target float64) float64 {
	res := math.RoundToEven(target*100) / 100
	return res
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

func EncryptAESGCM(textStr string, keyStr string, nonceStr string) string {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).

	plaintext := []byte(textStr)
	key := []byte(keyStr)
	nonce := []byte(nonceStr)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonceCipherArray := append(nonce, aesgcm.Seal(nil, nonce, plaintext, nil)...)

	return hex.EncodeToString(nonceCipherArray)
}

func DecryptAESGCM(encodedCipherStr string, keyStr string) string {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).

	nonceCipherArray, _ := hex.DecodeString(encodedCipherStr)

	ciphertext := nonceCipherArray[12:]
	nonce := nonceCipherArray[:12]
	key := []byte(keyStr)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return string(plaintext)
}

func ParseStringtoTime(date string) (time.Time, error) {
	layout := "2006-01-02T15:04:05Z07:00" // RFC3339
	d, err := time.Parse(layout, date)
	if err != nil {
		return d, err
	}
	return d, nil
}
func ParseDateSlashFormat(date string) string {
	t, _ := time.Parse(layoutISO, date)
	res := t.Format(layoutUS)
	m := res[:2]
	d := res[3:5]
	y := res[6:]
	return d + "/" + m + "/" + y

}
func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func SplitString(str string, cut string) []string {
	var resStr []string
	if strings.Contains(str, cut) {
		resStr = strings.Split(str, cut)
	} else {
		resStr = append(resStr, str)
	}

	return resStr
}
