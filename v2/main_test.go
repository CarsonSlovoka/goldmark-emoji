package emoji_test

import (
	"bytes"
	"fmt"
	emoji "github.com/CarsonSlovoka/goldmark-emoji/v2"
	"github.com/CarsonSlovoka/goldmark-emoji/v2/def"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"
	"os"
	"testing"
)

func TestNewEmojiExtender(t *testing.T) {
	mdGithub := goldmark.New(
		goldmark.WithExtensions(
			emoji.NewEmojiExtender(def.Github()),
		),
	)
	testutil.DoTestCaseFile(mdGithub, "test/github.txt", t)

	mdDefault := goldmark.New(
		goldmark.WithExtensions(
			emoji.NewEmojiExtender(), // 省略預設用Github來帶第
		),
	)
	testutil.DoTestCaseFile(mdDefault, "test/github.txt", t)

	markdown := goldmark.New(
		goldmark.WithExtensions(
			emoji.NewEmojiExtender(def.Github(), def.TW()),
		),
	)
	testutil.DoTestCaseFile(markdown, "test/all.txt", t)
}

func ExampleNewEmojiExtender() {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			emoji.NewEmojiExtender(def.Github(), def.TW()),
		),
	)

	if err := markdown.Convert([]byte(":blush:"), os.Stdout); err != nil {
		panic(err)
	}

	out := bytes.NewBuffer(make([]byte, 0))
	content := []byte(`Hello World. :blush: :臉紅:`)
	if err := markdown.Convert(content, out); err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output:
	// <p>😊</p>
	// <p>Hello World. 😊 😊</p>
}
