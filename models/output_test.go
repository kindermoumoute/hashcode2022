package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	o := &Output{}
	assert.Equal(t, `3
2 cheese peppers
0
1 basil
1 pineapple
2 mushrooms tomatoes
1 basil`, o.Generate())
}
