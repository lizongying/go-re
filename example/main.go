package main

import (
	"fmt"
	"github.com/lizongying/go-re/re"
)

func main() {
	html := `<html class="abc">....<div class="def">....</div><div class="gkl">123</div></html>`
	s, _ := re.NewSelectorFromStr(html)

	i8 := s.One(`(\d+)`).Int8()
	//123
	fmt.Println(i8)

	arr := s.Many(`(\d+)`)
	//123
	fmt.Println(arr)

	m := s.ManySelector(`>([^<]+)(<)`)
	n := m[4].One(`(\d+)`)
	//123
	fmt.Println(n)
}
