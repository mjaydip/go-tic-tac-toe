package game

import (
	"fmt"
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
	winner         int
	currentPlayer  int
	currentMove    int
	availableMoves int
	history        history.GameHistory
}

func getPlayerDetails(g *Game) {
	for i := range g.players {
		g.players[i].Rank = i + 1
		g.players[i].GetPlayerInfo()
	}
}

func (g *Game) checkWinner() {
	rules := utils.GetRuleSet()
	cm := g.currentMove
	cs := g.players[g.currentPlayer].Sign

	f := false
	for _, rs := range rules[cm] {
		f = false
		for _, s := range rs {
			if g.Board[s] != cs {
				f = true
			}
		}
		if f == false {
			g.winner = g.currentPlayer
			fmt.Printf("\n%v is the winner, congratulations!\n\n", g.players[g.currentPlayer].Name)
			break
		}
	}
}

func (g *Game) nextMove() {
	if g.currentPlayer == 0 {
		g.currentPlayer = 1
	} else {
		g.currentPlayer = 0
	}
	p := g.players[g.currentPlayer]
	msg := fmt.Sprintf("\n%v's(%v) move: ", p.Name, p.Sign)
	step, _ := utils.ScanValue(msg)
	s, _ := strconv.Atoi(step)
	s--
	g.Board[s] = p.Sign
	g.currentMove = s
	g.availableMoves--
	g.history = g.history.Push(g.Board)
}

// Play will start the game
func (g *Game) Play() {
	for g.winner <= 0 && g.availableMoves > 0 {
		g.nextMove()
		g.checkWinner()
	}
	if g.winner == 0 && g.availableMoves == 0 {
		fmt.Printf("Nobody wins!!!\n\n")
	}
	g.history.WriteHistory()
}

// InitGame to setup game
func (g *Game) InitGame() {
	for i := 0; i < 9; i++ {
		g.Board[i] = strconv.Itoa(i + 1)
	}
	g.currentMove = -1
	g.availableMoves = 9
}
