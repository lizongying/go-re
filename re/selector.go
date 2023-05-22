package re

import (
	"errors"
	"github.com/lizongying/go-re/pkg/utils"
	"regexp"
)

type Selector struct {
	str string
}

func (s *Selector) One(str string) string {
	return regexp.MustCompile(str).FindString(s.str)
}
func (s *Selector) OneInt(str string) (out int, err error) {
	r := regexp.MustCompile(str).FindString(s.str)
	out, err = utils.Str2Int(r)
	return
}
func (s *Selector) OneUint(str string) (out uint, err error) {
	r := regexp.MustCompile(str).FindString(s.str)
	out, err = utils.Str2Uint(r)
	return
}
func (s *Selector) OneInt8(str string) (out int8, err error) {
	r := regexp.MustCompile(str).FindString(s.str)
	out, err = utils.Str2Int8(r)
	return
}
func (s *Selector) OneUint8(str string) (out uint8, err error) {
	r := regexp.MustCompile(str).FindString(s.str)
	out, err = utils.Str2Uint8(r)
	return
}
func (s *Selector) OneInt16(str string) (out int16, err error) {
	r := regexp.MustCompile(str).FindString(s.str)
	out, err = utils.Str2Int16(r)
	return
}
func (s *Selector) OneUint16(str string) (out uint16, err error) {
	r := regexp.MustCompile(str).FindString(s.str)
	out, err = utils.Str2Uint16(r)
	return
}
func (s *Selector) OneInt32(str string) (out int32, err error) {
	r := regexp.MustCompile(str).FindString(s.str)
	out, err = utils.Str2Int32(r)
	return
}
func (s *Selector) OneUint32(str string) (out uint32, err error) {
	r := regexp.MustCompile(str).FindString(s.str)
	out, err = utils.Str2Uint32(r)
	return
}
func (s *Selector) OneInt64(str string) (out int64, err error) {
	r := regexp.MustCompile(str).FindString(s.str)
	out, err = utils.Str2Int64(r)
	return
}
func (s *Selector) OneUint64(str string) (out uint64, err error) {
	r := regexp.MustCompile(str).FindString(s.str)
	out, err = utils.Str2Uint64(r)
	return
}
func (s *Selector) OneSub(str string) (out string) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out = r[1]
	return
}
func (s *Selector) OneSubInt(str string) (out int, err error) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out, err = utils.Str2Int(r[1])
	return
}

func (s *Selector) OneSubUint(str string) (out uint, err error) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out, err = utils.Str2Uint(r[1])
	return
}

func (s *Selector) OneSubInt8(str string) (out int8, err error) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out, err = utils.Str2Int8(r[1])
	return
}

func (s *Selector) OneSubUint8(str string) (out uint8, err error) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out, err = utils.Str2Uint8(r[1])
	return
}

func (s *Selector) OneSubInt16(str string) (out int16, err error) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out, err = utils.Str2Int16(r[1])
	return
}

func (s *Selector) OneSubUint16(str string) (out uint16, err error) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out, err = utils.Str2Uint16(r[1])
	return
}

func (s *Selector) OneSubInt32(str string) (out int32, err error) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out, err = utils.Str2Int32(r[1])
	return
}

func (s *Selector) OneSubUint32(str string) (out uint32, err error) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out, err = utils.Str2Uint32(r[1])
	return
}

func (s *Selector) OneSubInt64(str string) (out int64, err error) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out, err = utils.Str2Int64(r[1])
	return
}

