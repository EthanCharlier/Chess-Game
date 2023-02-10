package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var turn string
	var start string
	var shift string
	var number_turn int

	turn = "white"

	print_turn(turn)
	default_chess := [][]string{
		{"a8", "black", "tower1"}, {"b8", "black", "knight1"}, {"c8", "black", "bishop1"}, {"d8", "black", "queen"}, {"e8", "black", "king"}, {"f8", "black", "bishop2"}, {"g8", "black", "knight2"}, {"h8", "black", "tower2"},
		{"a7", "black", "pawn1"}, {"b7", "black", "pawn2"}, {"c7", "black", "pawn3"}, {"d7", "black", "pawn4"}, {"e7", "black", "pawn5"}, {"f7", "black", "pawn6"}, {"g7", "black", "pawn7"}, {"h7", "black", "pawn8"},
		{"a6", "none", "empty"}, {"b6", "none", "empty"}, {"c6", "none", "empty"}, {"d6", "none", "empty"}, {"e6", "none", "empty"}, {"f6", "none", "empty"}, {"g6", "none", "empty"}, {"h6", "none", "empty"},
		{"a5", "none", "empty"}, {"b5", "none", "empty"}, {"c5", "none", "empty"}, {"d5", "none", "empty"}, {"e5", "none", "empty"}, {"f5", "none", "empty"}, {"g5", "none", "empty"}, {"h5", "none", "empty"},
		{"a4", "none", "empty"}, {"b4", "none", "empty"}, {"c4", "none", "empty"}, {"d4", "none", "empty"}, {"e4", "none", "empty"}, {"f4", "none", "empty"}, {"g4", "none", "empty"}, {"h4", "none", "empty"},
		{"a3", "none", "empty"}, {"b3", "none", "empty"}, {"c3", "none", "empty"}, {"d3", "none", "empty"}, {"e3", "none", "empty"}, {"f3", "none", "empty"}, {"g3", "none", "empty"}, {"h3", "none", "empty"},
		{"a2", "white", "pawn1"}, {"b2", "white", "pawn2"}, {"c2", "white", "pawn3"}, {"d2", "white", "pawn4"}, {"e2", "white", "pawn5"}, {"f2", "white", "pawn6"}, {"g2", "white", "pawn7"}, {"h2", "white", "pawn8"},
		{"a1", "white", "tower1"}, {"b1", "white", "knight1"}, {"c1", "white", "bishop1"}, {"d1", "white", "queen"}, {"e1", "white", "king"}, {"f1", "white", "bishop2"}, {"g1", "white", "knight2"}, {"h1", "white", "tower2"}}

	pawn_list := [][]string{{"white", "pawn1", "false"}, {"white", "pawn2", "false"}, {"white", "pawn3", "false"}, {"white", "pawn4", "false"}, {"white", "pawn5", "false"}, {"white", "pawn6", "false"}, {"white", "pawn7", "false"}, {"white", "pawn8", "false"},
		{"black", "pawn1", "false"}, {"black", "pawn2", "false"}, {"black", "pawn3", "false"}, {"black", "pawn4", "false"}, {"black", "pawn5", "false"}, {"black", "pawn6", "false"}, {"black", "pawn7", "false"}, {"black", "pawn8", "false"}}

	print_actual_chess(default_chess)

	for number_turn < 10 {
		fmt.Println("\n⬇ Who ⬇")
		fmt.Scan(&start)
		fmt.Println("⬇ Where ⬇")
		fmt.Scan(&shift)

		real_inputs := check_inputs(default_chess, pawn_list, start, shift, turn)
		start = real_inputs[0]
		shift = real_inputs[1]

		turn = change_turn(turn)
		print_turn(turn)
		print_actual_chess(default_chess)
		number_turn = number_turn + 1
	}
}

