package re

import (
	"errors"
	"regexp"
)

type Selector struct {
	str string
}

func (s *Selector) One(path string) (result *Result) {
	r := regexp.MustCompile(path).FindStringSubmatch(s.str)
	if len(r) < 2 {
		result = NewResult("")
		return
	}
	result = NewResult(r[1])
	return
}

func (s *Selector) Many(path string) (results []*Result) {
	r := regexp.MustCompile(path).FindAllStringSubmatch(s.str, -1)
	for _, i := range r {
		for _, ii := range i[1:] {
			results = append(results, NewResult(ii))
		}
	}
	return
}

func (s *Selector) OneSelector(path string) (selector *Selector) {
	r := regexp.MustCompile(path).FindStringSubmatch(s.str)
	if len(r) < 2 {
		s.str = ""
		return s
	}
	s.str = r[1]
	return s
}

func (s *Selector) ManySelector(path string) (selectors []*Selector) {
	r := regexp.MustCompile(path).FindAllStringSubmatch(s.str, -1)
	if len(r) < 2 {
		return
	}

	for _, i := range r {
		for _, ii := range i[1:] {
			selectors = append(selectors, &Selector{
				str: ii,
			})
		}
	}
	return
}

// NewSelectorFromBytes selector from bytes
func NewSelectorFromBytes(bs []byte) (selector *Selector, err error) {
	if len(bs) == 0 {
		err = errors.New("bs is empty")
		return
	}
	selector = &Selector{
		str: string(bs),
	}
	return
}

// NewSelectorFromStr selector from str
func NewSelectorFromStr(str string) (selector *Selector, err error) {
	if str == "" {
		err = errors.New("str is empty")
		return
	}
	selector = &Selector{
		str: str,
	}
	return
}
