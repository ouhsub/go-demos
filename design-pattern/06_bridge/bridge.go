package bridge

type IMsgSender interface {
	Send(string) error
}

type EmailMsgSender struct {
	emails []string
}

func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{emails: emails}
}

func (sender *EmailMsgSender) Send(msg string) error {
	return nil
}

type INotification interface {
	Notify(string) error
}

type ErrorNotification struct {
	sender IMsgSender
}

func NewErrorNotification(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{sender: sender}
}

func (notifier *ErrorNotification) Notify(msg string) error {
	return notifier.sender.Send(msg)
}
