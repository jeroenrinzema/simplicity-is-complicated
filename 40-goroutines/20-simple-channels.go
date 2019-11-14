package main

import "fmt"

func health(status chan bool) {
	status <- true
}

func main() {
	status := make(chan bool)
	// remember unbuffered channels could cause a deadlock if health is not ran in a seperate go routine
	go health(status)

	res := <-status
	fmt.Println(res)
}
