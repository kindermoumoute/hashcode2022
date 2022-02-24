package models

import (
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

type Input struct {
}

func ParseInput(s string) *Input {
	lines := strings.Split(s, "\n")
	_ = pkg.ParseIntList(lines[0], " ")
	return &Input{}
}
