package main

import "fmt"

//Link は文字列と取り出すn文字目を入力として受け取り，連結して返す関数.
func Link(s string, n [4]int) string {
	var result string
	runes := []rune(s)

	for i := 0; i < len(n); i++ {
		result = result + string(runes[n[i]])
	}

	return result
}

//「パタトクカシーー」という文字列の1,3,5,7文字目を取り出して連結した文字列を得よ．
func main() {

	s := "パタトクカシーー"
	nums := [4]int{1, 3, 5, 7}

	fmt.Printf("%v\n", s)
	fmt.Println(Link(s, nums))

}
