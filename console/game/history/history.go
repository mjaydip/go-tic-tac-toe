package history

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

// GameHistory stores local cache of steps of game
type GameHistory [][9]string

// Push adds step to history
func (gh GameHistory) Push(step [9]string) [][9]string {
	return append(gh, step)
}

// WriteHistory writes game steps in file to story last run history
func (gh GameHistory) WriteHistory() {
	ts := time.Now()
	bs := ""
	bs += fmt.Sprintf("History written at %v\n", ts)
	for _, h := range gh {
		bs += strings.Join(h[:], " ")
		bs += "\n"
	}
	err := ioutil.WriteFile("history.txt", []byte(bs), 0777)
	if err != nil {
		log.Fatalf("Unable to write history %v", err)
	}
}
