package main

// IOS is a concrete factory.
type IOS struct{}

func (i *IOS) createSMSNotification() ISMSNotification {
	return &IOSSMS{
		SMS: SMS{
			platform: "iOS",
		},
	}
}

func (i *IOS) createEmailNotification() IEmailNotification {
	return &IOSEmail{
		Email: Email{
			platform: "iOS",
		},
	}
}
