package text_process

import "testing"

func TestCharClean(t *testing.T) {

	cc := CharClean{make(map[string]string)}
	cc.loadCharProject("../data/char_project.txt")
	if cc.charProject["𣎴"] != "不" {
		t.Fatal("The loaded data does not meet expectations")
	}
	if cc.charProject["鶏"] != "鸡" {
		t.Fatal("The loaded data does not meet expectations")
	}

	if cc.normalize("wo 🈚di℃") != "wo 无dic" {
		t.Fatal("normalize not meet expectations")
	}
}
