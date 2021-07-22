package main

import (
	"fmt"
	"sync"
)

func countFrequency(inputChannel <- chan string, outputChannel chan <- int, wg *sync.WaitGroup){

	defer wg.Done()

	for inputString := range inputChannel{
		for i:=0;i<len(inputString);i++{
			asciiValues := int(inputString[i]) -97
			outputChannel <- asciiValues
		}
	}

	//close(outputChannel)

}

func main(){

	frequencyStore := make(map[int] int)
	var noOfStrings int
	fmt.Println("Enter the number of strings you have: ")
	fmt.Scan(&noOfStrings)
	input := make([]string, noOfStrings)
	fmt.Println("Enter the strings: ")
	for i:=0;i<noOfStrings;i++{
		fmt.Scan(&input[i])
	}

	inputChannel := make(chan string, noOfStrings)
	outputChannel := make(chan int, 100*noOfStrings)

	var wg sync.WaitGroup

	for i:=0;i<10;i++{
		wg.Add(1)
		go countFrequency(inputChannel,outputChannel,&wg)
	}

	for _,values:= range input{
		inputChannel <- values
	}
	close(inputChannel)
	wg.Wait()
	close(outputChannel)

	for {
		value,indicator:= <-outputChannel
		if indicator == false{
			break
		}
		frequencyStore[value]++
	}

	finalResult := make([]int,26)
	fmt.Println("{")
	for key,value := range frequencyStore{
		finalResult[key] = value
	}

	for i:=0;i<26;i++{
		fmt.Println(string(i+97), ": ",finalResult[i],",")
	}
	fmt.Println("}")


}