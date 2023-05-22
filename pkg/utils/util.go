package utils

import (
	"fmt"
	"strconv"
)

// Str2Int Convert string to int
func Str2Int(str string) (i int, err error) {
	i, err = strconv.Atoi(str)
	return
}

// Str2Uint Convert string to uint
func Str2Uint(str string) (i uint, err error) {
	i64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return
	}
	i = uint(i64)
	return
}

// Str2Int8  Convert string to int8
func Str2Int8(str string) (i int8, err error) {
	i64, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return
	}
	i = int8(i64)
	return
}

// Str2Uint8 Convert string to uint8
func Str2Uint8(str string) (i uint8, err error) {
	u64, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		return
	}
	i = uint8(u64)
	return
}

// Str2Int16  Convert string to int16
func Str2Int16(str string) (i int16, err error) {
	i64, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return
	}
	i = int16(i64)
	return
}

// Str2Uint16 Convert string to uint16
func Str2Uint16(str string) (i uint16, err error) {
	u64, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return
	}
	i = uint16(u64)
	return
}

// Str2Int32  Convert string to int32
func Str2Int32(str string) (i int32, err error) {
	i64, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return
	}
	i = int32(i64)
	return
}

// Str2Uint32 Convert string to uint32
func Str2Uint32(str string) (i uint32, err error) {
	u64, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return
	}
	i = uint32(u64)
	return
}

// Str2Int64  Convert string to int64
func Str2Int64(str string) (i int64, err error) {
	i, err = strconv.ParseInt(str, 10, 64)
	return
}

// Str2Uint64 Convert string to uint64
func Str2Uint64(str string) (i uint64, err error) {
	i, err = strconv.ParseUint(str, 10, 64)
	return
}

func Int2Str[T int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64](i T) (str string) {
	str = fmt.Sprintf("%d", i)
	return
}
