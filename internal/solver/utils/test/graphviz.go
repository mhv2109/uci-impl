package test

import (
	"fmt"
	"log"

	graphviz "github.com/goccy/go-graphviz"
	cgraph "github.com/goccy/go-graphviz/cgraph"
	"github.com/google/uuid"
	"github.com/notnil/chess"

	"github.com/mhv2109/uci-impl/internal/solver/utils"
)

var gv *graphviz.Graphviz

func init() {
	gv = graphviz.New()
}

type Graph struct {
	fpath  string
	format graphviz.Format

	graph     *cgraph.Graph
	root      *cgraph.Node
	resetRoot func()
}

func NewGraph(fpath string, format graphviz.Format) *Graph {
	graph := &Graph{}
	graph.fpath = fpath
	graph.format = format
	graph.openMove()
	return graph
}

func (graph *Graph) SearchStartedCallback(position *chess.Position, moves ...*chess.Move) {
	id := uuid.New()
	graph.resetRoot = func() {
		graph.root = graph.getOrCreateNode(&id)
		graph.root.SetLabel(position.String())
	}
	graph.resetRoot()
}

func (graph *Graph) CurrentMoveCallback(move *chess.Move, thisID, parentID *uuid.UUID,
	depth int, score, alpha, beta utils.CentiPawns) {

	graph.openMove()

	mstr := move.String()

	this := graph.getOrCreateNode(thisID)
	this.SetLabel(mstr)

	if parentID != nil {
		parent := graph.getOrCreateNode(parentID)
		graph.createEdge(this, parent, score, alpha, beta)
	} else {
		graph.resetRoot()
		graph.createEdge(this, graph.root, score, alpha, beta)
		graph.closeMove(move)
	}

}

func (graph *Graph) getOrCreateNode(id *uuid.UUID) *cgraph.Node {
	name := id.String()
	node, err := graph.graph.Node(name)
	if node == nil || err != nil {
		node, err = graph.graph.CreateNode(name)
		if node == nil || err != nil {
			log.Panicf("Error getting or creating graph node %s", err)
		}
	}
	return node
}

func (graph *Graph) createEdge(next, prev *cgraph.Node, score, alpha, beta utils.CentiPawns) {
	edgeName := prev.Name() + "__" + next.Name()
	edge, err := graph.graph.CreateEdge(edgeName, prev, next)
	if err != nil {
		log.Panicf("Error creating edge: %s", err)
	}
	edge.SetLabel(fmt.Sprintf("score: %d, alpha: %d, beta: %d", score, alpha, beta))
}

func (graph *Graph) openMove() {
	if graph.graph == nil {
		if g, err := gv.Graph(); err != nil {
			log.Panicf("Error initializing graph: %s", err)
		} else {
			graph.graph = g
		}
	}
}

func (graph *Graph) closeMove(move *chess.Move) {
	fpath := graph.fpath + "_" + move.String() + "." + string(graph.format)
	if err := gv.RenderFilename(graph.graph, graph.format, fpath); err != nil {
		log.Fatal(err)
	}
	graph.graph = nil
}
