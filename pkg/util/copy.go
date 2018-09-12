package util

import (
	"encoding/json"

	"github.com/jinsenglin/prototype-go/pkg/model"
)

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

	switch src.(type) {
	case model.User:
		var dst model.User
		json.Unmarshal(b, &dst)
		return dst
	default:
		var dst map[string]interface{}
		json.Unmarshal(b, &dst)
		return dst
	}
}
