package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//１行読み取り
func scan(word string) string {
	var scan = bufio.NewScanner(os.Stdin)
	fmt.Printf("%v : ", word)

	scan.Scan()

	return scan.Text()
}

//英小文字ならば(219 - 文字コード)の文字に置換、その他の文字はそのまま出力
//英小文字 [a-z] はutf-8の10進数で 97 ~ 123 で表現される
func Cipher(s string) string {
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {

		//置換
		if runes[i] > 96 && runes[i] < 123 {
			runes[i] = 219 - runes[i]
		}
	}
	return string(runes)
}

// 与えられた文字列の各文字を，以下の仕様で変換する関数cipherを実装せよ
// 英小文字ならば(219 - 文字コード)の文字に置換
// その他の文字はそのまま出力
// この関数を用い，英語のメッセージを暗号化・復号化せよ.
func main() {
	sent := strings.Trim(scan("input sentence"), "\n")
	cipherSent := Cipher(sent)
	fmt.Printf("cipherted sentence : %v\n", cipherSent)
	decodeSent := Cipher(cipherSent)
	fmt.Printf("decoded sentence : %v\n", decodeSent)
}