func print_actual_chess(default_chess [][]string) string {
	var chess_to_print string = ""
	var counter int = 1
	for _, piece := range default_chess {

		if piece[1] == "none" {
			if check_return(counter) {
				fmt.Println(" ··  ", strconv.Itoa(((counter/8)-8)*(-1)+1))
			} else {
				fmt.Print(" ··")
			}
		} else if piece[1] == "black" {
			if strings.Contains(piece[2], "pawn") {
				if check_return(counter) {
					fmt.Println(" ♙   ", strconv.Itoa(((counter/8)-8)*(-1)+1))
				} else {
					fmt.Print(" ♙ ")
				}
			}
			if strings.Contains(piece[2], "tower") {
				if check_return(counter) {
					fmt.Println(" ♖   ", strconv.Itoa(((counter/8)-8)*(-1)+1))
				} else {
					fmt.Print(" ♖ ")
				}
			}
			if strings.Contains(piece[2], "knight") {
				if check_return(counter) {
					fmt.Println(" ♘   ", strconv.Itoa(((counter/8)-8)*(-1)+1))
				} else {
					fmt.Print(" ♘ ")
				}
			}
			if strings.Contains(piece[2], "bishop") {
				if check_return(counter) {
					fmt.Println(" ♗   ", strconv.Itoa(((counter/8)-8)*(-1)+1))
				} else {
					fmt.Print(" ♗ ")
				}
			}
			if strings.Contains(piece[2], "queen") {
				if check_return(counter) {
					fmt.Println(" ♕   ", strconv.Itoa(((counter/8)-8)*(-1)+1))
				} else {
					fmt.Print(" ♕ ")
				}
			}
			if strings.Contains(piece[2], "king") {
				if check_return(counter) {
					fmt.Println(" ♔   ", strconv.Itoa(((counter/8)-8)*(-1)+1))
				} else {
					fmt.Print(" ♔ ")
				}
			}
		} else if piece[1] == "white" {
			if strings.Contains(piece[2], "pawn") {
				if check_return(counter) {
					fmt.Println(" ♟   ", strconv.Itoa(((counter/8)-8)*(-1)+1))
				} else {
					fmt.Print(" ♟ ")
				}
			}
			if strings.Contains(piece[2], "tower") {
				if check_return(counter) {
					fmt.Println(" ♜   ", strconv.Itoa(((counter/8)-8)*(-1)+1))
				} else {
					fmt.Print(" ♜ ")
				}
			}
			if strings.Contains(piece[2], "knight") {
				if check_return(counter) {
					fmt.Println(" ♞   ", strconv.Itoa(((counter/8)-8)*(-1)+1))
				} else {
					fmt.Print(" ♞ ")
				}
			}
			if strings.Contains(piece[2], "bishop") {
				if check_return(counter) {
					fmt.Println(" ♝   ", strconv.Itoa(((counter/8)-8)*(-1)+1))
				} else {
					fmt.Print(" ♝ ")
				}
			}
			if strings.Contains(piece[2], "queen") {
				if check_return(counter) {
					fmt.Println(" ♛   ", strconv.Itoa(((counter/8)-8)*(-1)+1))
				} else {
					fmt.Print(" ♛ ")
				}
			}
			if strings.Contains(piece[2], "king") {
				if check_return(counter) {
					fmt.Println(" ♚   ", strconv.Itoa(((counter/8)-8)*(-1)+1))
				} else {
					fmt.Print(" ♚ ")
				}
			}
		}
		counter += 1
	}
	fmt.Println("\n a  b  c  d  e  f  g  h")
	return chess_to_print
}

func check_return(counter int) bool {
	var return_print = []int{8, 16, 24, 32, 40, 48, 56, 64}
	for _, value := range return_print {
		if counter == value {
			return true
		}
	}
	return false
}

