package minimax

import (
	"log"

	lru "github.com/hashicorp/golang-lru"
	"github.com/mhv2109/uci-impl/internal/solver/utils"
)

type cacheValue struct {
	Score utils.CentiPawns
	Depth int
}

type cacheWrapper struct {
	cache *lru.ARCCache
}

func newCacheWrapper(size int) *cacheWrapper {
	arcCache, err := lru.NewARC(size)
	if err != nil {
		log.Panicf("Error initializing cache: %s", err)
	}
	return &cacheWrapper{arcCache}
}

func (wrapper *cacheWrapper) Get(FEN string, depth int) (score utils.CentiPawns, ok bool) {
	if value, ok := wrapper.cache.Get(FEN); ok {
		if v := value.(cacheValue); v.Depth <= depth {
			return v.Score, true
		}
	}
	return 0, false
}

func (wrapper *cacheWrapper) Add(FEN string, depth int, score utils.CentiPawns) {
	value := cacheValue{score, depth}
	wrapper.cache.Add(FEN, value)
}
