package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var gVote, bVote int =0, 0
	var n int
	var gMsg string = ""
	var bMsg string = ""

	fmt.Scanf("%d", &n)

	for i := 0; i < n; i++ {
		var msg string
		fmt.Scanf("%s", &msg)
		if strings.Contains(msg, "G:") {
			gMsg = gMsg+msg
		}else {
			bMsg = bMsg+msg
		}
	}
	fmt.Println(gMsg)
	fmt.Println(bMsg)
	gMsgArr := strings.Split(gMsg, " ")
	for _, s := range gMsgArr{
		data, err := strconv.Atoi(s)
		if err == nil {
			if data == 19 || data == 20  {
				gVote = gVote + 2
			}else {
				bVote = bVote + 2
			}
		}
	}
	bMsgArr := strings.Split(bMsg, " ")
	for _, s := range bMsgArr{
		data, err := strconv.Atoi(s)
		if err == nil {
			if data == 19 || data == 20  {
				gVote = gVote + 1
			}else {
				bVote = bVote + 1
			}
		}
	}

	if gVote > bVote {
		fmt.Println("Date")
	}else {
		fmt.Println("No Date")
	}

}
