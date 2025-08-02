package main

import "fmt"

// EmailNotifier is the concrete component for the notification system.
// This is because we need email notifications to always be sent making it a base behavior.
type EmailNotifier struct{}

func (e *EmailNotifier) send(message string) {
	fmt.Println("Sending email notification with message:", message)
}
