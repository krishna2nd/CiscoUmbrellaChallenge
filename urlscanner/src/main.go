package main

import (
	"userver"
)

// Creating a server and passing a channel
// It can be achived via sync also but I felt better via channels
// if multi thread server is starting via sharing same channel can used for
// gracefull shutdown of server.
func main() {
	ch := make(chan int)
	server := userver.New(ch)
	go server.Start()
	<-ch
}
