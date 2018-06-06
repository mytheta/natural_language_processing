package main

import (
	"fmt"
	"strings"
)

//コンマとピリオドの除去
func Parse(s string) string {
	s = strings.Trim(s, ".")
	s = strings.Replace(s, ",", "", -1)
	return s
}

//"Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
//という文を単語に分解し，各単語の（アルファベットの）文字数を先頭から出現順に並べたリストを作成せよ．
func main() {
	s := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."

	s = Parse(s)

	//func Fields(s string) []string
	//s文字列の空白文字を除去し,空白に従って分割されたsliceを返す.
	words := strings.Fields(s)

	//文字と文字数の出力
	fmt.Println(words, len(words))

	//分割された文字列の文字数を出力
	for i := 0; i < len(words); i++ {
		fmt.Printf("%d\n", len(words[i]))
	}

}
