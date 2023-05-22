package utils

import (
	"testing"
)

func TestStr2Int(t *testing.T) {
	i, e := Str2Int("20")
	t.Log(i, e)
}

func TestStr2Int8(t *testing.T) {
	i, e := Str2Int8("20a")
	t.Log(i, e)
}
