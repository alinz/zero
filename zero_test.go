package zero_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/alinz/zero"
)

func TestEncoderDecoder(t *testing.T) {
	var r io.Reader

	r = strings.NewReader("Hello World")

	// encode the given string reader

	var encoded bytes.Buffer

	w := zero.NewEncoder(&encoded)

	io.Copy(w, r)

	// decode the given zero size stream into string

	r = zero.NewDecoder(&encoded)

	var decoded bytes.Buffer

	io.Copy(&decoded, r)

	if decoded.String() != "Hello World" {
		t.Fatalf("expected decoded to be '%s' but got this '%s'", "Hello World", decoded.String())
	}
}
