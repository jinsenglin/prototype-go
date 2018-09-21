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
