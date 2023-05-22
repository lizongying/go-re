# go-re

通过正则提取字符串、数字等

[go-re](https://github.com/lizongying/go-re)
[document](https://pkg.go.dev/github.com/lizongying/go-re)

## Install

```
go get github.com/lizongying/go-re
```

## Usage

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