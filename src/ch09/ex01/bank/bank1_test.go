package bank

import (
	"fmt"
	"testing"

	"github.com/revenue-hack/go-practice/src/ch09/ex01/bank"
)

func TestWithdraw(t *testing.T) {
	bank.Deposit(200)
	is := bank.Withdraw(100)
	if !is {
		t.Errorf("got is %v\twant is true", is)
	}
	is = bank.Withdraw(100)
	if !is {
		t.Errorf("got is %v\twant is true", is)
	}
	is = bank.Withdraw(100)
	if is {
		t.Errorf("got is %v\twant is false", is)
	}
	if got := bank.Balance(); got != 0 {
		t.Errorf("Balance = %d\twant is 0", got)
	}
}

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
