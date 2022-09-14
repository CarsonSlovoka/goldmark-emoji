package def

// Github https://emojipedia.org/man-red-hair/
func Github() Emojis {
	return NewEmojis(
		NewEmoji("man: red hair", []rune{0x1F468, 0x200D, 0x1F9B0}, "red_haired_man", "man_red_hair"),
		NewEmoji("😊", []rune{0x1F60A}, "blush", "smiling_face_with_smiling_eyes", "微笑"),
	)
}
