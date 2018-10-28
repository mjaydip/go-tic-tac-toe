package game

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandleMove handle the move from client
func HandleMove(w http.ResponseWriter, req *http.Request) {
	mv := req.FormValue("move")
	fmt.Println("move " + mv)
	res := struct {
		Success bool `json:"success"`
	}{
		Success: true,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