func (s *Selector) OneSubUint64(str string) (out uint64, err error) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out, err = utils.Str2Uint64(r[1])
	return
}
func (s *Selector) MustOneInt(str string) (out int) {
	r := regexp.MustCompile(str).FindString(s.str)
	out, _ = utils.Str2Int(r)
	return
}
func (s *Selector) MustOneUint(str string) (out uint) {
	r := regexp.MustCompile(str).FindString(s.str)
	out, _ = utils.Str2Uint(r)
	return
}
func (s *Selector) MustOneInt8(str string) (out int8) {
	r := regexp.MustCompile(str).FindString(s.str)
	out, _ = utils.Str2Int8(r)
	return
}
func (s *Selector) MustOneUint8(str string) (out uint8) {
	r := regexp.MustCompile(str).FindString(s.str)
	out, _ = utils.Str2Uint8(r)
	return
}
func (s *Selector) MustOneInt16(str string) (out int16) {
	r := regexp.MustCompile(str).FindString(s.str)
	out, _ = utils.Str2Int16(r)
	return
}
func (s *Selector) MustOneUint16(str string) (out uint16) {
	r := regexp.MustCompile(str).FindString(s.str)
	out, _ = utils.Str2Uint16(r)
	return
}
func (s *Selector) MustOneInt32(str string) (out int32) {
	r := regexp.MustCompile(str).FindString(s.str)
	out, _ = utils.Str2Int32(r)
	return
}
func (s *Selector) MustOneUint32(str string) (out uint32) {
	r := regexp.MustCompile(str).FindString(s.str)
	out, _ = utils.Str2Uint32(r)
	return
}
func (s *Selector) MustOneInt64(str string) (out int64) {
	r := regexp.MustCompile(str).FindString(s.str)
	out, _ = utils.Str2Int64(r)
	return
}
func (s *Selector) MustOneUint64(str string) (out uint64) {
	r := regexp.MustCompile(str).FindString(s.str)
	out, _ = utils.Str2Uint64(r)
	return
}
func (s *Selector) MustOneSubInt(str string) (out int) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out, _ = utils.Str2Int(r[1])
	return
}
func (s *Selector) MustOneSubUint(str string) (out uint) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out, _ = utils.Str2Uint(r[1])
	return
}
func (s *Selector) MustOneSubInt8(str string) (out int8) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out, _ = utils.Str2Int8(r[1])
	return
}
func (s *Selector) MustOneSubUint8(str string) (out uint8) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out, _ = utils.Str2Uint8(r[1])
	return
}
func (s *Selector) MustOneSubInt16(str string) (out int16) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out, _ = utils.Str2Int16(r[1])
	return
}
func (s *Selector) MustOneSubUint16(str string) (out uint16) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out, _ = utils.Str2Uint16(r[1])
	return
}
func (s *Selector) MustOneSubInt32(str string) (out int32) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out, _ = utils.Str2Int32(r[1])
	return
}
func (s *Selector) MustOneSubUint32(str string) (out uint32) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out, _ = utils.Str2Uint32(r[1])
	return
}
func (s *Selector) MustOneSubInt64(str string) (out int64) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out, _ = utils.Str2Int64(r[1])
	return
}
func (s *Selector) MustOneSubUint64(str string) (out uint64) {
	r := regexp.MustCompile(str).FindStringSubmatch(s.str)
	if len(r) < 2 {
		return
	}
	out, _ = utils.Str2Uint64(r[1])
	return
}
func (s *Selector) Many(str string) []string {
	return regexp.MustCompile(str).FindAllString(s.str, -1)
}
func (s *Selector) ManyInt(str string) (out []int, err error) {
	r := regexp.MustCompile(str).FindAllString(s.str, -1)
	for _, v := range r {
		i, e := utils.Str2Int(v)
		if e != nil {
			err = e
			return
		}
		out = append(out, i)
	}

	return
}
func (s *Selector) ManyUint(str string) (out []uint, err error) {
	r := regexp.MustCompile(str).FindAllString(s.str, -1)
	for _, v := range r {
		i, e := utils.Str2Uint(v)
		if e != nil {
			err = e
			return
		}
		out = append(out, i)
	}

	return
}
func (s *Selector) ManyInt8(str string) (out []int8, err error) {
	r := regexp.MustCompile(str).FindAllString(s.str, -1)
	for _, v := range r {
		i, e := utils.Str2Int8(v)
		if e != nil {
			err = e
			return
		}
		out = append(out, i)
	}

	return
}
func (s *Selector) ManyUint8(str string) (out []uint8, err error) {
	r := regexp.MustCompile(str).FindAllString(s.str, -1)
	for _, v := range r {
		i, e := utils.Str2Uint8(v)
		if e != nil {
			err = e
			return
		}
		out = append(out, i)
	}

	return
}
func (s *Selector) ManyInt16(str string) (out []int16, err error) {
	r := regexp.MustCompile(str).FindAllString(s.str, -1)
	for _, v := range r {
		i, e := utils.Str2Int16(v)
		if e != nil {
			err = e
			return
		}
		out = append(out, i)
	}

	return
}
func (s *Selector) ManyUint16(str string) (out []uint16, err error) {
	r := regexp.MustCompile(str).FindAllString(s.str, -1)
	for _, v := range r {
		i, e := utils.Str2Uint16(v)
		if e != nil {
			err = e
			return
		}
		out = append(out, i)
	}

	return
}
func (s *Selector) ManyInt32(str string) (out []int32, err error) {
	r := regexp.MustCompile(str).FindAllString(s.str, -1)
	for _, v := range r {
		i, e := utils.Str2Int32(v)
		if e != nil {
			err = e
			return
		}
		out = append(out, i)
	}

	return
}
func (s *Selector) ManyUint32(str string) (out []uint32, err error) {
	r := regexp.MustCompile(str).FindAllString(s.str, -1)
	for _, v := range r {
		i, e := utils.Str2Uint32(v)
		if e != nil {
			err = e
			return
		}
		out = append(out, i)
	}

	return
}
func (s *Selector) ManyInt64(str string) (out []int64, err error) {
	r := regexp.MustCompile(str).FindAllString(s.str, -1)
	for _, v := range r {
		i, e := utils.Str2Int64(v)
		if e != nil {
			err = e
			return
		}
		out = append(out, i)
	}

	return
}
func (s *Selector) ManyUint64(str string) (out []uint64, err error) {
	r := regexp.MustCompile(str).FindAllString(s.str, -1)
	for _, v := range r {
		i, e := utils.Str2Uint64(v)
		if e != nil {
			err = e
			return
		}
		out = append(out, i)
	}

	return
}
func (s *Selector) MustManyInt(str string) (out []int) {
	r := regexp.MustCompile(str).FindAllString(s.str, -1)
	for _, v := range r {
		i, e := utils.Str2Int(v)
		if e != nil {
			continue
		}
		out = append(out, i)
	}

	return
}
func (s *Selector) MustManyUint(str string) (out []uint) {
	r := regexp.MustCompile(str).FindAllString(s.str, -1)
	for _, v := range r {
		i, e := utils.Str2Uint(v)
		if e != nil {
			continue
		}
		out = append(out, i)
	}

	return
}
func (s *Selector) MustManyInt8(str string) (out []int8) {
	r := regexp.MustCompile(str).FindAllString(s.str, -1)
	for _, v := range r {
		i, e := utils.Str2Int8(v)
		if e != nil {
			continue
		}
		out = append(out, i)
	}

	return
}
func (s *Selector) MustManyUint8(str string) (out []uint8) {
	r := regexp.MustCompile(str).FindAllString(s.str, -1)
	for _, v := range r {
		i, e := utils.Str2Uint8(v)
		if e != nil {
			continue
		}
		out = append(out, i)
	}

	return
}
func (s *Selector) MustManyInt16(str string) (out []int16) {
	r := regexp.MustCompile(str).FindAllString(s.str, -1)
	for _, v := range r {
		i, e := utils.Str2Int16(v)
		if e != nil {
			continue
		}
		out = append(out, i)
	}

	return
}
func (s *Selector) MustManyUint16(str string) (out []uint16) {
	r := regexp.MustCompile(str).FindAllString(s.str, -1)
	for _, v := range r {
		i, e := utils.Str2Uint16(v)
		if e != nil {
			continue
		}
		out = append(out, i)
	}

	return
}
func (s *Selector) MustManyInt32(str string) (out []int32) {
	r := regexp.MustCompile(str).FindAllString(s.str, -1)
	for _, v := range r {
		i, e := utils.Str2Int32(v)
		if e != nil {
			continue
		}
		out = append(out, i)
	}

	return
}
func (s *Selector) MustManyUint32(str string) (out []uint32) {
	r := regexp.MustCompile(str).FindAllString(s.str, -1)
	for _, v := range r {
		i, e := utils.Str2Uint32(v)
		if e != nil {
			continue
		}
		out = append(out, i)
	}

	return
}
func (s *Selector) MustManyInt64(str string) (out []int64) {
	r := regexp.MustCompile(str).FindAllString(s.str, -1)
	for _, v := range r {
		i, e := utils.Str2Int64(v)
		if e != nil {
			continue
		}
		out = append(out, i)
	}

	return
}
func (s *Selector) MustManyUint64(str string) (out []uint64) {
	r := regexp.MustCompile(str).FindAllString(s.str, -1)
	for _, v := range r {
		i, e := utils.Str2Uint64(v)
		if e != nil {
			continue
		}
		out = append(out, i)
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
