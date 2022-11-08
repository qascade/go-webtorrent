package testing

import (
	"testing"

	"github.com/qascade/go_webtorrent/pkg/file"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	want := "Hello World"
	got := file.Hello()
	assert.Equal(t, want, got, "not equal")
}
