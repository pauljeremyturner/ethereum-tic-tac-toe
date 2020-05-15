package web

/*

001 | 002 | 004
---------------
008 | 016 | 032
---------------
064 | 128 | 256

Use bit shifting to determine squares for circle and cross

e.g. if squares (0,0) bottom left and (1, 0) bottom middle and (2, 0) bottom right
are occupied by circles, the it value for circle would be 64 + 128 + 256 = 448

*/
// Game request model
// swagger:request gameState
type gameState struct {
	GameId       string `json:"GameId"`
	CirclePlayer string `json:"circle"`
	CrossPlayer  string `json:"cross"`
	Circles      int `json:"Circles"`
	Crosses      int `json:"Crosses"`
}

// Game response model
// swagger:request gameMove
type gameMove struct {
	gameId string `json:"GameId"`
	player Side
	move   byte `json:"move"`
}

// swagger:request side
type Side bool

const (
	Circle Side = true
	Cross  Side = false
)
