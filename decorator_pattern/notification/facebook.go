package main

import "fmt"

type FacebookNotifier struct {
	notifier INotifier
}

func (f *FacebookNotifier) send(message string) {
	f.notifier.send(message)
	fmt.Println("Sending Facebook notification with message:", message)
}
