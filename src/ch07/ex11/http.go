package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type dollars float32

type database struct {
	products map[string]dollars
	lock     sync.RWMutex
}

func main() {
	db := database{products: map[string]dollars{"shoes": 50, "sockes": 5}}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/read", db.read)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func (db *database) list(w http.ResponseWriter, req *http.Request) {
	db.withRLockContext(func() {
		for item, price := range db.products {
			fmt.Fprintf(w, "%s: %v\n", item, price)
		}
	})
}

func (db *database) read(w http.ResponseWriter, req *http.Request) {
	db.withRLockContext(func() {
		item, isHas := getItem(w, req)
		if !isHas {
			return
		}
		price, ok := db.products[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
		}
		fmt.Fprintf(w, "item: %s\tprice: %v\n", item, price)
	})
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	db.withRLockContext(func() {
		item, isHas := getItem(w, req)
		if !isHas {
			return
		}
		price, ok := db.products[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
		}
		fmt.Fprintf(w, "%s price is %v\n", item, price)
	})
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	db.withLockContext(func() {
		item, isHas := getItem(w, req)
		if !isHas {
			return
		}
		_, ok := db.products[item]
		// when item nothing
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		db.products[item] = convertDollars(req.URL.Query().Get("price"))
		fmt.Fprintf(w, "update item: %s\tprice: %v\n", item, db.products[item])
	})
}

func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	db.withLockContext(func() {
		item, isHas := getItem(w, req)
		if !isHas {
			return
		}
		_, ok := db.products[item]
		if !ok {
			fmt.Fprintf(w, "item don't match: %q\n", item)
			return
		}
		delete(db.products, item)
		fmt.Fprintf(w, "delete item: %s\n", item)
	})
}

func (db *database) create(w http.ResponseWriter, req *http.Request) {
	db.withLockContext(func() {
		item, isHas := getItem(w, req)
		if !isHas {
			return
		}
		_, ok := db.products[item]
		if ok {
			fmt.Fprintf(w, "already exist item: %q\n", item)
			return
		}
		price, isHas := getPrice(w, req)
		if !isHas {
			return
		}
		db.products[item] = price
		fmt.Fprintf(w, "success add item: %q\tprice: %v\n", item, price)
	})
}

func getPrice(w http.ResponseWriter, req *http.Request) (dollars, bool) {
	price := req.URL.Query().Get("price")
	// when price nothing
	if price == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such price: %s\n", price)
		return 0, false
	}
	return convertDollars(price), true
}

func getItem(w http.ResponseWriter, req *http.Request) (string, bool) {
	item := req.URL.Query().Get("item")
	// when item nothing
	if item == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return "", false
	}
	return item, true
}

func convertDollars(str string) dollars {
	intPrice, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return dollars(intPrice)

}
func (db *database) withRLockContext(fn func()) {
	db.lock.RLock()
	defer db.lock.RUnlock()
	fn()
}

func (db *database) withLockContext(fn func()) {
	db.lock.Lock()
	defer db.lock.Unlock()
	fn()
}
