package main

import (
	"fmt"
	"strings"
)

//単語バイグラム
func WordNgram(s string, n int) []string {
	//各単語をスライスに
	words := strings.Fields(s)
	nwords := []string{}

	for i := 0; i < len(words)-n+1; i++ {
		var word string
		for j := 0; j < n; j++ {
			word += words[i+j]
		}
		nwords = append(nwords, word)
	}
	return nwords
}

//文字バイグラム
func CharNgram(s string, n int) []string {

	//文字列を一文字ずつruneに
	runes := []rune(s)
	cwords := []string{}
	for i := 0; i < len(runes)-n+1; i++ {
		var word string
		for j := 0; j < n; j++ {
			word += string(runes[i+j])
		}
		cwords = append(cwords, word)
	}
	return cwords
}

//与えられたシーケンス（文字列やリストなど）からn-gramを作る関数を作成せよ．
//この関数を用い，"I am an NLPer"という文から単語bi-gram，文字bi-gramを得よ．
func main() {
	s := "I am an NLPer"

	wordsNgram := WordNgram(s, 2)
	charNgram := CharNgram(s, 2)
	fmt.Print("単語bigram:")
	fmt.Println(wordsNgram)
	fmt.Print("文字bigram:")
	fmt.Println(charNgram)

}
