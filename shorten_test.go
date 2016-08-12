package shorten

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandSeq(t *testing.T) {
	v := RandSeq(12)

	assert.Equal(t, 12, len(v))
}
