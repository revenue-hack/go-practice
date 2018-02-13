package main

import "log"

func main() {
	log.Println("start")

	ch := make(chan int, 3)
	for i := 0; i < 3; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}
	// これだと出力されない
	// おそらくchを待たずに終了してしまうから
	/*
		if len(ch) == 3 {
			for i := 0; i < len(ch); i++ {
				log.Println(<-ch)
			}
		}
	*/
	// これだとうまくいく
	// 上記との違いはforでchが来るまで待っているところ
	for {
		if len(ch) == 3 {
			for i := 0; i < 3; i++ {
				log.Println(<-ch)
			}
			return
		}
	}
}
