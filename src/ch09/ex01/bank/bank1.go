package bank

var deposits = make(chan int)
var balances = make(chan int)
var withdraw = make(chan *WithdrawType)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case wd := <-withdraw:
			if balance >= wd.amount {
				balance -= wd.amount
				wd.ch <- true
			} else {
				wd.ch <- false
			}

		}
	}
}

type WithdrawType struct {
	amount int
	ch     chan<- bool
}

func Withdraw(amount int) bool {
	ch := make(chan bool)
	var t = WithdrawType{amount, ch}
	withdraw <- &t
	return <-ch
}

func init() {
	go teller()
}
