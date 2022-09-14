package def_test

import (
	"fmt"
	"github.com/CarsonSlovoka/goldmark-emoji/v2/def"
	"testing"
)

func TestNewEmojis(t *testing.T) {
	es := def.NewEmojis(
		def.NewEmoji("man: red hair", []rune{0x1F468, 0x200D, 0x1F9B0}, "red_haired_man", "man_red_hair"),
		def.NewEmoji("😊", []rune{0x1F60A}, "blush", "smiling_face_with_smiling_eyes", "微笑"),
	)
	e1, exists := es.Get("man_red_hair")
	if !exists {
		t.Fatal("emoji not exist.")
	}
	if len(e1.Unicode) != 3 {
		t.Fatal("length not correct")
	}

	e1, _ = es.Get("blush")
	if e1.Name != "😊" {
		t.Fatal("name not correct")
	}
}

func ExampleNewEmojis() {
	es := def.NewEmojis(
		def.NewEmoji("man: red hair", []rune{0x1F468, 0x200D, 0x1F9B0}, "red_haired_man", "man_red_hair"),
		def.NewEmoji("😊", []rune{0x1F60A}, "blush", "smiling_face_with_smiling_eyes", "微笑"),
	)
	e, _ := es.Get("man_red_hair")
	fmt.Println(e.Name)
	fmt.Printf("%s\n", string(e.Unicode))
	fmt.Println(e.Aliases[0])
	// Output:
	// man: red hair
	// 👨‍🦰
	// red_haired_man
}
