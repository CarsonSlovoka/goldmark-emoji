package def

type Emoji struct {
	Name    string
	Unicode []rune   // 該emoji可能是有很多個codepoints組合而成，例如: red_haired_man: {👨 U+1F468, ‍ U+200D, 🦰 U+1F9B0} 是由這三者組合而成: https://apps.timwhitlock.info/unicode/inspect?s=%F0%9F%91%A8%E2%80%8D%F0%9F%A6%B0
	Aliases []string // 允許有很多別名
}

func NewEmoji(name string, unicode []rune, alias ...string) Emoji {
	if len(alias) == 0 {
		panic("Emoji must have at least one alias.")
	}
	return Emoji{name, unicode, alias}
}

type Emojis interface {
	Get(alias string) (*Emoji, bool)
	Add(*Emojis)
}

type emojis struct {
	list     []Emoji           // 紀錄初始狀態所擁有的Emoji
	dict     map[string]*Emoji // 以alias來對應Emoji
	children []Emojis          // Add方法不會更改原始的list而是把新的emoji「集」新增在children
}

func NewEmojis(es ...Emoji) Emojis {
	ems := &emojis{
		es,
		make(map[string]*Emoji, 0),
		make([]Emojis, 0),
	}

	// 建立dict資料
	for _, e := range ems.list {
		for _, alias := range e.Aliases {
			ems.dict[alias] = &e
		}
	}
	return ems
}

func (e *emojis) Add(emoji *Emojis) {
	e.children = append(e.children, *emoji)
}

func (e *emojis) Get(alias string) (*Emoji, bool) {
	// 從dict中找尋
	v, exists := e.dict[alias]
	if exists {
		return v, exists
	}

	// 從children找尋
	for _, es := range e.children {
		v, exists = es.Get(alias)
		if exists {
			return v, exists
		}
	}
	return nil, false
}
