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

//引数 x, y, zを受け取り「x時のyはz」という文字列を返す関数を実装せよ．
//さらに，x=12, y="気温", z=22.4として，実行結果を確認せよ．
func main() {

	//各x,y,zをターミナルから入力
	x := strings.Trim(scan("x"), "\n")
	y := strings.Trim(scan("y"), "\n")
	z := strings.Trim(scan("z"), "\n")

	fmt.Printf("「%v時の%vは%v」\n", x, y, z)
}
