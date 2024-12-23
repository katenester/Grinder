package main

import (
	"fmt"
)

var board [16]int
var currentPlayer int

func printCell(i int) {
	if board[i] == 0 {
		fmt.Print(".")
	} else if board[i] == 1 {
		fmt.Print("X")
	} else {
		fmt.Print("O")
	}
}

func printBoard() {
	printHelpBoard()
	fmt.Println("Текущее поле:")

	// Первая строка
	printCell(0)
	for i := 0; i < 5; i++ {
		fmt.Print("#")
	}
	printCell(1)
	for i := 0; i < 5; i++ {
		fmt.Print("#")
	}
	printCell(2)
	fmt.Println()

	// Вторая строка
	fmt.Print("#")
	for i := 0; i < 5; i++ {
		fmt.Print(" ")
	}
	fmt.Print("|")
	for i := 0; i < 5; i++ {
		fmt.Print(" ")
	}
	fmt.Print("#")
	fmt.Println()

	// Третья строка
	fmt.Print("#")
	for i := 0; i < 3; i++ {
		fmt.Print(" ")
	}
	printCell(3)
	fmt.Print("#")
	printCell(4)
	fmt.Print("#")
	printCell(5)
	for i := 0; i < 3; i++ {
		fmt.Print(" ")
	}
	fmt.Print("#")
	fmt.Println()

	// Четвертая строка
	printCell(6)
	fmt.Print(" ")
	fmt.Print("-")
	fmt.Print(" ")
	printCell(7)
	for i := 0; i < 3; i++ {
		fmt.Print(" ")
	}
	printCell(8)
	fmt.Print(" ")
	fmt.Print("-")
	fmt.Print(" ")
	printCell(9)
	fmt.Println()

	// Пятая строка
	fmt.Print("#")
	for i := 0; i < 3; i++ {
		fmt.Print(" ")
	}
	printCell(10)
	fmt.Print("#")
	printCell(11)
	fmt.Print("#")
	printCell(12)
	for i := 0; i < 3; i++ {
		fmt.Print(" ")
	}
	fmt.Print("#")
	fmt.Println()

	// Шестая строка
	fmt.Print("#")
	for i := 0; i < 5; i++ {
		fmt.Print(" ")
	}
	fmt.Print("|")
	for i := 0; i < 5; i++ {
		fmt.Print(" ")
	}
	fmt.Print("#")
	fmt.Println()

	// Седьмая строка
	printCell(13)
	for i := 0; i < 5; i++ {
		fmt.Print("#")
	}
	printCell(14)
	for i := 0; i < 5; i++ {
		fmt.Print("#")
	}
	printCell(15)
	fmt.Println()

	fmt.Println()
}

func printHelpBoard() {
	fmt.Println("Вспомогательное поле:")
	fmt.Println("0#####1#####2")
	fmt.Println("#     |     #")
	fmt.Println("#   3#4#5   #")
	fmt.Println("6---7   8---9")
	fmt.Println("#   101112  #")
	fmt.Println("#     |     #")
	fmt.Println("13####14###15")
	fmt.Println()
}

func main() {
	currentPlayer = 1
	fmt.Println(board)
	printBoard()
}
