package main

// INotifier is the component interface for the notification system.
type INotifier interface {
	send(message string)
}
