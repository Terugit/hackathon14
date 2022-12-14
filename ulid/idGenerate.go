package ulid

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

func IdGenerate() (string, error) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id, err := ulid.New(ulid.Timestamp(t), entropy)
	return id.String(), err
}
