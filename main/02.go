package main

import "fmt"

//Alternately は二つの文字列wo交互に連結する関数.
func Alternately(s1 string, s2 string) string {
	var result string
	runes1 := []rune(s1)
	runes2 := []rune(s2)

	for i := 0; i < len(runes1); i++ {
		result = result + string(runes1[i]) + string(runes2[i])
	}

	return result
}

//「パトカー」＋「タクシー」の文字を先頭から交互に連結して文字列「パタトクカシーー」を得よ
func main() {
	patrolCar := "パトカー"
	taxi := "タクシー"

	fmt.Println(Alternately(patrolCar, taxi))

}
