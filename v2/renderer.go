package emoji

import (
	"fmt"
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

// nodeRendererFunc 此函數可以把結果直接寫到writer之中，就能對輸出產生影響
func (r *emojiHTMLRenderer) nodeRendererFunc(writer util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	if !entering {
		return gast.WalkContinue, nil
	}
	node := n.(*ast.NodeEmoji)
	_, _ = fmt.Fprintf(writer, "%s", string(node.Value.Unicode))
	return 0, nil
}
