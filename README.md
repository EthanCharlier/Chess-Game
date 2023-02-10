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
First I imagine how to order the information in a list, and I concluded that I could make a list in 2D like that : <i><b>{{Position, Color, Piece}, ...}</b></i>
  * <b>Position</b> -> String whose first character is a letter and the second a number.
  * <b>Color</b> -> String containing the piece color and none otherwise.
  * <b>Piece</b> -> String containing the piece.

I don't know at all if this is the right way or not, maybe I could have made a discionary like this : <i><b>{Position: (Color, Piece), ...}</b></i>, because indeed the position is not supposed to vary.
