# goldmark-emoji

這是一個[goldmark](https://github.com/yuin/goldmark)的擴充功能

解析的對象為 `:joy:` 等類似範本

## USAGE

```go
package main

import (
  "bytes"
  "fmt"
  emoji "github.com/CarsonSlovoka/goldmark-emoji/v2"
  "github.com/yuin/goldmark"
  "os"
)

func Example() {
  markdown := goldmark.New(
    goldmark.WithExtensions(
      emoji.NewEmojiExtender(),
    ),
  )

  markdown.Convert([]byte(":smiling_face_with_smiling_eyes:"), os.Stdout)

  // Output:
  // <p>😊</p>
}
```

更多範例請參考[ExampleNewEmojiExtender](./v2/main_test.go)的代碼


## 參考資料

- [github.com/yuin/goldmark-emoji](https://github.com/yuin/goldmark-emoji#goldmark-emoji)
- [emojipedia.org](https://emojipedia.org/man-red-hair/)
