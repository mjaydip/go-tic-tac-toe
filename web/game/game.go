package game

import (
	"strconv"

	"github.com/mjaydip/go-examples/go-tic-tac-toe/web/board"
	"github.com/mjaydip/go-examples/go-tic-tac-toe/web/game/history"
	"github.com/mjaydip/go-examples/go-tic-tac-toe/web/player"
	"github.com/mjaydip/go-examples/go-tic-tac-toe/web/utils"
)

// Game is a type to manage game
type Game struct {
	Board          board.GameBoard
	players        [2]player.Player
	Winner         int
	CurrentPlayer  int
	currentMove    int
	availableMoves int
	history        history.GameHistory
}

func (g *Game) checkWinner() {
	rules := utils.GetRuleSet()
	cm := g.currentMove
	cs := g.players[g.CurrentPlayer].Sign

	f := false
	for _, rs := range rules[cm] {
		f = false
		for _, s := range rs {
			if g.Board[s] != cs {
				f = true
			}
		}
		if f == false {
			g.Winner = g.CurrentPlayer
			break
		}
	}
}

// NextMove is a func to set next move of user
func (g *Game) NextMove(mv int) {
	p := g.players[g.CurrentPlayer]
	g.Board[mv] = p.Sign
	g.currentMove = mv
	g.availableMoves--
	g.checkWinner()
	g.history = g.history.Push(g.Board)
	g.history.WriteHistory()
	if g.CurrentPlayer == 0 {
		g.CurrentPlayer = 1
	} else {
		g.CurrentPlayer = 0
	}
}

// InitGame to setup game
func (g *Game) InitGame() {
	for i := 0; i < 9; i++ {
		g.Board[i] = strconv.Itoa(i + 1)
	}
	g.currentMove = -1
	g.Winner = -1
	g.availableMoves = 9
	g.CurrentPlayer = 0
}
