package main

import (
	"github.com/mhv2109/uci-impl/internal/handler"
	"github.com/mhv2109/uci-impl/internal/solver/minimax"
)

// main program
func main() {
	solver := minimax.NewMinimaxSolver()
	server := handler.NewServer(solver)
	server.ServeForever()
}
