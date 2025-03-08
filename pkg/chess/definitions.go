package chess

type Piece uint8

const (
	Empty Piece = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
)

type Color uint8

const (
	White = iota
	Black
)

type Variant uint8

const (
	Standard = iota
)
