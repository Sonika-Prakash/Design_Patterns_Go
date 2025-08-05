package main

// IPlatform is the abstract factory interface.
type IPlatform interface {
	createSMSNotification() ISMSNotification
	createEmailNotification() IEmailNotification
}

func createPlatform(platformType string) IPlatform {
	switch platformType {
	case "Android":
		return &Android{}
	case "iOS":
		return &IOS{}
	default:
		return nil
	}
}
