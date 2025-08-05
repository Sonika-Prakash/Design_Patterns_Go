package main

import "fmt"

// ISMSNotification is the abstract product interface.
type ISMSNotification interface {
	sendSMS(message string)
	getPlatform() string
}

type SMS struct {
	platform string
}

func (s *SMS) sendSMS(message string) {
	fmt.Println("Sending SMS:", message)
}

func (s *SMS) getPlatform() string {
	return s.platform
}
