package emoji_test

import (
	"fmt"
	emoji "github.com/CarsonSlovoka/goldmark-emoji/v2"
	"github.com/yuin/goldmark"
	"testing"
)

func TestNewEmojiExtender(t *testing.T) {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			emoji.NewEmojiExtender(),
		),
	)
	fmt.Println(markdown)
	// testutil.DoTestCaseFile(markdown, "test/emoji.txt", t)
}
