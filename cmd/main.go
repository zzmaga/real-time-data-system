package main

import "fmt"

func main() {
	i := 0
	s := 0
	for i+1 < 10 {
		i = i + 2
		s = s + i
		if s > 16 {
			break
		}
	}
	i = i + 1
	fmt.Println(i)
}
