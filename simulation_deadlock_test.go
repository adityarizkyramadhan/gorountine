package gorountine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type nasabah struct {
	mutex   sync.Mutex
	name    string
	balance int
}

type nasabahInterface interface {
	lock()
	unlock()
	GetName() string
	GetBalance() int
	SetName(string)
	SetBalance(int)
}

func NewNasabah() nasabahInterface {
	return &nasabah{}
}

func (m *nasabah) lock() {
	m.mutex.Lock()
}

func (m *nasabah) unlock() {
	m.mutex.Unlock()
}

func (m *nasabah) GetName() string {
	return m.name
}

func (m *nasabah) GetBalance() int {
	return m.balance
}

func (m *nasabah) SetBalance(balance int) {
	m.balance = balance
}

func (m *nasabah) SetName(name string) {
	m.name = name
}

func TransferDeadLock(transferer nasabahInterface, receiver nasabahInterface, ammount int) {
	//lock pengirim
	transferer.lock()
	fmt.Println("Sedang lock pertama :", transferer.GetName())
	transferer.SetBalance(transferer.GetBalance() - ammount)
	time.Sleep(time.Second)
	//lock penerima
	receiver.lock()
	fmt.Println("Sedang lock kedua :", receiver.GetName())
	receiver.SetBalance(receiver.GetBalance() + ammount)

	transferer.unlock()
	receiver.unlock()
}

func TransferWithoutDeadLock(transferer nasabahInterface, receiver nasabahInterface, ammount int) {
	//lock pengirim
	transferer.lock()
	fmt.Println("Sedang lock pertama :", transferer.GetName())
	transferer.SetBalance(transferer.GetBalance() - ammount)
	time.Sleep(time.Second)
	transferer.unlock()
	//lock penerima
	receiver.lock()
	fmt.Println("Sedang lock kedua :", receiver.GetName())
	receiver.SetBalance(receiver.GetBalance() + ammount)
	time.Sleep(time.Second)
	receiver.unlock()
}

func TestTransferDeadLock(t *testing.T) {
	user1 := NewNasabah()
	user1.SetName("user1")
	user1.SetBalance(100)
	user2 := NewNasabah()
	user2.SetName("user2")
	user2.SetBalance(100)
	go TransferDeadLock(user1, user2, 10)
	go TransferDeadLock(user2, user1, 10)
	time.Sleep(time.Second * 2)
	fmt.Println("user1 :", user1.GetBalance())
	fmt.Println("user2 :", user2.GetBalance())
}

func TestTransferWithoutDeadLock(t *testing.T) {
	user1 := NewNasabah()
	user1.SetName("user1")
	user1.SetBalance(100)
	user2 := NewNasabah()
	user2.SetName("user2")
	user2.SetBalance(100)
	go TransferWithoutDeadLock(user1, user2, 10)
	go TransferWithoutDeadLock(user2, user1, 10)
	time.Sleep(time.Second * 2)
	fmt.Println("user1 :", user1.GetBalance())
	fmt.Println("user2 :", user2.GetBalance())
}
