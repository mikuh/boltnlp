package text_process

import "testing"

func TestCharClean(t *testing.T) {

	cc := NewCharClean([]string{"../data/char_project.txt"}, []string{"../data/char_keep.txt"})
	if cc.CharProject['𣎴'] != '不' {
		t.Fatal("The loaded data does not meet expectations")
	}
	if cc.CharProject['鶏'] != '鸡' {
		t.Fatal("The loaded data does not meet expectations")
	}

	ss := ""
	a := []rune("wo 🈚di℃@#$%^&")
	b := cc.Normalize(a)
	for _, char := range b {
		ss += string(char)
	}
	if ss != "wo 无dic@#$%^&" {
		t.Fatal("normalize not meet expectations")
	}

	c, indexs := cc.Clean(b, 3)
	for i, v := range []rune(c) {
		if b[indexs[i]] != v {
			t.Fatal("clean not meet expectations")
		}
	}
	if c != "wo无dic" {
		t.Fatal("clean target not meet expectations")
	}
	if len([]rune(c)) != len(indexs) {
		t.Fatal("clean Index meet expectations")
	}

}

func BenchmarkNormalize(b *testing.B) {
	cc := NewCharClean([]string{"../data/char_project.txt"}, []string{"../data/char_keep.txt"})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cc.Normalize([]rune("wo 🈚🈚🈚🈚🈚🈚🈚🈚di℃"))
	}
}
