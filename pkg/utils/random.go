package utils

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetRandom(min, max int) int {
	return rand.Intn(max-min) + min + 1 //r.Intn(max-min) + min + 1
}

func GetRandomMax(max int) int {
	return GetRandom(0, max) - 1
}

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

// 随机字符串
func RandStrCommon(size int, kind int) []byte {
	kinds, result := [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random kind
			kind = rand.Intn(3)
		}
		scope, base := kinds[kind][0], kinds[kind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}

func RandStr(size int) string {
	kind := 3
	kinds, result := [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		kind = rand.Intn(3)
		scope, base := kinds[kind][0], kinds[kind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}
func RandInt(i int64) int64 {
	if i == 0 {
		return 0
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(i)
}
func RandInt2(i int) int {
	if i == 0 {
		return 0
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(i)
}
