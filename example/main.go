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
