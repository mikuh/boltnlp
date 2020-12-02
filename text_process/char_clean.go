package text_process

import (
	"bufio"
	"io"
	"log"
	"os"

	// "regexp"
	"strings"
)

// CharClean the base struct
type CharClean struct {
	charProject map[string]string
}

// func init() {
// 	cc := CharClean{make(map[string]string)}
// }

func (cc CharClean) loadCharProject(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		line, _, c := r.ReadLine()
		if c == io.EOF {
			break
		}
		ab := strings.Split(string(line), "\t")
		if len(ab) == 2 {
			cc.charProject[ab[0]] = ab[1]
		}

	}
	log.Println("Init char project successfully.")

}

func (cc CharClean) normalize(sentence string) string {
	chs := strings.Split(sentence, "")
	s := ""
	for _, ch := range chs {
		transCh, ok := cc.charProject[ch]
		if ok {
			s += transCh
		} else {
			s += ch
		}
	}
	return s
}