func check_inputs(default_chess [][]string, pawn_list [][]string, start string, shift string, turn string) []string {
	var inputs []bool
	var inputs_values []string
	var start_input bool = false
	var shift_input bool = false

	var letter_input = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	var number_input = []string{"1", "2", "3", "4", "5", "6", "7", "8"}

	if start_input != true {
		if len(start) == 2 {
			for _, letter := range letter_input {
				for _, number := range number_input {
					if strings.HasPrefix(start, letter) && strings.HasSuffix(start, number) {
						start_input = true
					}
				}
			}
		} else {
			fmt.Println("Who input lenght not correct.")
		}
	}
	if shift_input != true {
		if len(shift) == 2 {
			for _, letter := range letter_input {
				for _, number := range number_input {
					if strings.HasPrefix(shift, letter) && strings.HasSuffix(shift, number) {
						shift_input = true
					}
				}
			}
		} else {
			fmt.Println("Where input lenght not correct.")
		}
	}

	inputs = []bool{start_input, shift_input}
	inputs_values = []string{start, shift}

	if inputs[0] != true || inputs[1] != true {
		var other_start string
		var other_shift string

		fmt.Println("Who or Where inputs not correct.\n⬇ Who ⬇")
		fmt.Scan(&other_start)
		fmt.Println("⬇ Where ⬇")
		fmt.Scan(&other_shift)
		return check_inputs(default_chess, pawn_list, other_start, other_shift, turn)
	} else if check_turn(default_chess, pawn_list, start, shift, turn) {
		var other_start string
		var other_shift string

		fmt.Println("⬇ Who ⬇")
		fmt.Scan(&other_start)
		fmt.Println("⬇ Where ⬇")
		fmt.Scan(&other_shift)
		return check_inputs(default_chess, pawn_list, other_start, other_shift, turn)
	}
	return inputs_values
}

func print_turn(turn string) {
	if turn == "white" {
		fmt.Println(" ----- White  Turn -----")
	} else if turn == "black" {
		fmt.Println(" ----- Black  Turn -----")
	}
	return
}

func check_turn(default_chess [][]string, pawn_list [][]string, start string, shift string, turn string) bool {
	var start_color string
	var shift_color string
	var start_piece []string
	var shift_piece []string

	for _, chessPiece := range default_chess {
		if chessPiece[0] == start {
			start_piece = chessPiece
			start_color = chessPiece[1]
		} else if chessPiece[0] == shift {
			shift_piece = chessPiece
			shift_color = chessPiece[1]
		}
	}
	if start_color == turn {
		if start_color == shift_color {
			fmt.Println("Shift impossible, same color.")
			return true
		} else if shift_color == "none" {
			check_shift(default_chess, pawn_list, start_piece, shift_piece)
			return false
		} else {
			check_offensive()
			return false
		}
	}
	fmt.Println("Not the right color.")
	return true
}

func check_double_shift(pawn_list [][]string, pawn_find []string) bool {
	for i, pawn := range pawn_list {
		if pawn[0] == pawn_find[1] && pawn[1] == pawn_find[2] {
			if pawn[2] == "false" {
				pawn_list[i][2] = "true"
				return true
			}
		}
	}
	return false
}

func check_in(default_chess [][]string, shift string) bool {
	for _, position := range default_chess {
		if shift == position[0] && position[2] == "empty" {
			return true
		}
	}
	return false
}

