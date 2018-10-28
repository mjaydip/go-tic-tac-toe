package player

import (
	"fmt"

	"github.com/mjaydip/go-examples/go-tic-tac-toe/console/utils"
)

// Player is type to hold player information
type Player struct {
	Name string
	Rank int
	Sign string
}

// GetPlayerInfo will get player basic details like name
func (p *Player) GetPlayerInfo() {
	msg := fmt.Sprintf("\nPlease enter player %v name: ", p.Rank)
	p.Name, _ = utils.ScanValue(msg)

	msg = fmt.Sprintf("Please enter %v's sign (x/o): ", p.Name)
	p.Sign, _ = utils.ScanValue(msg)
}
