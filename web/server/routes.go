package server

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// HandleMove handle the move from client
func HandleMove(w http.ResponseWriter, req *http.Request) {
	mv := req.FormValue("move")
	cp := g.CurrentPlayer
	m, _ := strconv.Atoi(mv)
	m--
	g.NextMove(m)
	res := struct {
		Success       bool `json:"success"`
		Winner        int  `json:"winner"`
		CurrentPlayer int  `json:"currentPlayer"`
	}{
		Success:       true,
		Winner:        g.Winner,
		CurrentPlayer: cp,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
