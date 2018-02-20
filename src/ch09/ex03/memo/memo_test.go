package memo

import (
	"testing"

	"github.com/revenue-hack/go-practice/src/ch09/ex03/memo"
	"github.com/revenue-hack/go-practice/src/ch09/ex03/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}

func TestCancel(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.SequentialCancel(t, m)
	memotest.Sequential(t, m)
}

func TestConcurrentCancel(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.ConcurrentCancel(t, m)
	memotest.Concurrent(t, m)
}
