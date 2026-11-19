// Package rawvalidate holds a native Go fuzz harness (built with go-118-fuzz-build) for
// bson.Raw.Validate: the public API an application calls to check a BSON document it received
// from an untrusted source (a file, a socket, a wire message) before reading it.
package rawvalidate

import (
	"testing"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// FuzzRawValidate validates arbitrary bytes as a BSON document; when they validate, it reads the
// document back through the Raw accessors a caller would use next.
func FuzzRawValidate(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		r := bson.Raw(data)
		if err := r.Validate(); err != nil {
			return
		}
		elems, err := r.Elements()
		if err != nil {
			return
		}
		for _, e := range elems {
			_ = e.Key()
			_ = e.Value().String()
		}
		_ = r.String()
	})
}
