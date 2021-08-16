package flyweight

type color string

const (
	RED   color = "红"
	BLACK color = "黑"
)

type role string

const (
	CHE   role = "车"
	MA    role = "马"
	XIANG role = "象"
	SHI   role = "士"
	SHUAI role = "帥"
	PAO   role = "炮"
	BING  role = "兵"
)

// Chess contains the attributes by a chess unit
type Chess struct {
	id    int
	role  role
	color color
}

var chessFactory *ChessFactory

// NewChess ...
func NewChess(id int, role role, color color) *Chess {
	return &Chess{id, role, color}
}

// ChessFactory define the factory function of chess
type ChessFactory struct {
	chesses map[int]*Chess
}

// GetChessFactory init and return the ChessFactory pointer
func GetChessFactory() *ChessFactory {
	if chessFactory == nil {
		chessFactory = &ChessFactory{chesses: make(map[int]*Chess)}
		chessFactory.initalize()
	}
	return chessFactory
}

func (cf *ChessFactory) initalize() {
	cf.chesses[1] = NewChess(1, CHE, BLACK)
	cf.chesses[2] = NewChess(2, MA, BLACK)
	cf.chesses[3] = NewChess(3, XIANG, BLACK)
	cf.chesses[4] = NewChess(4, SHI, BLACK)
	cf.chesses[5] = NewChess(5, SHUAI, BLACK)
	cf.chesses[6] = NewChess(6, PAO, BLACK)
	cf.chesses[7] = NewChess(7, BING, BLACK)
	cf.chesses[8] = NewChess(8, CHE, RED)
	cf.chesses[9] = NewChess(9, MA, RED)
	cf.chesses[10] = NewChess(10, XIANG, RED)
	cf.chesses[11] = NewChess(11, SHI, RED)
	cf.chesses[12] = NewChess(12, SHUAI, RED)
	cf.chesses[13] = NewChess(13, PAO, RED)
	cf.chesses[14] = NewChess(14, BING, RED)
}

// Get return the Chess by given id
// if not in cache then create and return
func (cf *ChessFactory) Get(id int) *Chess {
	return cf.chesses[id]
}

// ChessUnit restore a Chess pointer and it's position on 2d
type ChessUnit struct {
	chess *Chess
	posx  int
	posy  int
}

// ChessBoard restore a 9*10 board
type ChessBoard struct {
	board [9][10]*Chess
}

// NewChessBoard return a initial board
func NewChessBoard() *ChessBoard {
	return &ChessBoard{}
}

// Move moves the chess from a positon to another position on the board
func (b *ChessBoard) Move(unit *ChessUnit, posx, posy int) {
	b.board[unit.posx][unit.posy] = nil
	unit.posx = posx
	unit.posy = posy
	b.board[posx][posy] = unit.chess
}
