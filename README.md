# goldmark-emoji

這是一個[goldmark](https://github.com/yuin/goldmark)的擴充功能

解析的對象為 `:joy:` 等類似範本

## USAGE

更多範例請參考[ExampleNewEmojiExtender](./v2/main_test.go)的代碼

```go
package main

import (
  "bytes"
  "fmt"
  emoji "github.com/CarsonSlovoka/goldmark-emoji/v2"
  "github.com/CarsonSlovoka/goldmark-emoji/v2/def"
  "github.com/yuin/goldmark"
  "os"
)

func Example() {
  markdown := goldmark.New(
    goldmark.WithExtensions(
      emoji.NewEmojiExtender(
        // def.Github(), def.TW(), ... // 可以不給，預設會用Github來當作標準, 您可以再加入喜歡的表情符號清單，如果找不到滿意的也可以自己建立
      ),
    ),
  )

  markdown.Convert([]byte(":blush:"), os.Stdout)

  // Output:
  // <p>😊</p>
}
```


### 自建立表情清單

如果要自建立表情符號清單餵入，可以參考def資料夾的{[github-emoji.go](v2/def/github-emoji.go), [ch-tw.go](v2/def/ch-tw.go)}

```go
func MyEmoji() Emojis {
  return NewEmojis(
    NewEmoji("desc ...", []rune{0x1F97A}, "QQ", "Pleading Face", ""), // 🥺
    // ...
    // NewEmoji("desc ...", []rune{}, ""),
  )
}

markdown := goldmark.New(
  goldmark.WithExtensions(
    emoji.NewEmojiExtender(
      def.Github(), MyEmoji()
    ),
  ),
)

markdown.Convert([]byte(":QQ:"), os.Stdout)
```

## 參考資料

- [github.com/yuin/goldmark-emoji](https://github.com/yuin/goldmark-emoji#goldmark-emoji)
- [emojipedia.org](https://emojipedia.org/man-red-hair/)
