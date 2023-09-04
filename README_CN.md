# go-re

golang实现的re选择器，使用更简单。

[go-re](https://github.com/lizongying/go-re)

[document](https://pkg.go.dev/github.com/lizongying/go-re)

[english](./README.md)

## 安装

```
go get github.com/lizongying/go-re
```

## 用法

```go
package main

import (
	"fmt"
	"github.com/lizongying/go-re/re"
)

func main() {
	html := `<html class="abc">....<div class="def">....</div><div class="gkl">123</div></html>`
	s, _ := re.NewSelectorFromStr(html)

	i, e := s.OneInt8(`\d+`)
	//123
	fmt.Println(i, e)

	u64, e := s.OneUint64(`\d+`)
	//123
	fmt.Println(u64, e)

	m := s.OneSub(`class="([^"]+)`)
	//abc
	fmt.Println(m, e)
}

```