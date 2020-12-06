package text_process

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

// CharClean the base struct
type CharClean struct {
	CharProject map[rune]rune
	CharKeep    map[rune]bool
}

// NewCharClean for new a CharClean object
func NewCharClean(charMapFiles []string, charKeepFiles []string) *CharClean {
	cc := CharClean{make(map[rune]rune), make(map[rune]bool)}
	var wg sync.WaitGroup
	wg.Add(2)
	go cc.loadCharProject(charMapFiles, &wg)
	go cc.loadCharKeep(charKeepFiles, &wg)
	wg.Wait()
	return &cc
}

func (cc *CharClean) loadCharProject(filepaths []string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, filepath := range filepaths {
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
				cc.CharProject[[]rune(ab[0])[0]] = []rune(ab[1])[0]
			}

		}
	}

	log.Println("Init char project successfully.")

}

func (cc *CharClean) loadCharKeep(filepaths []string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, filepath := range filepaths {
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
			_line := []rune(string(line))
			cc.CharKeep[_line[0]] = true
		}
	}
	log.Println("Init char keep set successfully.")
}

//Normalize the rune Slice
func (cc *CharClean) Normalize(charSlice []rune) []rune {
	normalSlice := make([]rune, 0)
	for _, char := range charSlice {
		trans, ok := cc.CharProject[char]
		if ok {
			normalSlice = append(normalSlice, trans)
		} else {
			normalSlice = append(normalSlice, char)
		}
	}
	return normalSlice
}

// Clean the char Slice return the join string
func (cc *CharClean) Clean(charSlice []rune, maxCRC int) (string, []int) {
	s := ""
	sIndex := make([]int, 0)
	topCharCount := 0
	var topChar rune
	for i, char := range charSlice {
		isKeep, ok := cc.CharKeep[char]
		if ok && isKeep {
			if char != topChar {
				topCharCount = 1
				topChar = char
			} else {
				topCharCount++
			}
			if topCharCount <= maxCRC {
				s += string(char)
				sIndex = append(sIndex, i)
			}
		}
	}

	return s, sIndex
}
