package angrypurpletiger_test

import (
	"testing"

	apt "github.com/ormembaar/angry-purple-tiger"
)

func Test_Sum(t *testing.T) {
	data := []byte("112CuoXo7WCcp6GGwDNBo6H5nKXGH45UNJ39iEefdv2mwmnwdFt8")
	expected := "feisty-glass-dalmatian"

	digest := apt.Sum(data)
	if digest != expected {
		t.Errorf("bad digest: expected %s, got %s", expected, digest)
	}
}
