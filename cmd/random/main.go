package main

import (
	"uci-impl/internal/handler"
	"uci-impl/internal/solver/random"
)

// main program loop
func main() {
	solver := random.NewRandomSolver()
	server := handler.NewServer(solver)
	server.ServeForever()
}
