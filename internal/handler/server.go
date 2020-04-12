package handler

import (
	"bufio"
	"os"
	"strings"

	"uci-impl/internal/solver"
)

type Server struct {
	debug   bool
	handler *UCIInputHandler
}

func NewServer(debug bool, s solver.Solver) *Server {
	return &Server{
		debug:   debug,
		handler: NewHandler(s)}
}

func (server *Server) ServeForever() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		go server.process(text)
	}
}

func (server *Server) process(text string) {
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	// split input by whitespace
	input := strings.Fields(text)

	server.handler.Handle(input)
}
