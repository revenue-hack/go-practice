package memo

import (
	"errors"
	"fmt"
)

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

type request struct {
	key      string
	response chan<- result
	done     <-chan struct{}
}

var cancel = errors.New("cancel")

type Memo struct{ requests chan request }

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	fmt.Println("new")
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response, done}
	res := <-response
	if res.err != nil {
		select {
		case <-done:
			return nil, res.err
		default:
			return memo.Get(key, done)
		}
	}
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	fmt.Println("server start")
	cache := make(map[string]*entry)
	for req := range memo.requests {
		fmt.Println("server for")
		e := cache[req.key]
		if e == nil {
			fmt.Println("no cache")
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key, req.done)
		} else {
			select {
			case <-e.ready:
				if e.res.err != nil {
					delete(cache, req.key)
					e = nil
				}
			default:
				// nothing
			}
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string, done <-chan struct{}) {
	fmt.Println("call")
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	fmt.Println("deliver")
	<-e.ready
	fmt.Println("deliver response")
	response <- e.res
}