func check_shift(default_chess [][]string, pawn_list [][]string, start_piece []string, shift_piece []string) [][]string {
	var shift bool = false
	letters := []string{"/", "/", "a", "b", "c", "d", "e", "f", "g", "h", "/", "/"}
	var indice_letter_start_piece int
	var indice_letter_shift_piece int
	var shift_list []string

	indice_number_start_piece, err := strconv.Atoi(string(start_piece[0][1]))
	indice_number_shift_piece, err := strconv.Atoi(string(shift_piece[0][1]))

	for i, letter := range letters {
		if string(start_piece[0][0]) == letter {
			indice_letter_start_piece = i
		}
		if string(shift_piece[0][0]) == letter {
			indice_letter_shift_piece = i
		}
	}
	fmt.Println(indice_letter_start_piece, indice_letter_shift_piece)

	if err != nil {
		fmt.Println(err)
	} else {
		if strings.Contains(start_piece[2], "pawn") {
			if start_piece[1] == "white" {
				if letters[indice_letter_shift_piece]+strconv.Itoa(indice_number_shift_piece) == letters[indice_letter_start_piece]+strconv.Itoa(indice_number_start_piece+2) {
					if check_double_shift(pawn_list, start_piece) == true {
						move(default_chess, start_piece, shift_piece)
					} else if check_double_shift(pawn_list, start_piece) == false {
						fmt.Println("Impossible.")
					}
				} else if letters[indice_letter_shift_piece]+strconv.Itoa(indice_number_shift_piece) == letters[indice_letter_start_piece]+strconv.Itoa(indice_number_start_piece+1) {
					move(default_chess, start_piece, shift_piece)
				}

			} else if start_piece[1] == "black" {
				if letters[indice_letter_shift_piece]+strconv.Itoa(indice_number_shift_piece) == letters[indice_letter_start_piece]+strconv.Itoa(indice_number_start_piece-2) {
					if check_double_shift(pawn_list, start_piece) == true {
						move(default_chess, start_piece, shift_piece)
						shift = true
					} else if check_double_shift(pawn_list, start_piece) == false {
						fmt.Println("Impossible.")
					}
				} else if letters[indice_letter_shift_piece]+strconv.Itoa(indice_number_shift_piece) == letters[indice_letter_start_piece]+strconv.Itoa(indice_number_start_piece-1) {
					move(default_chess, start_piece, shift_piece)
					shift = true
				} else {
					fmt.Println("err") ////////////////////
				}
			}

		} else if strings.Contains(start_piece[2], "knight") {
			if check_in(default_chess, letters[indice_letter_start_piece-1]+strconv.Itoa(indice_number_start_piece+2)) { //HG
				shift_list = append(shift_list, letters[indice_letter_start_piece-1]+strconv.Itoa(indice_number_start_piece+2))
			}
			if check_in(default_chess, letters[indice_letter_start_piece+1]+strconv.Itoa(indice_number_start_piece+2)) { //HD
				shift_list = append(shift_list, letters[indice_letter_start_piece+1]+strconv.Itoa(indice_number_start_piece+2))
			}
			if check_in(default_chess, letters[indice_letter_start_piece+2]+strconv.Itoa(indice_number_start_piece+1)) { //DH
				shift_list = append(shift_list, letters[indice_letter_start_piece+2]+strconv.Itoa(indice_number_start_piece+1))
			}
			if check_in(default_chess, letters[indice_letter_start_piece+2]+strconv.Itoa(indice_number_start_piece-1)) { //DB
				shift_list = append(shift_list, letters[indice_letter_start_piece+2]+strconv.Itoa(indice_number_start_piece-1))
			}
			if check_in(default_chess, letters[indice_letter_start_piece+1]+strconv.Itoa(indice_number_start_piece-2)) { //BD
				shift_list = append(shift_list, letters[indice_letter_start_piece+1]+strconv.Itoa(indice_number_start_piece-2))
			}
			if check_in(default_chess, letters[indice_letter_start_piece-1]+strconv.Itoa(indice_number_start_piece-2)) { //BG
				shift_list = append(shift_list, letters[indice_letter_start_piece-1]+strconv.Itoa(indice_number_start_piece-2))
			}
			if check_in(default_chess, letters[indice_letter_start_piece-2]+strconv.Itoa(indice_number_start_piece-1)) { //GB
				shift_list = append(shift_list, letters[indice_letter_start_piece-2]+strconv.Itoa(indice_number_start_piece-1))
			}
			if check_in(default_chess, letters[indice_letter_start_piece-2]+strconv.Itoa(indice_number_start_piece+1)) { //GH
				shift_list = append(shift_list, letters[indice_letter_start_piece-2]+strconv.Itoa(indice_number_start_piece+1))
			}

			for _, position := range shift_list {
				if shift_piece[0] == position {
					move(default_chess, start_piece, shift_piece)
					shift = true
				}
			}

		} else if strings.Contains(start_piece[2], "tower") {
			var test bool = false
			var counter1, counter2, counter3, counter4 int = 1, 1, 1, 1
			for test == false {
				test_check, position, check := check_shift_linear(default_chess, counter1, counter2, counter3, counter4, letters, indice_letter_start_piece, indice_number_start_piece)
				if test_check {
					if check == 1 {
						shift_list = append(shift_list, position)
						counter1 += 1
						test = false
					} else if check == 2 {
						shift_list = append(shift_list, position)
						counter2 += 1
						test = false
					} else if check == 3 {
						shift_list = append(shift_list, position)
						counter3 += 1
						test = false
					} else if check == 4 {
						shift_list = append(shift_list, position)
						counter4 += 1
						test = false
					}
				} else {
					test = true
				}
			}

			for _, position := range shift_list {
				if shift_list[0] == position {
					move(default_chess, start_piece, shift_piece)
					shift = true
				}
			}

		} else if strings.Contains(start_piece[2], "bishop") {
			var test bool = false
			var counter1, counter2, counter3, counter4 int = 1, 1, 1, 1
			for test == false {
				test_check, position, check := check_shift_diagonal(default_chess, counter1, counter2, counter3, counter4, letters, indice_letter_start_piece, indice_number_start_piece)
				if test_check {
					if check == 1 {
						shift_list = append(shift_list, position)
						counter1 += 1
						test = false
					} else if check == 2 {
						shift_list = append(shift_list, position)
						counter2 += 1
						test = false
					} else if check == 3 {
						shift_list = append(shift_list, position)
						counter3 += 1
						test = false
					} else if check == 4 {
						shift_list = append(shift_list, position)
						counter4 += 1
						test = false
					}
				} else {
					test = true
				}
			}

			for _, position := range shift_list {
				if shift_list[0] == position {
					move(default_chess, start_piece, shift_piece)
					shift = true
				}
			}

		} else if strings.Contains(start_piece[2], "king") {
			if check_in(default_chess, letters[indice_letter_start_piece]+strconv.Itoa(indice_number_start_piece+1)) { //H
				shift_list = append(shift_list, letters[indice_letter_start_piece]+strconv.Itoa(indice_number_start_piece+1))
			}
			if check_in(default_chess, letters[indice_letter_start_piece+1]+strconv.Itoa(indice_number_start_piece)) { //D
				shift_list = append(shift_list, letters[indice_letter_start_piece+1]+strconv.Itoa(indice_number_start_piece))
			}
			if check_in(default_chess, letters[indice_letter_start_piece]+strconv.Itoa(indice_number_start_piece-1)) { //B
				shift_list = append(shift_list, letters[indice_letter_start_piece]+strconv.Itoa(indice_number_start_piece-1))
			}
			if check_in(default_chess, letters[indice_letter_start_piece-1]+strconv.Itoa(indice_number_start_piece)) { //G
				shift_list = append(shift_list, letters[indice_letter_start_piece-1]+strconv.Itoa(indice_number_start_piece))
			}

			for _, position := range shift_list {
				if shift_piece[0] == position {
					move(default_chess, start_piece, shift_piece)
					shift = true
				}
			}

		} else if strings.Contains(start_piece[2], "queen") {
			var test bool = false
			var counter1, counter2, counter3, counter4 int = 1, 1, 1, 1
			for test == false {
				test_check, position, check := check_shift_diagonal(default_chess, counter1, counter2, counter3, counter4, letters, indice_letter_start_piece, indice_number_start_piece)
				if test_check {
					if check == 1 {
						shift_list = append(shift_list, position)
						counter1 += 1
						test = false
					} else if check == 2 {
						shift_list = append(shift_list, position)
						counter2 += 1
						test = false
					} else if check == 3 {
						shift_list = append(shift_list, position)
						counter3 += 1
						test = false
					} else if check == 4 {
						shift_list = append(shift_list, position)
						counter4 += 1
						test = false
					}
				} else {
					test = true
				}
			}

			test = false
			counter1, counter2, counter3, counter4 = 1, 1, 1, 1
			for test == false {
				test_check, position, check := check_shift_linear(default_chess, counter1, counter2, counter3, counter4, letters, indice_letter_start_piece, indice_number_start_piece)
				if test_check {
					if check == 1 {
						shift_list = append(shift_list, position)
						counter1 += 1
						test = false
					} else if check == 2 {
						shift_list = append(shift_list, position)
						counter2 += 1
						test = false
					} else if check == 3 {
						shift_list = append(shift_list, position)
						counter3 += 1
						test = false
					} else if check == 4 {
						shift_list = append(shift_list, position)
						counter4 += 1
						test = false
					}
				} else {
					test = true
				}
			}

			for _, position := range shift_list {
				if shift_list[0] == position {
					move(default_chess, start_piece, shift_piece)
					shift = true
				}
			}

		}
	}
	if shift == false {
		fmt.Println("err")
	}
	return pawn_list
}

