package popcount

import (
	"testing"

	"github.com/revenue-hack/go-practice/src/ch09/ex02/popcount"
)

func TestPopCount(t *testing.T) {
	ch := make(chan struct{}, 10)
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
		go func() {
			if popcount.PopCount(10) != 2 {
				t.Error("TestPopCount func popCount not 10")
			}
			<-ch
		}()
	}
}
