package entities

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/june-style/go-sample/domain/derrors"
)

func GenToken(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", derrors.Wrap(err)
	}
	return base64.RawURLEncoding.EncodeToString(b), nil

}
