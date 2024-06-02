package entities

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"sort"
	"strings"
)

func GenMD5(food string, salt ...string) string {
	m := hmac.New(md5.New, mixSalt(salt))
	m.Write([]byte(food))
	return hex.EncodeToString(m.Sum(nil))
}

func GenSHA1(food string, salt ...string) string {
	m := hmac.New(sha1.New, mixSalt(salt))
	m.Write([]byte(food))
	return hex.EncodeToString(m.Sum(nil))
}

func GenSHA256(food string, salt ...string) string {
	m := hmac.New(sha256.New, mixSalt(salt))
	m.Write([]byte(food))
	return hex.EncodeToString(m.Sum(nil))
}

func GenSHA384(food string, salt ...string) string {
	m := hmac.New(sha512.New384, mixSalt(salt))
	m.Write([]byte(food))
	return hex.EncodeToString(m.Sum(nil))
}

func GenSHA512(food string, salt ...string) string {
	m := hmac.New(sha512.New, mixSalt(salt))
	m.Write([]byte(food))
	return hex.EncodeToString(m.Sum(nil))
}

const delimiter = "#"

// recommended to mix secret salt
func mixSalt(salt []string) []byte {
	sort.Strings(salt)
	return []byte(strings.Join(salt, delimiter))
}
