package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	retVal := Hello("don")
	assert.Equal(t, "Hello don", retVal)
}
