package main

import (
	"fmt"
	"strings"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i :=0; i<t; i++  {
		var str string
		var count int = 0
		fmt.Scanf("%s", &str)
		str = strings.ToUpper(str)
		vowelMap := map[string]int{
			"A" : 1,
			"E" : 1,
			"I" : 1,
			"O" : 1,
			"U" : 1,
		}
		for s, _ := range str {
			if  vowelMap[string(str[s])] == 1 {
				count++
			}
		}
		fmt.Println(count)
	}
}
