package def

func TW() Emojis {
	return NewEmojis(
		NewEmoji("๐ ่็ด็ฌ่", []rune{0x1F60A}, "blush", "smiling_face_with_smiling_eyes", "่็ด", "่็ด็ฌ"), // ๐
		NewEmoji("๐ Slightly Smiling Face", []rune{0x1F642}, "blush", "smiling_face_with_smiling_eyes", "ๅพฎ็ฌ"),
		NewEmoji("๐ Grinning Face with Smiling Eyes", []rune{0x1F604}, "้ๅฟ็ฌ"),
		NewEmoji("๐ Beaming Face with Smiling Eyes", []rune{0x1F601}, "้ ็ฎ็ฌ", "ๆกไฝๅ็ฌ", "ๆกไฝๅ", "ๆๅผ"),
		NewEmoji("๐ smiling face with sunglasses", []rune{0x1f60e}, "็ฅๆฐฃ", "่ทฉ", "ๅพๆ"),

		NewEmoji("\U0001F97A Pleading Face", []rune{0x1F97A}, "ๆท็ผๆฑชๆฑช", "็ฅๆฑ่", "็ฅๆฑ"), // ๐ฅบ

		NewEmoji("๐ Folded Hands", []rune{0x1F64F}, "ๆ่จ", "่ซๆฑ", "ๆๆ"),
	)
}
