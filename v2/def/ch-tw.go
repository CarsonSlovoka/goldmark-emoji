package def

func TW() Emojis {
	return NewEmojis(
		NewEmoji("😊 臉紅笑臉", []rune{0x1F60A}, "blush", "smiling_face_with_smiling_eyes", "臉紅", "臉紅笑"), // 😊
		NewEmoji("🙂 Slightly Smiling Face", []rune{0x1F642}, "blush", "smiling_face_with_smiling_eyes", "微笑"),
		NewEmoji("😄 Grinning Face with Smiling Eyes", []rune{0x1F604}, "開心笑"),
		NewEmoji("😁 Beaming Face with Smiling Eyes", []rune{0x1F601}, "頑皮笑", "惡作劇笑", "惡作劇", "捉弄"),
		NewEmoji("😎 smiling face with sunglasses", []rune{0x1f60e}, "神氣", "跩", "得意"),

		NewEmoji("\U0001F97A Pleading Face", []rune{0x1F97A}, "淚眼汪汪", "祈求臉", "祈求"), // 🥺

		NewEmoji("🙏 Folded Hands", []rune{0x1F64F}, "拜託", "請求", "拜拜"),
	)
}
