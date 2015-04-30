package parser

import (
	"bytes"
	"github.com/spf13/hugo/parser"
)

func Fuzz(data []byte) int {
	score := 0
	p, err := parser.ReadFrom(bytes.NewReader(data))

	if err != nil {
		return 0
	}

	if p.FrontMatter() != nil {
		score++
	}

	m, err := p.Metadata()

	if err != nil && m != nil {
		score++
	}

	if p.IsRenderable() {
		score++
	}

	return score
}
