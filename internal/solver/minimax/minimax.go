package minimax

import (
	"math"
	"math/rand"

	"github.com/mhv2109/uci-impl/internal/handler"
	"github.com/mhv2109/uci-impl/internal/handler/info"
	"github.com/mhv2109/uci-impl/internal/solver/utils"
	"github.com/notnil/chess"
)

type submitCallback func([]string) bool
type searchCallback func(*chess.Position, ...*chess.Move)
type moveCallback func(*chess.Move, int, utils.CentiPawns, utils.CentiPawns, utils.CentiPawns)

type minimaxAlgo struct {
	MaxDepth int
	HashSize int

	player  chess.Color
	submit  submitCallback
	emitter handler.Emitter

	cache *cacheWrapper

	searchStartedCallbacks  []searchCallback
	currentMoveCallbacks    []moveCallback
	bestMoveCallbacks       []moveCallback
	searchFinishedCallbacks []searchCallback
}

func newMinimaxAlgo(maxDepth int, hashSize int, submit func([]string) bool,
	emitter handler.Emitter) *minimaxAlgo {

	// maxDepth must be >= 1
	if maxDepth < 1 {
		maxDepth = 1
	}

	// Benchmarks calculated about 1021 entries per MB
	cache := newCacheWrapper(hashSize * 1021)

	minimax := &minimaxAlgo{
		maxDepth,
		hashSize,
		chess.NoColor,
		submit,
		emitter,
		cache,
		make([]searchCallback, 0),
		make([]moveCallback, 0, 1),
		make([]moveCallback, 0, 1),
		make([]searchCallback, 0)}

	minimax.Init()

	return minimax
}

func (minimax *minimaxAlgo) Init() {
	//minimax.AddCurrentMoveCallback(minimax.infoCurrentMove)
	minimax.AddBestMoveCallback(minimax.infoBestMove)
}

func (minimax *minimaxAlgo) AddSearchStaretedCallback(callback searchCallback) {
	minimax.searchStartedCallbacks = append(minimax.searchStartedCallbacks, callback)
}

func (minimax *minimaxAlgo) AddCurrentMoveCallback(callback moveCallback) {
	minimax.currentMoveCallbacks = append(minimax.currentMoveCallbacks, callback)
}

func (minimax *minimaxAlgo) AddBestMoveCallback(callback moveCallback) {
	minimax.bestMoveCallbacks = append(minimax.bestMoveCallbacks, callback)
}

func (minimax *minimaxAlgo) AddSearchFinishedCallback(callback searchCallback) {
	minimax.searchFinishedCallbacks = append(minimax.searchFinishedCallbacks, callback)
}

func (minimax *minimaxAlgo) executeSearchStartedCallbacks(position *chess.Position, moves ...*chess.Move) {
	executeSearchCallbacks(position, moves, minimax.searchStartedCallbacks)
}

func (minimax *minimaxAlgo) executeCurrentMoveCallbacks(move *chess.Move,
	depth int, score, alpha, beta utils.CentiPawns) {
	executeMoveCallbacks(move, depth, score, alpha, beta, minimax.currentMoveCallbacks)
}

func (minimax *minimaxAlgo) executeBestMoveCallbacks(move *chess.Move,
	depth int, score, alpha, beta utils.CentiPawns) {
	executeMoveCallbacks(move, depth, score, alpha, beta, minimax.bestMoveCallbacks)
}

func (minimax *minimaxAlgo) executeSearchFinishedCallbacks(position *chess.Position, moves ...*chess.Move) {
	executeSearchCallbacks(position, moves, minimax.searchFinishedCallbacks)
}

func executeMoveCallbacks(move *chess.Move, depth int, score,
	alpha, beta utils.CentiPawns, callbacks []moveCallback) {
	for _, callback := range callbacks {
		callback(move, depth, score, alpha, beta)
	}
}

func executeSearchCallbacks(position *chess.Position, moves []*chess.Move, callbacks []searchCallback) {
	for _, callback := range callbacks {
		callback(position, moves...)
	}
}

func (minimax *minimaxAlgo) Start(position *chess.Position, moves ...*chess.Move) {
	minimax.player = position.Turn()
	minimax.executeSearchStartedCallbacks(position, moves...)
	minimax.maxStep(position, 0, -math.MaxInt64, math.MaxInt64, moves...)
}

