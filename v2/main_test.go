package emoji_test

import (
	emoji "github.com/CarsonSlovoka/goldmark-emoji/v2"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"
	"testing"
)

func TestNewEmojiExtender(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			emoji.NewEmojiExtender(),
		),
	)
	testutil.DoTestCaseFile(markdown, "test/emoji.txt", t)
}
