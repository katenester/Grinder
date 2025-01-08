package service

import (
	"Grinder/server/internal/models"
	"Grinder/server/internal/repository"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

type GameService struct {
	repo repository.Game
}

func NewGameService(repo repository.Game) *GameService {
	return &GameService{
		repo: repo,
	}
}

func (g *GameService) CreateRoom(players []models.Player) (int, error) {
	return g.repo.CreateRoom(players)
}

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

// GameServer - логика игры c сервером
func (g *GameService) GameServer(idRoom int) error {
	count1 := 6
	count2 := 6
	//currentPlayer := 1 // 1-текущий (первый) игрок , 2 - второй/ сервер
	players := g.repo.GetPlayer(idRoom)
	// Игроки расставляют по 6 фишек на поле - 1-ый этап
	for turns := 0; turns < 12; turns++ {
		// Вывод поля для игрока (отправка игроку)
		err := sendPlayer(printBoard(g.repo.GetBoard(idRoom)), players[turns%2])
		if err != nil {
			log.Print("GameSever error")
			return err
		}
		// Ставим фишку
		// Если пользователь
		if turns%2 == 0 {
			g.placePieces(idRoom, turns%2, players)
		} else {
			// Логика для компьютера
			g.computerPlacePiece(idRoom, turns%2, players)
		}

		// Проверка на наличие мельницы после хода
		mills := checkAndGetMills(g.repo.GetBoard(idRoom), turns%2) // Получаем построенную мельницу
		if len(mills) > 0 {                                         // Проверяем, есть ли найденные мельницы
			for _, mill := range mills {
				// Проверяем, была ли уже построена эта мельница
				if !isMillAlreadyBuilt(g.repo.GetMillsBuilt(idRoom, turns%2), mill) {
					g.repo.AppendMillsBuilt(idRoom, turns%2, mill) // Добавляем построенную мельницу в список
					//err = sendPlayer("Игрок "+player.Name+", выберите позицию для размещения фишки (0-15): ", players[player-1])
					//if err != nil {
					//	log.Println(err)
					//}
					err = sendPlayer(fmt.Sprintf("Игрок %s построил новую мельницу! Текущие мельницы: %v\n", players[turns%2].Name, g.repo.GetMillsBuilt(idRoom, turns%2)), players[0])
					if err != nil {
						log.Println(err)
					}
					// Удаляем фишку противника
					if turns%2 == 0 {
						g.removeOpponentPiece(idRoom, turns%2, players)
						count2--
					} else {
						g.removeOpponentPieceComputer(idRoom, turns%2, players)
						count1--
					}
				} else {
					err = sendPlayer(fmt.Sprintf("Игрок %s построил мельницу, но она уже была построена ранее.\n", players[turns%2].Name), players[0])
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
		//// Смена игрока и повтор
		//currentPlayer = 3 - currentPlayer
	}
	currentPlayer := 0
	// Второй этап
	for {
		// Вывод поля для игрока (отправка игроку)
		err := sendPlayer(printBoard(g.repo.GetBoard(idRoom)), players[0])
		if err != nil {
			log.Print("GameSever error")
			return err
		}
		if currentPlayer == 0 {
			// Ход игрока
			sendPlayer(fmt.Sprintf("Игрок %d, выберите перемещение\n", currentPlayer), players[0])
			var from, to int

			// Проверка ввода для перемещения "from"
			for {
				sendPlayer("С какой позиции хотите переместить фишку? ", players[0])
				position, err := receivePlayer(players[currentPlayer])
				positionInt, err := strconv.Atoi(position)
				if err != nil || positionInt < 0 || positionInt >= len(g.repo.GetBoard(idRoom)) {
					sendPlayer("Некорректный ввод, попробуйте снова. ", players[0])
					continue
				}
				from = positionInt
				break
			}

			// Проверка ввода для перемещения "to"
			for {
				sendPlayer("На какую позицию хотите переместить фишку?  ", players[0])
				var position string
				// Получаем позицию от игрока
				position, err = receivePlayer(players[0])
				if err != nil {
					log.Println("placePieces error 2")
				}
				positionInt, err := strconv.Atoi(position)
				if err != nil || positionInt < 0 || positionInt >= len(g.repo.GetBoard(idRoom)) {
					sendPlayer("Некорректный ввод, попробуйте снова. ", players[0])
					continue
				}
				to = positionInt
				break
			}

			// Проверяем допустимость перемещения
			if isValidMove(g.repo.GetBoard(idRoom), neighbors, currentPlayer, from, to, count1) {
				board := g.repo.GetBoard(idRoom)
				board[to] = board[from] // Перемещаем фишку
				board[from] = 0         // Освобождаем исходное место
				g.repo.SetBoard(idRoom, board)
				mills := checkAndGetMills(g.repo.GetBoard(idRoom), currentPlayer)
				if len(mills) > 0 {
					for _, mill := range mills {
						if !isMillAlreadyBuilt(g.repo.GetMillsBuilt(idRoom, currentPlayer), mill) {
							g.repo.AppendMillsBuilt(idRoom, currentPlayer, mill)
							sendPlayer(fmt.Sprintf("Игрок %d построил новую мельницу! Текущие мельницы: %v\n", currentPlayer, g.repo.GetMillsBuilt(idRoom, currentPlayer)), players[0])
							g.removeOpponentPiece(idRoom, currentPlayer, players)
							count2--
						}
					}
				}
				if count1 == 2 || isLocked(g.repo.GetBoard(idRoom), neighbors, 1) {
					sendPlayer(fmt.Sprintf("Игрок %s победил!", players[1].Name), players[0])
					break
				}
				if count2 == 2 || isLocked(g.repo.GetBoard(idRoom), neighbors, 2) {
					sendPlayer(fmt.Sprintf("Игрок %s победил!", players[0].Name), players[0])
					break
				}
			} else {
				sendPlayer("Неверный ход, попробуйте снова", players[0])
			}
		} else {
			// Ход компьютера
			sendPlayer(fmt.Sprintf("Игрок %s (Компьютер) делает ход...", players[1].Name), players[0])
			// Логика выбора хода компьютера
			from, to := g.computerMove(g.repo.GetBoard(idRoom), neighbors, currentPlayer, count2)
			board := g.repo.GetBoard(idRoom)
			board[to] = board[from] // Перемещаем фишку
			board[from] = 0         // Освобождаем исходное место
			// Устанавливаем в репозитории
			g.repo.SetBoard(idRoom, board)
			sendPlayer(fmt.Sprintf("Компьютер переместил фишку с %d на %d\n", from, to), players[0])
			mills := checkAndGetMills(board, currentPlayer)
			if len(mills) > 0 {
				for _, mill := range mills {
					if !isMillAlreadyBuilt(g.repo.GetMillsBuilt(idRoom, currentPlayer), mill) {
						g.repo.AppendMillsBuilt(idRoom, currentPlayer, mill)
						sendPlayer(fmt.Sprintf("Игрок %s построил новую мельницу! Текущие мельницы: %v\n", players[1].Name, g.repo.GetMillsBuilt(idRoom, currentPlayer)), players[0])
						g.removeOpponentPieceComputer(idRoom, currentPlayer, players)
						count1--
					}
				}
			}
			if count1 == 2 || isLocked(g.repo.GetBoard(idRoom), neighbors, 1) {
				sendPlayer(fmt.Sprintf("Игрок %s победил!", players[1].Name), players[0])
				break
			}
			if count2 == 2 || isLocked(g.repo.GetBoard(idRoom), neighbors, 2) {
				sendPlayer(fmt.Sprintf("Игрок %s победил!", players[0].Name), players[0])
				break
			}
		}
		// Смена игрока
		currentPlayer = (currentPlayer + 1) % 2
	}
	return nil
}

// Game - логика игры
func (g *GameService) GameUser(idRoom int) error {
	count1 := 6
	count2 := 6
	players := g.repo.GetPlayer(idRoom)
	// Игроки расставляют по 6 фишек на поле - 1-ый этап
	for turns := 0; turns < 12; turns++ {
		// Вывод поля для игрока (отправка игроку)
		sendPlayer(printBoard(g.repo.GetBoard(idRoom)), players[0])
		// Вывод поля для игрока (отправка игроку)
		sendPlayer(printBoard(g.repo.GetBoard(idRoom)), players[1])
		// Ставим фишку
		g.placePieces(idRoom, turns%2, players)
		// Проверка на наличие мельницы после хода
		mills := checkAndGetMills(g.repo.GetBoard(idRoom), turns%2) // Получаем построенную мельницу
		if len(mills) > 0 {                                         // Проверяем, есть ли найденные мельницы
			for _, mill := range mills {
				// Проверяем, была ли уже построена эта мельница
				if !isMillAlreadyBuilt(g.repo.GetMillsBuilt(idRoom, turns%2), mill) {
					g.repo.AppendMillsBuilt(idRoom, turns%2, mill) // Добавляем построенную мельницу в список
					sendPlayer(fmt.Sprintf("Игрок %s построил новую мельницу! Текущие мельницы: %v\n", players[turns%2].Name, g.repo.GetMillsBuilt(idRoom, turns%2)), players[0])
					sendPlayer(fmt.Sprintf("Игрок %s построил новую мельницу! Текущие мельницы: %v\n", players[turns%2].Name, g.repo.GetMillsBuilt(idRoom, turns%2)), players[1])
					// Удаляем фишку противника
					g.removeOpponentPiece(idRoom, turns%2, players)
					if turns%2 == 0 {
						count2--
					} else {
						count1--
					}
				} else {
					err := sendPlayer(fmt.Sprintf("Игрок %s построил мельницу, но она уже была построена ранее.\n", players[turns%2].Name), players[0])
					if err != nil {
						log.Println(err)
					}
					err = sendPlayer(fmt.Sprintf("Игрок %s построил мельницу, но она уже была построена ранее.\n", players[turns%2].Name), players[1])
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
	}

	currentPlayer := 0
	// 2-ой этап: движение фишек
	for {
		// Вывод поля для игрока (отправка игроку)
		sendPlayer(printBoard(g.repo.GetBoard(idRoom)), players[0])
		// Вывод поля для игрока (отправка игроку)
		sendPlayer(printBoard(g.repo.GetBoard(idRoom)), players[1])
		sendPlayer(fmt.Sprintf("Игрок %s, выберите перемещение\n", players[currentPlayer].Name), players[currentPlayer])

		var from, to int

		// Проверка ввода для перемещения "from"
		for {
			sendPlayer("С какой позиции хотите переместить фишку? ", players[currentPlayer])
			position, err := receivePlayer(players[currentPlayer])
			positionInt, err := strconv.Atoi(position)
			if err != nil || positionInt < 0 || positionInt >= len(g.repo.GetBoard(idRoom)) {
				sendPlayer("Некорректный ввод, попробуйте снова. ", players[currentPlayer])
				continue
			}
			from = positionInt
			break
		}

		// Проверка ввода для перемещения "to"
		for {
			sendPlayer("На какую позицию хотите переместить фишку?  ", players[currentPlayer])
			// Получаем позицию от игрока
			position, err := receivePlayer(players[currentPlayer])
			if err != nil {
				log.Println("placePieces error 2")
			}
			positionInt, err := strconv.Atoi(position)
			if err != nil || positionInt < 0 || positionInt >= len(g.repo.GetBoard(idRoom)) {
				sendPlayer("Некорректный ввод, попробуйте снова. ", players[currentPlayer])
				continue
			}
			to = positionInt
			break
		}

		var localCount int
		if currentPlayer == 0 {
			localCount = count1
		} else {
			localCount = count2
		}

		// Проверяем, допустимо ли перемещение
		if isValidMove(g.repo.GetBoard(idRoom), neighbors, currentPlayer, from, to, localCount) {
			board := g.repo.GetBoard(idRoom)
			board[to] = board[from] // Перемещаем фишку
			board[from] = 0         // Освобождаем исходное место
			g.repo.SetBoard(idRoom, board)
			mills := checkAndGetMills(g.repo.GetBoard(idRoom), currentPlayer)
			if len(mills) > 0 {
				for _, mill := range mills {
					if !isMillAlreadyBuilt(g.repo.GetMillsBuilt(idRoom, currentPlayer), mill) {
						g.repo.AppendMillsBuilt(idRoom, currentPlayer, mill)
						sendPlayer(fmt.Sprintf("Игрок %d построил новую мельницу! Текущие мельницы: %v\n", currentPlayer, g.repo.GetMillsBuilt(idRoom, currentPlayer)), players[0])
						g.removeOpponentPiece(idRoom, currentPlayer, players)
						if currentPlayer == 0 {
							count2--
						} else {
							count1--
						}
					}
				}
			}
			if count1 == 2 || isLocked(g.repo.GetBoard(idRoom), neighbors, 0) {
				sendPlayer(fmt.Sprintf("Игрок %s победил!", players[1].Name), players[0])
				sendPlayer(fmt.Sprintf("Игрок %s победил!", players[1].Name), players[1])
				break
			}
			if count2 == 2 || isLocked(g.repo.GetBoard(idRoom), neighbors, 1) {
				sendPlayer(fmt.Sprintf("Игрок %s победил!", players[0].Name), players[0])
				sendPlayer(fmt.Sprintf("Игрок %s победил!", players[0].Name), players[1])
				break
			}
		} else {
			sendPlayer("Неверный ход, попробуйте снова", players[currentPlayer])
		}

		// Смена игрока
		currentPlayer = (currentPlayer + 1) % 2
	}
	return nil
}
func printCell(i int, board [16]int) string {
	if board[i] == 0 {
		return "."
	} else if board[i] == 1 {
		return "X"
	} else {
		return "O"
	}
}

func printBoard(board [16]int) string {
	var sb strings.Builder
	sb.WriteString("Вспомогательное поле:\n")
	sb.WriteString("0#####1#####2\n")
	sb.WriteString("#     |     #\n")
	sb.WriteString("#   3#4#5   #\n")
	sb.WriteString("6---7   8---9\n")
	sb.WriteString("#   101112  #\n")
	sb.WriteString("#     |     #\n")
	sb.WriteString("13####14###15\n")
	sb.WriteString("\n")
	sb.WriteString("Текущее поле:\n")

	// Первая строка
	sb.WriteString(printCell(0, board))
	for i := 0; i < 5; i++ {
		sb.WriteString("#")
	}
	sb.WriteString(printCell(1, board))
	for i := 0; i < 5; i++ {
		sb.WriteString("#")
	}
	sb.WriteString(printCell(2, board))
	sb.WriteString("\n")

	// Вторая строка
	sb.WriteString("#")
	for i := 0; i < 5; i++ {
		sb.WriteString(" ")
	}
	sb.WriteString("|")
	for i := 0; i < 5; i++ {
		sb.WriteString(" ")
	}
	sb.WriteString("#")
	sb.WriteString("\n")

	// Третья строка
	sb.WriteString("#")
	for i := 0; i < 3; i++ {
		sb.WriteString(" ")
	}
	sb.WriteString(printCell(3, board))
	sb.WriteString("#")
	sb.WriteString(printCell(4, board))
	sb.WriteString("#")
	sb.WriteString(printCell(5, board))
	for i := 0; i < 3; i++ {
		sb.WriteString(" ")
	}
	sb.WriteString("#")
	sb.WriteString("\n")

	// Четвертая строка
	sb.WriteString(printCell(6, board))
	sb.WriteString(" ")
	sb.WriteString("-")
	sb.WriteString(" ")
	sb.WriteString(printCell(7, board))
	for i := 0; i < 3; i++ {
		sb.WriteString(" ")
	}
	sb.WriteString(printCell(8, board))
	sb.WriteString(" ")
	sb.WriteString("-")
	sb.WriteString(" ")
	sb.WriteString(printCell(9, board))
	sb.WriteString("\n")

	// Пятая строка
	sb.WriteString("#")
	for i := 0; i < 3; i++ {
		sb.WriteString(" ")
	}
	sb.WriteString(printCell(10, board))
	sb.WriteString("#")
	sb.WriteString(printCell(11, board))
	sb.WriteString("#")
	sb.WriteString(printCell(12, board))
	for i := 0; i < 3; i++ {
		sb.WriteString(" ")
	}
	sb.WriteString("#")
	sb.WriteString("\n")

	// Шестая строка
	sb.WriteString("#")
	for i := 0; i < 5; i++ {
		sb.WriteString(" ")
	}
	sb.WriteString("|")
	for i := 0; i < 5; i++ {
		sb.WriteString(" ")
	}
	sb.WriteString("#")
	sb.WriteString("\n")

	// Седьмая строка
	sb.WriteString(printCell(13, board))
	for i := 0; i < 5; i++ {
		sb.WriteString("#")
	}
	sb.WriteString(printCell(14, board))
	for i := 0; i < 5; i++ {
		sb.WriteString("#")
	}
	sb.WriteString(printCell(15, board))
	sb.WriteString("\n")
	sb.WriteString("\n")
	return sb.String()
}

// Функция для расстановки фишек
func (g *GameService) placePieces(idRoom int, player int, players []models.Player) {
	board := g.repo.GetBoard(idRoom)
	var position string
	for {
		// Отправляем message игрокам
		err := sendPlayer("Игрок "+players[player].Name+", выберите позицию для размещения фишки (0-15): ", players[player])
		if err != nil {
			log.Println("placePieces error")
		}
		// Получаем позицию от игрока
		position, err = receivePlayer(players[player])
		if err != nil {
			log.Println("placePieces error 2")
		}
		positionInt, err := strconv.Atoi(position)
		if err != nil {
			err = sendPlayer("Некорректный ввод, попробуйте снова.", players[player])
			if err != nil {
				log.Println("placePieces error 3")
			}
		} else {
			//fmt.Println(position)
			if positionInt >= 0 && positionInt < 16 && board[positionInt] == 0 { // Проверяем, свободно ли место
				// Меняем доску локально
				board[positionInt] = player
				// Меняем глобально (на уравне ссылок и репозитория)
				g.repo.SetBoard(idRoom, board)
				break
			}
			err = sendPlayer("Некорректный ввод, попробуйте снова.", players[player])
			if err != nil {
				log.Println("placePieces error 4")
			}
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

func (g *GameService) removeOpponentPiece(idRoom int, player int, players []models.Player) {
	opponent := (player + 1) % 2
	board := g.repo.GetBoard(idRoom)
	var position string
	for {
		err := sendPlayer(fmt.Sprintf("Игрок %s, выберите позицию для удаления фишки противника: ", players[player].Name), players[player])
		if err != nil {
			log.Println("removeOpponentPiece error")
		}

		// Получаем позицию от игрока
		position, err = receivePlayer(players[player])
		if err != nil {
			log.Println("removeOpponentPiece error 2")
		}
		positionInt, err := strconv.Atoi(position)
		if err != nil {
			err = sendPlayer("Некорректный ввод, попробуйте снова.", players[player])
			if err != nil {
				log.Println("removeOpponentPiece error 3")
			}
		} else {
			if positionInt >= 0 && positionInt < 16 && board[positionInt] == opponent {
				// Меняем доску локально
				board[positionInt] = 0
				// Меняем глобально (на уравне ссылок и репозитория)
				g.repo.SetBoard(idRoom, board)
				break
			}
			err = sendPlayer("Некорректный ввод, попробуйте снова.", players[player])
			if err != nil {
				log.Println("removeOpponentPiece error 4")
			}
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

func isValidMove(board [16]int, neighbors map[int][]int, currentPlayer, from, to, count int) bool {
	// Проверка, что фишка принадлежит текущему игроку и цель свободна
	if board[from] != currentPlayer || board[to] != 0 {
		return false
	}

	// Проверка, что 'to' - сосед 'from' в зависимости от массива neighbors
	if count > 3 {
		for _, neighbor := range neighbors[from] {
			if neighbor == to {
				return true // Перемещение на соседнюю позицию
			}
		}
		return false
	}
	return true
}

func isLocked(board [16]int, neighbors map[int][]int, currentPlayer int) bool {
	for i, piece := range board {
		if piece == currentPlayer { // Если на позиции стоит фишка текущего игрока
			// Проверяем соседние клетки
			canMove := false
			for _, neighbor := range neighbors[i] {
				if board[neighbor] == 0 { // Если соседняя клетка свободна
					canMove = true
					break
				}
			}
			if canMove { // Если есть хотя бы один возможный ход
				return false
			}
		}
	}
	return true // Если ни одна фишка не может сделать ход
}

func (g *GameService) computerPlacePiece(idRoom int, currentPlayer int, players []models.Player) {
	//rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел
	freePositions := []int{} // Массив для хранения свободных позиций
	board := g.repo.GetBoard(idRoom)
	// Находим все свободные позиции
	for i := 0; i < len(board); i++ {
		if board[i] == 0 { // Если клетка свободна (0)
			freePositions = append(freePositions, i) // Добавляем номер клетки
		}
	}

	if len(freePositions) > 0 {
		// Выбираем случайную позицию из доступных
		randomIndex := rand.Intn(len(freePositions))
		selectedPosition := freePositions[randomIndex]

		// Ставим фишку компьютера на выбранную позицию
		sendPlayer(fmt.Sprintf("Компьютер поставил фишку на позицию %d\n", selectedPosition), players[0])

		// Меняем доску локально
		board[selectedPosition] = currentPlayer
		// Меняем глобально (на уравне ссылок и репозитория)
		g.repo.SetBoard(idRoom, board)
		//err := sendPlayer("Некорректный ввод, попробуйте снова.", player)
	}
	//else {
	//	//// Можно будет убрать
	//	//err := sendPlayer("Некорректный ввод, попробуйте снова.", player)
	//	//if err != nil {
	//	//	log.Println("error computerPlacePiece")
	//	//}
	//}
}

func (g *GameService) removeOpponentPieceComputer(idRoom int, player int, players []models.Player) {
	opponent := (player + 1) % 2
	opponentPositions := []int{} // Массив для хранения позиций фишек противника
	board := g.repo.GetBoard(idRoom)
	// Находим все позиции фишек противника
	for i := 0; i < len(board); i++ {
		if board[i] == opponent {
			opponentPositions = append(opponentPositions, i)
		}
	}

	if len(opponentPositions) > 0 {
		// Выбираем случайную позицию для удаления
		randomIndex := rand.Intn(len(opponentPositions))
		selectedPosition := opponentPositions[randomIndex]
		// Ставим фишку компьютера на выбранную позицию
		sendPlayer(fmt.Sprintf("Компьютер удалил фишку противника с позиции %d\n", selectedPosition), players[0])
		// Меняем доску локально
		board[selectedPosition] = 0
		// Меняем глобально (на уравне ссылок и репозитория)
		g.repo.SetBoard(idRoom, board)
	}
	//else {
	//	// Можно убрать
	//	fmt.Println("Нет фишек противника для удаления.")
	//}
}
func (g *GameService) computerMove(board [16]int, neighbors map[int][]int, currentPlayer int, count int) (int, int) {
	availableMoves := []string{}

	// Находим все доступные перемещения для компьютера
	for from := 0; from < len(board); from++ {
		if board[from] == currentPlayer { // Если на позиции стоит фишка компьютера
			if count > 3 {
				// Можно перемещаться только на соседние свободные клетки
				for _, to := range neighbors[from] {
					if isValidMove(board, neighbors, currentPlayer, from, to, count) {
						availableMoves = append(availableMoves, fmt.Sprintf("%d->%d", from, to))
					}
				}
			} else {
				// Можно перемещаться на любую свободную клетку
				for to := 0; to < len(board); to++ {
					if isValidMove(board, neighbors, currentPlayer, from, to, count) {
						availableMoves = append(availableMoves, fmt.Sprintf("%d->%d", from, to))
					}
				}
			}
		}
	}

	if len(availableMoves) > 0 {
		// Выбираем случайное перемещение
		move := availableMoves[rand.Intn(len(availableMoves))]
		from, to := parseMove(move)
		return from, to // Возвращаем from и to, если ход возможен
	}

	return -1, -1 // Возвращаем -1 для from и to, если нет доступных ходов
}

// Функция для парсинга перемещения из строки
func parseMove(move string) (int, int) {
	var from, to int
	fmt.Sscanf(move, "%d->%d", &from, &to)
	return from, to
}

func sendPlayer(message string, player models.Player) error {
	// Если сервер => ничего не отправляем
	if player.Conn == nil {
		return nil
	}
	dat, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = player.Conn.Write(dat)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func receivePlayer(player models.Player) (string, error) {
	dat := make([]byte, 1024)
	n, err := player.Conn.Read(dat)
	return string(dat[:n]), err
}
