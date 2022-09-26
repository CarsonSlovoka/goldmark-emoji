package emoji

import (
	"github.com/CarsonSlovoka/goldmark-emoji/v2/def"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type emojiExtender struct {
	emojis def.Emojis
}

var DefaultEmojis def.Emojis

func init() {
	DefaultEmojis = def.Github()
}

func NewEmojiExtender(elems ...def.Emojis) goldmark.Extender {
	if len(elems) == 0 {
		elems = []def.Emojis{DefaultEmojis}
	}
	emojis := elems[0]
	for _, es := range elems[1:] {
		emojis.Add(&es)
	}
	return &emojiExtender{emojis}
}

// Extend adds a hashtag parser to a Goldmark parser
func (e *emojiExtender) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		// https://github.com/yuin/goldmark/blob/20df1691ad91648cd872e6472301d7f6d325b448/parser/parser.go#L566-L578
		// parser.WithBlockParsers(),

		// https://github.com/yuin/goldmark/blob/20df1691ad91648cd872e6472301d7f6d325b448/parser/parser.go#L594-L610
		parser.WithInlineParsers(
			util.Prioritized( // {any, priority} 它的any指的是一個對象，一個當被Parse成功之後，可能會返回的產物(包含在ast.Node之中)
				NewEmojiParser(e.emojis), 1000,
			),
		),
	)
	m.Renderer().AddOptions(
		renderer.WithNodeRenderers(
			util.Prioritized(NewEmojiHTMLRenderer(), 200),
		),
	)
}
