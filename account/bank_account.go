package account

import "sync"

const (
	open   = iota
	closed = iota
)

// Define the Account type here.
type Account struct {
	amount int64
	status int
	mu     sync.RWMutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{amount: amount, status: open}
}

func (a *Account) Balance() (int64, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if a.status == closed {
		return 0, false
	}
	return a.amount, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.status == closed {
		return a.amount, false
	}
	if a.amount+amount < 0 {
		return a.amount, false
	}
	a.amount += amount
	return a.amount, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.status == closed{
		return 0, false
	}
	
	//defer 
	a.status = closed
	//a.mu.Unlock()
	return a.amount, true
}
