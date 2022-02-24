package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	in := ParseInput(`3
2 cheese peppers
0
1 basil
1 pineapple
2 mushrooms tomatoes
1 basil`)

	assert.Len(t, in, 0)
}
