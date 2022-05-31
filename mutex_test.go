package golanggoroutine

import (
	"sync"
	"testing"
	"time"
)

// Mutex digunakan untuk mengunci variabel saat diakses oleh goroutine
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	println(x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	balance int
}

func (account *BankAccount) addBalance(amount int) {
	account.RWMutex.Lock()
	account.balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) getBalance() int {
	account.RWMutex.RLock()
	balance := account.balance
	account.RWMutex.RUnlock()

	return balance
}

// Rw Mutex (read & write)
func TestRwMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.addBalance(1)
				println(account.getBalance())
			}
		}()
	}

	time.Sleep(7 * time.Second)
	println("Total Balance", account.getBalance())
}


// Eaxmple DeadLock
type Rekening struct {
	RWMutex sync.RWMutex
	Name    string
	Balance int
}

func Transfer(user_1 *Rekening, user_2 *Rekening, amount int) {
	user_1.Lock()
	println("Lock by", user_1.Name)
	user_1.setBalance(-amount)

	time.Sleep(1 * time.Second)

	user_2.Lock()
	println("Lock by", user_2.Name)
	user_2.setBalance(amount)

	time.Sleep(1 * time.Second)

	user_1.Unlock()
	user_2.Unlock()
}

func (rekening *Rekening) setBalance(amount int) {
	rekening.Balance += amount
}

func (rekening *Rekening) Lock() {
	rekening.RWMutex.Lock()
}

func (rekening *Rekening) Unlock() {
	rekening.RWMutex.Unlock()
}

func TestDeadLock(t *testing.T) {
	user_1 := Rekening{Name: "User 1", Balance: 10000}
	user_2 := Rekening{Name: "User 2", Balance: 10000}

	go Transfer(&user_1, &user_2, 1000)
	go Transfer(&user_2, &user_1, 2000)

	time.Sleep(2 * time.Second)

	println("Name", user_1.Name, "Total Balance", user_1.Balance)
	println("Name", user_2.Name, "Total Balance", user_2.Balance)
}
