package emoji

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type emojiParser struct{}

func NewEmojiParser() parser.InlineParser {
	return &emojiParser{}
}

func (e *emojiParser) Trigger() []byte {
	return []byte{':'}
}

func (e *emojiParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	return nil
}
