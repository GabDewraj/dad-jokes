package secrets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSecret(t *testing.T) {
	assertWithTest := assert.New(t)
	assertWithTest.Nil(vault)
}
