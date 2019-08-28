package main

import (
	"fmt"
	"github.com/bimalvasan/wcnp-notifier/collectors"
	"github.com/bimalvasan/wcnp-notifier/notifiers"
)

func main()  {
	fmt.Println("Starting the application...")
	collectors.GetUnassignedTickets()

	notifiers.SendMessageToChannel("", "Test message")
}