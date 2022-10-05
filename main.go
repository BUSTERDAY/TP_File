package main

import (
	"fmt"
	"os"
)

var enigme []string
var liste_mot []string

func liste() []string {
	var tmp string
	data, _ := os.ReadFile("Hangman.txt")

	for index, char := range data {
		if char == 13 {
			liste_mot = append(liste_mot, tmp)
			tmp = ""
		} else if char != 10 {
			tmp += string(char)
		}
		if index == len(data)-1 {
			liste_mot = append(liste_mot, tmp)
		}
	}
	return liste_mot
}

func main() {
	liste()
	fmt.Println(liste_mot)
	phase1()
	phase2()
	phase3()
	//phase4()
	fmt.Println(enigme)
}

func phase1() []string {
	enigme = append(enigme, liste_mot[0])
	return enigme
}

func phase2() []string {
	enigme = append(enigme, liste_mot[len(liste_mot)-1])
	return enigme
}

func phase3() []string {

}

func phase4() []string {

}
