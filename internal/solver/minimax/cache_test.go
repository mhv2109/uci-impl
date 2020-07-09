package minimax

import (
	"flag"
	"math/rand"
	"os"
	"runtime"
	"runtime/debug"
	"testing"
)

var benchmarkGetCacheSize *bool

func TestMain(m *testing.M) {
	benchmarkGetCacheSize = flag.Bool("benchmarkGetCacheSize", false, "Runs Benchmark to test cache size")
	flag.Parse()
	os.Exit(m.Run())
}

func BenchmarkGetCacheSize(b *testing.B) {
	if benchmarkGetCacheSize == nil || !*benchmarkGetCacheSize {
		b.SkipNow()
	}

	debug.SetGCPercent(-1)

	for _, size := range []int{1, 100, 1000, 10000, 100000, 130745, 1000000, 2000000, 3000000} {
		cache := newCacheWrapper(size)

		for i := 0; i < size; i++ {
			cache.Add(randomString(56), 1, 1)
		}

		runtime.GC()

		b.Logf("Memory stats for size=%d:\n", size)
		printMemUsage(b)
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func printMemUsage(t *testing.B) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	t.Logf("Alloc = %v MiB", bToMb(m.Alloc))
	t.Logf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	t.Logf("\tSys = %v MiB", bToMb(m.Sys))
	t.Logf("\tNumGC = %v\n\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