func check_shift_linear(default_chess [][]string, counter1, counter2, counter3, counter4 int, letters []string, indice_letter_start_piece int, indice_number_start_piece int) (bool, string, int) {
	for _, piece := range default_chess {
		if piece[0] == letters[indice_letter_start_piece]+strconv.Itoa(indice_number_start_piece+counter1) && piece[2] == "empty" { //H
			return true, letters[indice_letter_start_piece] + strconv.Itoa(indice_number_start_piece+counter1), 1
		}
		if piece[0] == letters[indice_letter_start_piece+counter2]+strconv.Itoa(indice_number_start_piece) && piece[2] == "empty" { //D
			return true, letters[indice_letter_start_piece+counter2] + strconv.Itoa(indice_number_start_piece), 2
		}
		if piece[0] == letters[indice_letter_start_piece]+strconv.Itoa(indice_number_start_piece-counter3) && piece[2] == "empty" { //B
			return true, letters[indice_letter_start_piece] + strconv.Itoa(indice_number_start_piece-counter3), 3
		}
		if piece[0] == letters[indice_letter_start_piece-counter4]+strconv.Itoa(indice_number_start_piece) && piece[2] == "empty" { //G
			return true, letters[indice_letter_start_piece-counter4] + strconv.Itoa(indice_number_start_piece), 4
		}
	}
	return false, "", 0
}

