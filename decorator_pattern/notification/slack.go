package main

import "fmt"

// SlackNotifier is a concrete decorator for the Notification system.
type SlackNotifier struct {
	notifier INotifier
}

func (s *SlackNotifier) send(message string) {
	s.notifier.send(message)
	fmt.Println("Sending Slack notification with message:", message)
}
