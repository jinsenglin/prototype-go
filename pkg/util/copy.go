//    Copyright 2018 cclin
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// Package util ...
package util

import (
	"encoding/json"

	"github.com/jinsenglin/prototype-go/pkg/model"
)

// CopyIntSlice ...
func CopyIntSlice(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

// CopyStringIntMap ...
func CopyStringIntMap(src map[string]int) map[string]int {
	dst := make(map[string]int)
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

// CopyStruct ...
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