func check_shift_diagonal(default_chess [][]string, counter1, counter2, counter3, counter4 int, letters []string, indice_letter_start_piece int, indice_number_start_piece int) (bool, string, int) {
	for _, piece := range default_chess {
		if piece[0] == letters[indice_letter_start_piece+counter1]+strconv.Itoa(indice_number_start_piece+counter1) && piece[2] == "empty" { //HD
			return true, letters[indice_letter_start_piece+counter1] + strconv.Itoa(indice_number_start_piece+counter1), 1
		}
		if piece[0] == letters[indice_letter_start_piece-counter2]+strconv.Itoa(indice_number_start_piece+counter2) && piece[2] == "empty" { //HG
			return true, letters[indice_letter_start_piece-counter2] + strconv.Itoa(indice_number_start_piece+counter2), 2
		}
		if piece[0] == letters[indice_letter_start_piece+counter3]+strconv.Itoa(indice_number_start_piece-counter3) && piece[2] == "empty" { //BD
			return true, letters[indice_letter_start_piece+counter3] + strconv.Itoa(indice_number_start_piece-counter3), 3
		}
		if piece[0] == letters[indice_letter_start_piece-counter4]+strconv.Itoa(indice_number_start_piece-counter4) && piece[2] == "empty" { //BG
			return true, letters[indice_letter_start_piece-counter4] + strconv.Itoa(indice_number_start_piece-counter4), 4
		}
	}
	return false, "", 0
}

func check_offensive() { //pb offensive, offensive droite avec pion
	fmt.Println("off")
}

func move(default_chess [][]string, start_piece []string, shift_piece []string) {
	var piece1 []string = []string{start_piece[0], start_piece[1], start_piece[2]}
	var piece2 []string = []string{shift_piece[0], shift_piece[1], shift_piece[2]}

	start_piece[1] = piece2[1]
	start_piece[2] = piece2[2]

	shift_piece[1] = piece1[1]
	shift_piece[2] = piece1[2]

	for i, piece := range default_chess {
		if piece[0] == piece1[0] && piece[1] == piece1[1] && piece[2] == piece1[2] {
			default_chess[i] = start_piece
		} else if piece[0] == piece2[0] && piece[1] == piece2[1] && piece[2] == piece2[2] {
			default_chess[i] = shift_piece
		}
	}
}

func change_turn(turn string) string {
	if turn == "white" {
		turn = "black"
	} else if turn == "black" {
		turn = "white"
	}
	return turn
}
