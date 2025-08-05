package main

// Android is the concrete factory.
type Android struct{}

func (a *Android) createSMSNotification() ISMSNotification {
	return &AndroidSMS{
		SMS: SMS{
			platform: "Android",
		},
	}
}

func (a *Android) createEmailNotification() IEmailNotification {
	return &AndroidEmail{
		Email: Email{
			platform: "Android",
		},
	}
}
