package controllers

import (
	"fmt"

	"github.com/bimalvasan/wcnp-notifier/collectors"
	"github.com/bimalvasan/wcnp-notifier/notifiers"
)

func Execute()  {
	fmt.Println("Retriving unassigned tickets from Jira.")
	collectors.GetUnassignedTickets()

	fmt.Println("Notifying unassigned ticket details in slack.")
	notifiers.SendMessageToChannel("", "Test message")
}