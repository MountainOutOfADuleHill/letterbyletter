/*
December 12th, 2016 /r/DailyProgrammer challenge
Difficulty: Easy
*/


package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	sentence1 := reader.ReadLine('\n')
	sentence2 := reader.ReadLine('\n')
	

	fmt.Println(sentence1)

	if len(sentence1) == len(sentence2) {
		for i := 0; i < len(sentence1); i++ {
			if sentence1[i] != sentence2[i] {
				sentence1 = replaceStringRune(sentence1, sentence2, i)
				fmt.Println(sentence1)
			}
		}
	}else{
		fmt.Println("String Mismatch")
	}
}

func replaceStringRune(sentence1, sentence2 string, i int) string {
	choppedS1 := []rune(sentence1)
	choppedS1[i] = rune(sentence2[i])
	return string(choppedS1)
}