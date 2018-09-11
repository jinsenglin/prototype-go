package util

import "encoding/json"

func CopyIntSlice(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

func CopyStringIntMap(src map[string]int) map[string]int {
	dst := make(map[string]int)
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func CopyStruct(src interface{}) interface{} {
	b, _ := json.Marshal(src)
	var dst interface{}
	json.Unmarshal(b, &dst)
	return dst
}
