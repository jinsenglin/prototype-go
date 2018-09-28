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

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	if result := Sqrt(1); result != 1 {
		t.Errorf("Input 1 | Expected Output 1 | Returned Output %v", result)
	}

	if result := Sqrt(4); result != 2 {
		t.Errorf("Input 4 | Expected Output 2 | Returned Output %v", result)
	}
}

func TestSqrt9(t *testing.T) {
	if result := Sqrt(9); result != 3 {
		t.Fatalf("Input 9 | Expected Output 3 | Returned Output %v", result)
	}
}
