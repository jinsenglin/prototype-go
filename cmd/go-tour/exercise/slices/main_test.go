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

package main

import "testing"

func TestPic(t *testing.T) {
	if result := Pic(2, 2); len(result) == 2 {
		for idx, ele := range result {
			if len(ele) != 2 {
				t.Errorf("Input 2 | Expected Output 2 | Returned Output %v | Index %v", len(ele), idx)
			}
		}
	} else {
		t.Errorf("Input 2 | Expected Output 2 | Returned Output %v", len(result))
	}
}
