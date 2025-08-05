package main

import "fmt"

func main() {
	platform := createPlatform("Android") // or "iOS"
	if platform == nil {
		fmt.Println("Unknown platform type")
		return
	}

	sms := platform.createSMSNotification()
	email := platform.createEmailNotification()
	fmt.Printf("Email notifier created for platform: %s\n", email.getPlatform())
	fmt.Printf("SMS notifier created for platform: %s\n", sms.getPlatform())

	sms.sendSMS("Hello via SMS!")
	email.sendEmail("Hello via Email!")
}
