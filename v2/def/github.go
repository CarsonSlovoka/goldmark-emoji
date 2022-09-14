package def

// Github https://emojipedia.org/man-red-hair/
func Github() Emojis {
	return NewEmojis(
		NewEmoji("man: read hair", []rune{0x1F468, 0x200D, 0x1F9B0}, "red_haired_man", "man_red_hair"),
		NewEmoji("Smile", []rune{0x1F60A}, "blush", "smiling_face_with_smiling_eyes"),
	)
}
