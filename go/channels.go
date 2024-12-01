package main

import (
	"fmt"
	"time"
)
// chan <- write to channel
// <- chan read from channel

func sendEmail(message string) {
	// go keyword is the start of a gorutine function
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

func test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

func waitForDBs(numDBs int, dbChan chan struct{}) {
	for _ = range numDBs {
		<-dbChan
	}
}

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, &count
}


func main() {
	test("Hello there Kaladin!")
	test("Hi there Shallan!")
	test("Hey there Dalinar!")
}

func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, len(emails))
	for i := 0; i < len(emails); i++ {
		ch<- emails[i]
	}
	return ch
}

func concurrentFib(n int) []int {
	ch := make(chan int)
	s := []int{}
	go fibonacci(n, ch)
	for i := range ch {
		s = append(s, i)
	}
	return s
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func logMessages(chEmails, chSms chan string) {
	for {
		select {
		case e, ok := <- chEmails:
			if !ok {
				return
			}
			logEmail(e)
		case s, ok := <- chSms:
			if !ok {
				return
			}
			logSms(s)
		}
	}
}

func logSms(sms string) {
	fmt.Println("SMS:", sms)
}

func logEmail(email string) {
	fmt.Println("Email:", email)
}

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	select {
	case <-snapshotTicker:
		takeSnapshot(logChan)
	case <-saveAfter:
		saveSnapshot(logChan)
		return
	default:
		waitForData(logChan)
		time.Sleep(500 * time.Millisecond)
	}
}

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}

//  A declared but uninitialized channel is nil just like a slice
//
// var s []int       // s is nil
// var c chan string // c is nil
//
// var s = make([]int, 5) // s is initialized and not nil
// var c = make(chan int) // c is initialized and not nil
//
// A send to a nil channel blocks forever
//
// A receive from a nil channel blocks forever
//
// var c chan string // c is nil
// fmt.Println(<-c)  // blocks
//
// A send to a closed channel panics
//
// var c = make(chan int, 100)
// close(c)
// c <- 1 // panic: send on closed channel
//
// A receive from a closed channel returns the zero value immediately
//
// var c = make(chan int, 100)
// close(c)
// fmt.Println(<-c) // 0
//

// func processMessages(messages []string) []string {
// 	ch := make(chan string)
// 	s := []string{}
// 	for _, m := range messages {
// 		go func() {
// 			ch<- process(m)
// 		}()
// 		s = append(s, <-ch)
// 	}
// 	return s
// }

func processMessages(messages []string) []string {
	if len(messages) == 0 {
		return []string{}
	}

	ch := make(chan string, len(messages))

	// go func(m string) m is the signature of the func() and msg will be the variable wich is put to m
	for _, msg := range messages {
		go func(m string) {
			ch <- process(m)
		}(msg)
	}

	processedMessages := make([]string, len(messages))
	for i := 0; i < len(messages); i++ {
		processedMessages[i] = <-ch
	}

	return processedMessages
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}