func (minimax *minimaxAlgo) maxStep(state *chess.Position, depth int,
	alpha, beta utils.CentiPawns, moves ...*chess.Move) utils.CentiPawns {

	var bestMove *chess.Move

	if gameWon(state) || depth >= minimax.MaxDepth {
		alpha = minimax.score(state)
	} else {
		if validMoves := getMoves(state, moves...); len(validMoves) == 0 {
			alpha = minimax.score(state)
		} else {
			for _, move := range validMoves {
				nextState := state.Update(move)

				var score utils.CentiPawns

				nextStateString := nextState.String()
				if value, ok := minimax.cache.Get(nextStateString, depth); ok {
					score = value
				} else {
					score = minimax.minStep(nextState, depth, alpha, beta)
					minimax.cache.Add(nextStateString, depth, score)
				}

				if score > alpha {
					alpha = score
					if depth <= 0 {
						minimax.submit([]string{move.String()})
						minimax.executeBestMoveCallbacks(move, depth, score, alpha, beta)
						bestMove = move
					}
				}
				minimax.executeCurrentMoveCallbacks(move, depth, score, alpha, beta)
				if alpha >= beta {
					break
				}
			}
		}
	}

	if depth == 0 {
		minimax.executeSearchFinishedCallbacks(state, bestMove)
	}

	return alpha
}

func (minimax *minimaxAlgo) minStep(state *chess.Position, depth int,
	alpha, beta utils.CentiPawns, moves ...*chess.Move) utils.CentiPawns {

	if gameWon(state) || depth >= minimax.MaxDepth {
		beta = minimax.score(state)
	} else {
		if validMoves := getMoves(state, moves...); len(validMoves) == 0 {
			beta = minimax.score(state)
		} else {
			for _, move := range validMoves {
				nextState := state.Update(move)

				score := minimax.maxStep(nextState, depth+1, alpha, beta)

				if score < beta {
					// dont submit opponent's moves!
					beta = score
				}
				minimax.executeCurrentMoveCallbacks(move, depth, score, alpha, beta)
				if alpha >= beta {
					break
				}
			}
		}
	}

	return beta
}

func (minimax *minimaxAlgo) infoCurrentMove(move *chess.Move, depth int, score, alpha, beta utils.CentiPawns) {
	i := info.Info{}
	i.SetDepth(depth)
	i.SetCurrmove(move.String())
	i.SetScore(info.CP, int(score))
	minimax.emitter.EmitInfo(i)
}

func (minimax *minimaxAlgo) infoBestMove(move *chess.Move, depth int, score, alpha, beta utils.CentiPawns) {
	i := info.Info{}
	i.SetPv([]string{move.String()})
	i.SetScore(info.CP, int(score))
	minimax.emitter.EmitInfo(i)
}

func (minimax *minimaxAlgo) score(state *chess.Position) utils.CentiPawns {
	if winner, ok := getWinner(state); ok {
		// return winning/losing score
		if winner == minimax.player {
			return utils.MaxScore
		}
		return -utils.MaxScore
	} else if state.Status() == chess.NoMethod {
		// return normal score
		board := state.Board()
		if minimax.player == chess.White {
			return utils.WhiteAdvantage(board)
		}
		return utils.BlackAdvantage(board)
	}
	// return draw score
	return 0
}

func getWinner(state *chess.Position) (chess.Color, bool) {
	if !gameWon(state) {
		return chess.NoColor, false
	}
	if turn := state.Turn(); turn == chess.Black {
		return chess.White, true
	}
	return chess.Black, true
}

func gameWon(state *chess.Position) bool {
	outcome := state.Status()
	if outcome == chess.Checkmate {
		return true
	}
	return false
}

func getMoves(state *chess.Position, moves ...*chess.Move) []*chess.Move {
	if len(moves) > 0 {
		return randomize(moves)
	}
	return randomize(state.ValidMoves())
}

func randomize(moves []*chess.Move) []*chess.Move {
	for i := range moves {
		j := rand.Intn(i + 1)
		moves[i], moves[j] = moves[j], moves[i]
	}
	return moves
}
