package emoji

import (
	"github.com/CarsonSlovoka/goldmark-emoji/v2/ast"
	gast "github.com/yuin/goldmark/ast"
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

func (e *emojiParser) Parse(parent gast.Node, block text.Reader, pc parser.Context) gast.Node {
	content, _ := block.PeekLine()
	if len(content) < 1 {
		return nil
	}
	// ...
	return &ast.NodeEmoji{}
}
