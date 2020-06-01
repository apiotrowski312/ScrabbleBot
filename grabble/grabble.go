package grabble

type grabble struct {
	board         Board
	players       []Player
	bag           Bag
	lettterPoints LettersPoint
}

// TODO: Place word

// TODO: Exception when all letters were used (plus 50 points)
