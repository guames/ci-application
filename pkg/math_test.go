package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestName(t *testing.T) {
	soma := Sum(10,10)
	assert.Equal(t, soma, 20)
}
