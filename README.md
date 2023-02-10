# Chess-Game
## About The Project

I challenged myself to reproduce a chess game in Go, I didn't know this language before starting this project. So there will be a lot of rookie error. So if you see any errors or have any suggestions please let me know.

Here is my plan :
* Imagine how to contain the chessboard.
* Imagine how to display the chessboard.
* Imagine how the pieces will move.
* Imagine an attack system.
* And finally imagine a condition to win.

<br>
<br>

### How to contain the chessboard
First I imagine how to order the information in a list, and I concluded that I could make a list in 2D like that : `{{Position, Color, Piece}, ...}`
  * <b>Position</b> -> String whose first character is a letter and the second a number.
  * <b>Color</b> -> String containing the piece color and none otherwise.
  * <b>Piece</b> -> String containing the piece.


Here is the definition of my chessboard :
```sh
  default_chess := [][]string{
		{"a8", "black", "tower1"}, {"b8", "black", "knight1"}, {"c8", "black", "bishop1"}, {"d8", "black", "queen"}, {"e8", "black", "king"}, {"f8", "black", "bishop2"}, {"g8", "black", "knight2"}, {"h8", "black", "tower2"},
		{"a7", "black", "pawn1"}, {"b7", "black", "pawn2"}, {"c7", "black", "pawn3"}, {"d7", "black", "pawn4"}, {"e7", "black", "pawn5"}, {"f7", "black", "pawn6"}, {"g7", "black", "pawn7"}, {"h7", "black", "pawn8"},
		{"a6", "none", "empty"}, {"b6", "none", "empty"}, {"c6", "none", "empty"}, {"d6", "none", "empty"}, {"e6", "none", "empty"}, {"f6", "none", "empty"}, {"g6", "none", "empty"}, {"h6", "none", "empty"},
		{"a5", "none", "empty"}, {"b5", "none", "empty"}, {"c5", "none", "empty"}, {"d5", "none", "empty"}, {"e5", "none", "empty"}, {"f5", "none", "empty"}, {"g5", "none", "empty"}, {"h5", "none", "empty"},
		{"a4", "none", "empty"}, {"b4", "none", "empty"}, {"c4", "none", "empty"}, {"d4", "none", "empty"}, {"e4", "none", "empty"}, {"f4", "none", "empty"}, {"g4", "none", "empty"}, {"h4", "none", "empty"},
		{"a3", "none", "empty"}, {"b3", "none", "empty"}, {"c3", "none", "empty"}, {"d3", "none", "empty"}, {"e3", "none", "empty"}, {"f3", "none", "empty"}, {"g3", "none", "empty"}, {"h3", "none", "empty"},
		{"a2", "white", "pawn1"}, {"b2", "white", "pawn2"}, {"c2", "white", "pawn3"}, {"d2", "white", "pawn4"}, {"e2", "white", "pawn5"}, {"f2", "white", "pawn6"}, {"g2", "white", "pawn7"}, {"h2", "white", "pawn8"},
		{"a1", "white", "tower1"}, {"b1", "white", "knight1"}, {"c1", "white", "bishop1"}, {"d1", "white", "queen"}, {"e1", "white", "king"}, {"f1", "white", "bishop2"}, {"g1", "white", "knight2"}, {"h1", "white", "tower2"}}
 ```
I don't know at all if this is the right way or not, maybe I could have made a discionary like this : `{Position: (Color, Piece), ...}`, because indeed the position is not supposed to vary.

<br>
<br>

### How to display the chessboard
Second I imagine how to print the chessboard, I decided to create a `counter` that will count each piece display and would make a line break when the `counter` is a multiple of 8. I also decided to display the letters for each column and the numbers for each line for simplicity.
<br>

For pawns I created `check_return` function, which contains `var return_print = []int{8, 16, 24, 32, 40, 48, 56, 64}`, that return `true` if `counter` is equal to a value of `return_print` and `true` otherwise :
```
func check_return(counter int) bool {
	var return_print = []int{8, 16, 24, 32, 40, 48, 56, 64}
	for _, value := range return_print {
		if counter == value {
			return true
		}
	}
	return false
}
```
<br>


I go through my chessboard and I test `color` and if the position is on the right then I add the number of the line : `strconv.Itoa(((counter/8)-8)*(-1)+1)` :
* If `color` is equal to `none`, I print `·`:
```
if piece[1] == "none" {
	if check_return(counter) {
		fmt.Println(" ··  ", strconv.Itoa(((counter/8)-8)*(-1)+1))
	} else {
		fmt.Print(" ··")
	}
}
```

* If `color` is equal to `white`, I look what piece it's and I print it :
```
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
```

* If `color` is equal to `black`, I do the same :
```
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
}
```
Of course I don't forget to add `1` to `counter` and print letters.
```
	counter += 1
}
fmt.Println("\n a  b  c  d  e  f  g  h")
```

