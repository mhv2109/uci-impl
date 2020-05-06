package handler

import (
	"bufio"
	"os"
	"strings"

	"github.com/mhv2109/uci-impl/internal/solver"
)

// Server continuously reads from stdin and submits commands to the Handler.
type Server struct {
	handler *UCIInputHandler
}

// NewServer returns a newly initialized Server instance.
func NewServer(s solver.Solver) *Server {
	return &Server{
		handler: NewHandler(s)}
}

func (server *Server) ServeForever() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		server.process(text)
	}
}

func (server *Server) process(text string) {
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	// split input by whitespace
	input := strings.Fields(text)

	server.handler.Handle(input)
}
