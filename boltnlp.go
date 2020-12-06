package main

import (
	"fmt"
	"github.com/mikuh/boltnlp/text_process"
)

func main() {
	fmt.Println("hello")

	cc := text_process.NewCharClean([]string{"./data/char_project.txt",}, []string{"./data/char_keep.txt", })
	fmt.Println(cc.Clean([]rune("我草你妈%……&bbbbb*aaaaa"), 3))
}
