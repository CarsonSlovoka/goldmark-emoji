package emoji

import (
	"fmt"
	"github.com/CarsonSlovoka/goldmark-emoji/v2/ast"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type emojiParser struct{}

const (
	beginPrefix = ':'
	endSuffix   = ':'
)

func NewEmojiParser() parser.InlineParser {
	return &emojiParser{}
}

func (e *emojiParser) Trigger() []byte {
	return []byte{beginPrefix}
}

// Parse 如果進入到此函數，表示Trigger已經被觸發(該條件符合)
func (e *emojiParser) Parse(parent gast.Node, block text.Reader, pc parser.Context) gast.Node {
	content, _ := block.PeekLine()
	if len(content) < 1 {
		return nil
	}
	pos := 1 // pos 0 為Trigger所觸發的byte，所以跳過它
	for ; pos < len(content); pos++ {
		c := content[pos]
		if !(util.IsAlphaNumeric(c) || c == '_' || c == '-') {
			break
		}
	}

	if pos >= len(content) || content[pos] != endSuffix {
		return nil
	}

	block.Advance(pos + 1)
	alias := content[1:pos]
	fmt.Println(alias)

	return &ast.NodeEmoji{}
}
