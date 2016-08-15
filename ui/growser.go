/*
TODO:
Create templates:
Index with ids,
get set player data
templates with forms
add javascript chessboard
*/
package main

import (
	"fmt"
	"github.com/polypmer/ghess"
	"net/http"
)

type ChessHandler struct {
	g *Board
}

// boardHandler for playing game
// Takes url param pgn move
func (h *ChessHandler) playGameHandler(w http.ResponseWriter,
	r *http.Request) {
	move := r.URL.Path[len("/play/"):]
	e := h.g.ParseMove(move)
	if e != nil {
		fmt.Fprintln(w, e.Error())
	}
	fmt.Fprintln(w, h.g.String())
}

func (h *ChessHandler) newGameHandler(w http.ResponseWriter,
	r *http.Request) {
	h.g = ghess.NewBoard()
	fmt.Fprintln(w, h.g.String())
}

func (h *ChessHandler) showGameHandler(w http.ResponseWriter,
	r *http.Request) {
	//print board
}

func main() {
	// So HandlFunc takes a custom Handler
	// Which is forcement takes into a reader and writer
	// and then it will print whatever is written to the
	// writer
	h := new(ChessHandler)

	// Server Part
	http.HandleFunc("/play/", h.playGameHandler)
	http.HandleFunc("/board/", h.showGameHandler)
	http.HandleFunc("/new/", h.newGameHandler)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
