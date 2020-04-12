package solver

type SearchParams struct {
	Ponder    bool
	Wtime     int
	Btime     int
	Winc      int
	Binc      int
	Movestogo int
	Depth     int
	Nodes     int
	Mate      int
	Movetime  int
	Infinite  bool
}

func NewSearchParams() *SearchParams {
	return &SearchParams{
		false, -1, -1, -1, -1, -1,
		-1, -1, -1, -1, false}
}

type Solver interface {
	GetOption(string) *string
	SetOption(string, string)
	GetOptions() []*Option
	SetPosition(string, ...string)
	SetStartPosition(...string)
	DoMove(string)
	StartSearch(*SearchParams, ...string) chan []string
	StopSearch()
	PonderHit()
}
