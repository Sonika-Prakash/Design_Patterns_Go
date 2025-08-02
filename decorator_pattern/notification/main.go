package main

func main() {
	notif := &EmailNotifier{}
	slackNotif := &SlackNotifier{
		notifier: notif,
	}
	facebookAndSlackNotif := &FacebookNotifier{
		notifier: slackNotif,
	}
	facebookAndSlackNotif.send("Alert: New message received!")
}
