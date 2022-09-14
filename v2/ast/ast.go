package ast

import (
	"fmt"
	"github.com/CarsonSlovoka/goldmark-emoji/v2/def"
	"github.com/yuin/goldmark/ast"
)

type NodeEmoji struct {
	ast.BaseInline // 嵌入預設結構，如此才不用實現所有ast.Node所要求的所有方法(BaseInline已經有提供大部分的方法)

	// ↓ 以下成員內容可依需求自定義
	Alias string
	Value *def.Emoji
}

func NewNodeEmoji(alias string, value *def.Emoji) *NodeEmoji {
	return &NodeEmoji{Alias: alias, Value: value}
}

func (e *NodeEmoji) Dump(source []byte, level int) {
	// 可以寫入您有興趣想知道的訊息
	m := map[string]string{
		"Kind": fmt.Sprintf("%d", e.Kind()),
	}
	ast.DumpHelper(e, source, level, m, nil)
}

var nodeKindEmoji = ast.NewNodeKind("Emoji")

func (e *NodeEmoji) Kind() ast.NodeKind {
	return nodeKindEmoji
}
