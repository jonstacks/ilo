package ilo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	random := randomString(10)
	msg := "Generates a random string of specified length"
	assert.Equal(t, 10, len(random), msg)
}
