package encode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	hash, err := HashPassword("hello")
	assert.Nil(t, err)
	assert.NotEmpty(t, hash)

}

func TestVerify(t *testing.T) {
	hash, err := HashPassword("hello")
	assert.Nil(t, err)
	assert.NotEmpty(t, hash)

	ok, err := VerifyPassword("hello", hash)
	assert.Nil(t, err)
	assert.True(t, ok)
}
