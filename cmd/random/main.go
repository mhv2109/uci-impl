package main

import (
	"github.com/mhv2109/uci-impl/internal/handler"
	"github.com/mhv2109/uci-impl/internal/solver/random"
)

// main program loop
func main() {
	solver := random.NewRandomSolver()
	server := handler.NewServer(solver)
	server.ServeForever()
}
