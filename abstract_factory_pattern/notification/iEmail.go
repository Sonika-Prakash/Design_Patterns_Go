package main

import "fmt"

// IEmailNotification is the abstract product interface for email notifications.
type IEmailNotification interface {
	sendEmail(message string)
	getPlatform() string
}

type Email struct {
	platform string
}

func (e *Email) sendEmail(message string) {
	fmt.Println("Sending Email:", message)
}

func (e *Email) getPlatform() string {
	return e.platform
}
