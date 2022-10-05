package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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
	phase4()
	fmt.Println(enigme)
	random_number()
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
	for index, mot := range liste_mot {
		if mot == "before" {
			coord, _ := strconv.Atoi(liste_mot[index+1])
			enigme = append(enigme, liste_mot[coord-1])
		}
	}
	return enigme
}

func phase4() []string {
	for index, mot := range liste_mot {
		if strings.Contains(mot, "now") {
			mot_temp := liste_mot[index-1]
			coords := int(mot_temp[1])/len(liste_mot) - 1
			enigme = append(enigme, liste_mot[coords])
		}
	}
	return enigme
}

func random_number() {
	rand.Seed(time.Now().UTC().UnixNano())
	RandomInteger := rand.Intn(1100000000000)
	fmt.Println(RandomInteger)
}
