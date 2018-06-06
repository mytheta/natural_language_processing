package main

import (
	"fmt"
	"strings"
)

//"Hi He Lied Because Boron Could Not Oxidize Fluorine.
//New Nations Might Also Sign Peace Security Clause. Arthur King Can." という文を単語に分解し，
//1, 5, 6, 7, 8, 9, 15, 16, 19番目の単語は先頭の1文字，それ以外の単語は先頭に2文字を取り出し，
//取り出した文字列から単語の位置（先頭から何番目の単語か）への連想配列（辞書型もしくはマップ型）を作成せよ．
func main() {
	s := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."

	//.の削除
	//s = strings.Trim(s, ".")
	s = strings.Replace(s, ".", "", -1)

	words := strings.Fields(s)

	wordsIndex := make(map[string]int)

	for i := 0; i < len(words); i++ {
		runes := []rune(words[i])

		//文字が何番目か？
		num := i + 1
		switch num {
		case 1, 5, 6, 7, 8, 9, 15, 16, 19: //1, 5, 6, 7, 8, 9, 15, 16, 19番目の単語は先頭の1文字
			wordsIndex[string(runes[0])] = num
		default: //上記以外の単語は先頭の2文字
			wordsIndex[string(runes[0])+string(runes[1])] = num
		}
	}

	fmt.Println(wordsIndex)
}
