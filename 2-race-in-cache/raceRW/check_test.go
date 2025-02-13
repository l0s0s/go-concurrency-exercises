//////////////////////////////////////////////////////////////////////
//
// DO NOT EDIT THIS PART
// Your task is to edit `main.go`
//

package raceRW

import (
	"strconv"
	"testing"
)

func TestMain(t *testing.T) {
	cache := Run()

	cacheLen := len(cache.cache)
	pagesLen := cache.pages.Len()
	if cacheLen != CacheSize {
		t.Errorf("Incorrect cache size %v", cacheLen)
	}
	if pagesLen != CacheSize {
		t.Errorf("Incorrect pages size %v", pagesLen)
	}
}

func TestLRU(t *testing.T) {
	loader := Loader{
		DB: GetMockDB(),
	}
	cache := New(&loader)

	for i := 0; i < 100; i++ {
		cache.Get("Test" + strconv.Itoa(i))
	}

	if len(cache.cache) != 100 {
		t.Errorf("cache not 100: %d", len(cache.cache))
	}
	cache.Get("Test0")
	cache.Get("Test101")
	if _, ok := cache.cache["Test0"]; !ok {
		t.Errorf("0 evicted incorrectly: %v", cache.cache)
	}

}

func BenchmarkCache(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Run()
	}
}