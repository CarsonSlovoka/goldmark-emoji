package emoji

import "github.com/yuin/goldmark/renderer"

type emojiHTMLRenderer struct{}

func NewEmojiHTMLRenderer() renderer.NodeRenderer {
	return &emojiHTMLRenderer{}
}

func (e emojiHTMLRenderer) RegisterFuncs(renderer.NodeRendererFuncRegisterer) {

}
