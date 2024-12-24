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
	var position string
	for {
		fmt.Printf("Игрок %d, выберите позицию для удаления фишки противника: ", player)
		fmt.Scan(&position)
		positionInt, err := strconv.Atoi(position)
		if err != nil {
			fmt.Println("Некорректный ввод, попробуйте снова.")
		} else {
			if positionInt >= 0 && positionInt < 16 && board[positionInt] == opponent {
				board[positionInt] = 0 // Удаляем фишку противника
				break
			}
			fmt.Println("Некорректный ввод или позиция не занята противником, попробуйте снова.")
		}
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

func isValidMove(board [16]int, neighbors map[int][]int, currentPlayer, from, to int) bool {
	// Проверка, что фишка принадлежит текущему игроку и цель свободна
	if board[from] != currentPlayer || board[to] != 0 {
		return false
	}

	// Проверка, что 'to' - сосед 'from' в зависимости от массива neighbors
	for _, neighbor := range neighbors[from] {
		if neighbor == to {
			return true // Перемещение на соседнюю позицию
		}
	}
	return false
}

func main() {
	board := [16]int{} // Изначально заполненный нулями массив
	currentPlayer := 1
	millsBuilt := make(map[int][][]int) // Хранит список построенных мельниц для каждого игрока

	fmt.Println(board)
	printBoard(board)
	var gameModeInt int
	var gameMode string
	for {
		fmt.Println("Выберите режим игры:")
		fmt.Println("1 - Игрок против игрока")
		fmt.Println("2 - Игрок против компьютера")
		fmt.Print("Ваш выбор: ")
		fmt.Scan(&gameMode)
		var err error
		gameModeInt, err = strconv.Atoi(gameMode)
		if err != nil {
			fmt.Println("Некорректный ввод, попробуйте снова.")
		} else {
			break
		}
	}

	if gameModeInt == 1 {
		// Игроки расставляют по 6 фишек на поле - 1-ый этап
		for turns := 0; turns < 12; turns++ {
			// Вывод поля
			printBoard(board)
			// Ставим фишку
			placePieces(&board, currentPlayer)
			// Проверка на наличие мельницы после хода
			mills := checkAndGetMills(board, currentPlayer) // Получаем построенную мельницу
			if len(mills) > 0 {                             // Проверяем, есть ли найденные мельницы
				for _, mill := range mills {
					// Проверяем, была ли уже построена эта мельница
					if !isMillAlreadyBuilt(millsBuilt[currentPlayer], mill) {
						millsBuilt[currentPlayer] = append(millsBuilt[currentPlayer], mill) // Добавляем построенную мельницу в список
						fmt.Printf("Игрок %d построил новую мельницу! Текущие мельницы: %v\n", currentPlayer, millsBuilt[currentPlayer])
						// Удаляем фишку противника
						removeOpponentPiece(&board, currentPlayer)
					} else {
						fmt.Printf("Игрок %d построил мельницу, но она уже была построена ранее.\n", currentPlayer)
					}
				}
			}

			// Смена игрока и повтор
			currentPlayer = 3 - currentPlayer
		}
		printBoard(board)

		// 2-ой этап: движение фишек
		var neighbors = map[int][]int{
			0:  {1, 6},       // клетка 0
			1:  {0, 2, 4},    // клетка 1
			2:  {1, 9},       // клетка 2
			3:  {4, 7},       // клетка 3
			4:  {1, 3, 5},    // клетка 4
			5:  {4, 8},       // клетка 5
			6:  {0, 7, 13},   // клетка 6
			7:  {3, 6, 10},   // клетка 7
			8:  {5, 9, 12},   // клетка 8
			9:  {2, 8, 15},   // клетка 9
			10: {7, 11},      // клетка 10
			11: {10, 12, 14}, // клетка 11
			12: {8, 11},      // клетка 12
			13: {6, 14},      // клетка 13
			14: {11, 13, 15}, // клетка 14
			15: {9, 14},      // клетка 15
		}

		for {
			printBoard(board)
			fmt.Printf("Игрок %d, выберите перемещение\n", currentPlayer)

			var from, to int

			// Проверка ввода для перемещения "from"
			for {
				fmt.Print("С какой позиции хотите переместить фишку? ")
				var position string
				fmt.Scan(&position)
				positionInt, err := strconv.Atoi(position)
				if err != nil || positionInt < 0 || positionInt >= len(board) {
					fmt.Println("Некорректный ввод, попробуйте снова.")
					continue
				}
				from = positionInt
				break
			}

			// Проверка ввода для перемещения "to"
			for {
				fmt.Print("На какую позицию хотите переместить фишку? ")
				var position string
				fmt.Scan(&position)
				positionInt, err := strconv.Atoi(position)
				if err != nil || positionInt < 0 || positionInt >= len(board) {
					fmt.Println("Некорректный ввод, попробуйте снова.")
					continue
				}
				to = positionInt
				break
			}

			// Проверяем, допустимо ли перемещение
			if isValidMove(board, neighbors, currentPlayer, from, to) {
				board[to] = board[from]                         // Перемещаем фишку
				board[from] = 0                                 // Освобождаем исходное место
				mills := checkAndGetMills(board, currentPlayer) // Проверяем наличие мельниц
				if len(mills) > 0 {
					for _, mill := range mills {
						if !isMillAlreadyBuilt(millsBuilt[currentPlayer], mill) {
							millsBuilt[currentPlayer] = append(millsBuilt[currentPlayer], mill)
							fmt.Printf("Игрок %d построил новую мельницу! Текущие мельницы: %v\n", currentPlayer, millsBuilt[currentPlayer])
							removeOpponentPiece(&board, currentPlayer)
						}
					}
				}
			} else {
				fmt.Println("Неверный ход, попробуйте снова")
				continue
			}

			// Смена игрока
			currentPlayer = 3 - currentPlayer
		}
	}
}
