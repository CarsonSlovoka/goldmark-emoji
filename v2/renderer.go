package emoji

import (
	"github.com/CarsonSlovoka/goldmark-emoji/v2/ast"
	gast "github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

type emojiHTMLRenderer struct{}

func NewEmojiHTMLRenderer() renderer.NodeRenderer {
	return &emojiHTMLRenderer{}
}

func (r *emojiHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.NodeKindEmoji, r.nodeRendererFunc)
}

func (r *emojiHTMLRenderer) nodeRendererFunc(writer util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	return 0, nil
}
