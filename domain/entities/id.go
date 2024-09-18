package entities

import (
	"crypto/rand"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/rs/xid"
)

func GenXID() string {
	return xid.New().String()
}

func GenUUID() string {
	return uuid.NewString()
}

var entropy *ulid.MonotonicEntropy

func init() {
	entropy = ulid.Monotonic(rand.Reader, 0)
}

func GenULID() string {
	return ulid.MustNew(ulid.Timestamp(Now()), entropy).String()
}
