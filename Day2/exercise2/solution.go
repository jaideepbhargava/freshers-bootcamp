package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
func giveRating(wg *sync.WaitGroup, average *int) {
	defer wg.Done()
	rand.Seed( time.Now().UnixNano())
	duration := rand.Intn(100)
	time.Sleep(time.Millisecond*100*time.Duration(duration))
	*average += rand.Intn(100)
}

func main(){
	var wg sync.WaitGroup

	var average int

	for studentNumber:=0; studentNumber<200;studentNumber++{
		wg.Add(1)
		go giveRating(&wg,&average)
	}
	wg.Wait()
	fmt.Println(float64(average)/float64(200))
}