package main

import "fmt"

func main() {
	sentence := []rune("hello from JD")
	reverse(sentence)
	begin, end := 0, 0
	for i := begin + 1; i < len(sentence); i++ {
		if string(sentence[i]) == " " {
			end = i
			reverse(sentence[begin:end])
			begin = i + 1
		} else if i == len(sentence)-1 {
			end = i + 1
			reverse(sentence[begin:end])
		}
	}
	//reverse(sentence)
	fmt.Println(string(sentence))
}

func reverse(sentence []rune) {
	for i := 0; i < len(sentence)/2; i++ {
		sentence[i], sentence[len(sentence)-1-i] = sentence[len(sentence)-1-i], sentence[i]
	}

}
