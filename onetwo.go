// program that reads in two text files (input1.txt and input2.txt), puts each word one at a time into a seperate channel, and write them to an output file.
// uses 4 goroutines, 3 channels, and a timer

package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)

	//create channels
	ch1 := make(chan string) //for input 1
	ch2 := make(chan string) //for input 2
	ch3 := make(chan string) //for output
	//timer set to 2 seconds
	timer := time.NewTimer(time.Second * 2)

	// create output file
	f, err := os.Create("output.txt")
	if err != nil {
		f.Close()
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	// start of goroutines

	//goroutines to read from files
	go readfile1(ch1)
	go readfile2(ch2)

	//goroutines to insert strings into output channel (ch3)
	go func() {
		defer close(ch1)
		defer close(ch2)
	Forever:
		for {
			select {
			//check for input from channel 1
			case str, ok := <-ch1:
				if !ok {
					break Forever
				}
				ch3 <- str + " "
			//check for input from channel 2
			case str, ok := <-ch2:
				if !ok {
					break Forever
				}
				ch3 <- str + " "
			//check for if timer is up
			case <-timer.C:
				ch3 <- "\n"
				timer.Reset(time.Second * 2)
			}
		}
	}()

	//goroutine to write to file
	go func() {
		defer close(ch3)
		for str := range ch3 {
			if _, err := w.WriteString(str); err != nil {
				panic(err)
			}
			//print statment for testing:
			fmt.Print(str)
		}
	}()

	//delay in program to give goroutines sufficient time to run
	time.Sleep(20 * time.Second)

	//write any buffered data to file
	w.Flush()

}

//function to read from input1.txt
func readfile1(ch chan string) {

	file, err := os.Open("input1.txt")
	defer file.Close()
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		ch <- word
		//delay of 0.2 seconds
		time.Sleep(200 * time.Millisecond)
	}

}

//function to read from input2.txt
func readfile2(ch chan string) {

	file, err := os.Open("input2.txt")
	defer file.Close()
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		ch <- word
		//delay of 0.3 seconds
		time.Sleep(300 * time.Millisecond)
	}

}
