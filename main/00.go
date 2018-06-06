package main

import "fmt"

func Reverse(s string) string {

	//文字列を示すstringの中身はbyte配列のため，
	//文字列を文字単位で扱うには rune を使う．
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//文字列"stressed"の文字を逆に（末尾から先頭に向かって）並べた文字列を得よ．
func main() {
	s := "stressed"

	fmt.Printf("%v\n", s)
	fmt.Println(Reverse(s))

}
