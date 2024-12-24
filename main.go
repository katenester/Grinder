package main

import (
	"fmt"
	"strconv"
)

func printCell(i int, board [16]int) {
	if board[i] == 0 {
		fmt.Print(".")
	} else if board[i] == 1 {
		fmt.Print("X")
	} else {
		fmt.Print("O")
	}
}

func printBoard(board [16]int) {
	printHelpBoard()
	fmt.Println("Текущее поле:")

	// Первая строка
	printCell(0, board)
	for i := 0; i < 5; i++ {
		fmt.Print("#")
	}
	printCell(1, board)
	for i := 0; i < 5; i++ {
		fmt.Print("#")
	}
	printCell(2, board)
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
	printCell(3, board)
	fmt.Print("#")
	printCell(4, board)
	fmt.Print("#")
	printCell(5, board)
	for i := 0; i < 3; i++ {
		fmt.Print(" ")
	}
	fmt.Print("#")
	fmt.Println()

	// Четвертая строка
	printCell(6, board)
	fmt.Print(" ")
	fmt.Print("-")
	fmt.Print(" ")
	printCell(7, board)
	for i := 0; i < 3; i++ {
		fmt.Print(" ")
	}
	printCell(8, board)
	fmt.Print(" ")
	fmt.Print("-")
	fmt.Print(" ")
	printCell(9, board)
	fmt.Println()

	// Пятая строка
	fmt.Print("#")
	for i := 0; i < 3; i++ {
		fmt.Print(" ")
	}
	printCell(10, board)
	fmt.Print("#")
	printCell(11, board)
	fmt.Print("#")
	printCell(12, board)
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
	printCell(13, board)
	for i := 0; i < 5; i++ {
		fmt.Print("#")
	}
	printCell(14, board)
	for i := 0; i < 5; i++ {
		fmt.Print("#")
	}
	printCell(15, board)
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

// Функция для расстановки фишек
func placePieces(board *[16]int, player int) {
	var position string
	for {
		fmt.Printf("Игрок %d, выберите позицию для размещения фишки (0-15): ", player)
		fmt.Scan(&position)
		positionInt, err := strconv.Atoi(position)
		if err != nil {
			fmt.Println("Некорректный ввод, попробуйте снова.")
		} else {
			fmt.Println(position)
			if positionInt >= 0 && positionInt < 16 && board[positionInt] == 0 { // Проверяем, свободно ли место
				board[positionInt] = player
				break
			}
			fmt.Println("Некорректный ввод, попробуйте снова.")
		}
	}
}

// Функция для проверки мельниц
func checkAndGetMills(board [16]int, player int) [][]int {
	mills := [][3]int{
		{0, 1, 2}, {3, 4, 5}, {10, 11, 12}, {13, 14, 15}, // Горизонтальные линии
		{0, 6, 13}, {3, 7, 10}, {5, 8, 12}, {2, 9, 15}, // Вертикальные линии
	}

	foundMills := [][]int{} // Срез для хранения найденных мельниц
	for _, mill := range mills {
		if board[mill[0]] == player && board[mill[1]] == player && board[mill[2]] == player {
			// Преобразуем [3]int в []int перед добавлением
			foundMills = append(foundMills, []int{mill[0], mill[1], mill[2]}) // Добавляем найденную мельницу
		}
	}
	return foundMills // Возвращаем все найденные мельницы
}

func removeOpponentPiece(board *[16]int, player int) {
	opponent := 3 - player
	var position int
	for {
		fmt.Printf("Игрок %d, выберите позицию для удаления фишки противника: ", player)
		fmt.Scan(&position)
		if position >= 0 && position < 16 && board[position] == opponent {
			board[position] = 0 // Удаляем фишку противника
			break
		}
		fmt.Println("Некорректный ввод или позиция не занята противником, попробуйте снова.")
	}
}

// Функция для проверки, была ли уже построена мельница
func isMillAlreadyBuilt(builtMills [][]int, mill []int) bool {
	for _, builtMill := range builtMills {
		if len(builtMill) == 3 && builtMill[0] == mill[0] && builtMill[1] == mill[1] && builtMill[2] == mill[2] {
			return true
		}
	}
	return false
}

/*func isNumber(value interface{}) bool {
	switch value.(type) {
	case int, int8, int16, int32, int64, float32, float64:
		return true
	default:
		return false
	}
}*/

func main() {
	board := [16]int{} // Изначально заполненный нулями массив
	currentPlayer := 1
	millsBuilt := make(map[int][][]int) // Хранит список построенных мельниц для каждого игрока

	fmt.Println(board)
	printBoard(board)
	var gameMode int
	fmt.Println("Выберите режим игры:")
	fmt.Println("1 - Игрок против игрока")
	fmt.Println("2 - Игрок против компьютера")
	fmt.Print("Ваш выбор: ")
	fmt.Scan(&gameMode)

	if gameMode == 1 {
		for turns := 0; turns < 6; turns++ {
			currentPlayer = 1
			printBoard(board)
			placePieces(&board, currentPlayer)
			// Проверка на наличие мельницы после хода
			mills := checkAndGetMills(board, currentPlayer) // Получаем построенную мельницу
			if len(mills) > 0 {                             // Проверяем, есть ли найденные мельницы
				for _, mill := range mills {
					// Проверяем, была ли уже построена эта мельница
					if !isMillAlreadyBuilt(millsBuilt[currentPlayer], mill) {
						millsBuilt[currentPlayer] = append(millsBuilt[currentPlayer], mill) // Добавляем построенную мельницу в список
						fmt.Printf("Игрок %d построил новую мельницу! Текущие мельницы: %v\n", currentPlayer, millsBuilt[currentPlayer])
						removeOpponentPiece(&board, currentPlayer)
					} else {
						fmt.Printf("Игрок %d построил мельницу, но она уже была построена ранее.\n", currentPlayer)
					}
				}
			}

			// Смена игрока и повтор
			currentPlayer = 2

			printBoard(board)
			placePieces(&board, currentPlayer)
			// Проверка на наличие мельницы после хода
			mills = checkAndGetMills(board, currentPlayer) // Получаем построенную мельницу
			if len(mills) > 0 {                            // Проверяем, есть ли найденные мельницы
				for _, mill := range mills {
					// Проверяем, была ли уже построена эта мельница
					if !isMillAlreadyBuilt(millsBuilt[currentPlayer], mill) {
						millsBuilt[currentPlayer] = append(millsBuilt[currentPlayer], mill) // Добавляем построенную мельницу в список
						fmt.Printf("Игрок %d построил новую мельницу! Текущие мельницы: %v\n", currentPlayer, millsBuilt[currentPlayer])
						removeOpponentPiece(&board, currentPlayer)
					} else {
						fmt.Printf("Игрок %d построил мельницу, но она уже была построена ранее.\n", currentPlayer)
					}
				}
			}
		}
		printBoard(board)
	}

}
