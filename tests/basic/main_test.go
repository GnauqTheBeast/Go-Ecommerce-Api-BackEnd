package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 3, Add(1, 2), "Dung roi day ae")
}
