package emoji_test

import (
	"bytes"
	"fmt"
	emoji "github.com/CarsonSlovoka/goldmark-emoji/v2"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"
	"os"
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

func ExampleNewEmojiExtender() {
	markdown := goldmark.New(
		goldmark.WithExtensions(
			emoji.NewEmojiExtender(),
		),
	)

	if err := markdown.Convert([]byte(":smiling_face_with_smiling_eyes:"), os.Stdout); err != nil {
		panic(err)
	}

	out := bytes.NewBuffer(make([]byte, 0))
	content := []byte(`Hello World. :smiling_face_with_smiling_eyes: :微笑:`)
	if err := markdown.Convert(content, out); err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output:
	// <p>😊</p>
	// <p>Hello World. 😊 😊</p>
}
