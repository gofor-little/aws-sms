package sms

// Data is the data required to build and send an SMS message.
type Data struct {
	SenderID    string `json:"senderID"`
	PhoneNumber string `json:"phoneNumber"`
	Message     string `json:"message"`
}
