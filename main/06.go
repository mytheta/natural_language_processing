package main

import "fmt"

func isSame(word string, words []string) bool {
	for _, element := range words {
		if element == word {
			return true
		}
	}
	return false
}

//文字バイグラム
func CharBigram(s string, n int) []string {

	//文字列を一文字ずつruneに
	runes := []rune(s)
	cwords := []string{}
	for i := 0; i < len(runes)-n+1; i++ {
		var word string
		for j := 0; j < n; j++ {
			word += string(runes[i+j])
		}
		if !isSame(word, cwords) {
			cwords = append(cwords, word)
		}
	}
	return cwords
}

func Union(x []string, y []string) []string {
	Union := []string{}
	for _, char := range x {
		Union = append(Union, string(char))
	}
	for _, char2 := range y {
		if !isSame(char2, Union) {
			Union = append(Union, string(char2))
		}
	}
	return Union
}

func Intersection(x []string, y []string) []string {
	intersection := []string{}
	for _, char := range x {
		for _, char2 := range y {
			if string(char) == string(char2) {
				intersection = append(intersection, string(char))
			}
		}
	}
	return intersection
}

//XとYの差集合
func Difference(x []string, y []string) []string {
	difference := []string{}
	for _, char := range x {
		if !isSame(char, y) {
			difference = append(difference, string(char))
		}
	}
	return difference
}

//"paraparaparadise"と"paragraph"に含まれる文字bi-gramの集合を，
//それぞれ,XとYとして求め，XとYの和集合，積集合，差集合を求めよ．
//さらに，'se'というbi-gramがXおよびYに含まれるかどうかを調べよ．
func main() {
	s1 := "paraparaparadise"
	s2 := "paragraph"

	X := CharBigram(s1, 2)
	Y := CharBigram(s2, 2)

	fmt.Print("X : ", X, "\n")
	fmt.Print("Y : ", Y, "\n")
	fmt.Print("和集合 : ", Union(X, Y), "\n")
	fmt.Print("積集合 : ", Intersection(X, Y), "\n")
	fmt.Print("差集合 : ", Difference(X, Y), "\n")
	fmt.Print("Xにseがあるか : ", isSame("se", X), "\n")
	fmt.Print("Yにseがあるか : ", isSame("se", Y), "\n")
}
