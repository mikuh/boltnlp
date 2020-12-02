package boltnlp

import (
	"fmt"
	"github.com/mikuh/boltnlp/text_process"
)


func main(){
	fmt.Println("hello")
	cc := text_process.CharClean{make(map[string]string)}
	cc.loadCharProject("data/char_project.txt")
	fmt.Println(cc.normalize("⾒dao了"))
}