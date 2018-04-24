package main

import (
	"fmt"
)

type Sort struct {
	param []int
}

func main() {
	sort := new(Sort)
	sort.param = []int{1, 32, 432, 53, 23, 43, 12, 42, 11, 24, 56, 54, 22}
	fmt.Println(sort.param)
	sort.Bubble()
	fmt.Println(sort.param)
}

func (s *Sort) Bubble() {
	length := len(s.param)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if s.param[j] > s.param[j+1] {
				temp := s.param[j]
				s.param[j] = s.param[j+1]
				s.param[j+1] = temp
			}
		}
	}
}
