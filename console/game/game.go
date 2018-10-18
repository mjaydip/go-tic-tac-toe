package game

import (
	"fmt"
	"strconv"

	"github.com/mjaydip/go-examples/tic-tac-toe/board"
	"github.com/mjaydip/go-examples/tic-tac-toe/game/history"
	"github.com/mjaydip/go-examples/tic-tac-toe/player"
	"github.com/mjaydip/go-examples/tic-tac-toe/utils"
)

// Game is a type to manage game
type Game struct {
	board          board.GameBoard
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
			if g.board[s] != cs {
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
	g.board[s] = p.Sign
	g.currentMove = s
	g.availableMoves--
	g.history = g.history.Push(g.board)
}

// Play will start the game
func (g *Game) Play() {
	for g.winner <= 0 && g.availableMoves > 0 {
		g.nextMove()
		g.PrintGameState()
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
		g.board[i] = strconv.Itoa(i + 1)
	}
	getPlayerDetails(g)
	g.currentMove = -1
	g.availableMoves = 9
	fmt.Printf("\nLet's play tic-tac-toe, %s and %s\n", g.players[0].Name, g.players[1].Name)
	g.PrintGameState()
}

// PrintGameState is to print game
func (g *Game) PrintGameState() {
	g.board.Print()
}
