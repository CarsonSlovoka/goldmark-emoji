package emoji

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type emojiExtender struct{}

func NewEmojiExtender() goldmark.Extender {
	return &emojiExtender{}
}

// Extend adds a hashtag parser to a Goldmark parser
func (e *emojiExtender) Extend(m goldmark.Markdown) {

	m.Parser().AddOptions(
		// https://github.com/yuin/goldmark/blob/20df1691ad91648cd872e6472301d7f6d325b448/parser/parser.go#L566-L578
		// parser.WithBlockParsers(),

		// https://github.com/yuin/goldmark/blob/20df1691ad91648cd872e6472301d7f6d325b448/parser/parser.go#L594-L610
		parser.WithInlineParsers(
			util.Prioritized(
				NewEmojiParser(), 1000,
			),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(NewHTMLRenderer(), 200),
		),
	)
}

func NewEmojiParser() parser.InlineParser {
	return nil
}

func NewHTMLRenderer() renderer.NodeRenderer {
	return nil
}
