package main

import (
	"fmt"
	"sync"
)

func addMoney(mutex *sync.Mutex, money int, tracker *int, wg *sync.WaitGroup){

	defer wg.Done()
	mutex.Lock()
	*tracker += money
	mutex.Unlock()
}

func withdrawMoney(mutex *sync.Mutex, money int, tracker *int, wg *sync.WaitGroup){
	defer wg.Done()
	mutex.Lock()
	*tracker -= money
	mutex.Unlock()
}

func main(){
	var tracker int = 0
	var balance int = 500
	var mutex = &sync.Mutex{}

	var wg sync.WaitGroup

	wg.Add(2)
	go addMoney(mutex, 150, &tracker, &wg)
	go withdrawMoney(mutex, 100, &tracker, &wg)

	wg.Wait()
	mutex.Lock()
	if tracker + balance < 0{
		fmt.Println("Your balance is not sufficient")
		balance -= tracker
	}
	balance+=tracker
	fmt.Println("Your current balance is :", balance)
	mutex.Unlock()

}