package main

import "fmt"

type BookInterface interface {
	String()
}
type Book struct {
	Title  string
	Author string
	Intor  string
}

func (s *Book) String() {
	fmt.Println(s.Title, s.Author, s.Intor)
}

type MyBook1 struct {
	bo      Book
	Content string
}

type MyBook2 struct {
	Book
	Content string
}

func main() {
	println("a")
	my1 := &MyBook1{
		Content: "xxx",
	}
	my1.bo = Book{
		Title:  "Go",
		Author: "songtianyi",
		Intor:  "GoGoGo",
	}

	my2 := &MyBook2{
		Content: "xxx",
	}
	my2.Book = Book{
		Title:  "Go",
		Author: "songtianyi",
		Intor:  "GoGoGo",
	}

	//fmt.Println(BookInterface(my1))
	fmt.Println(BookInterface(my2))
}

func maintest() {
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
