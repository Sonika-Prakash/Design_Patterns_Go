package main

import "fmt"

type SlackNotifier struct {
	notifier INotifier
}

func (s *SlackNotifier) send(message string) {
	s.notifier.send(message)
	fmt.Println("Sending Slack notification with message:", message)
}
