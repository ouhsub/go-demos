package template

import "fmt"

type ISMS interface {
	Send(string, int) error
}

type sms struct {
	ISMS
}

func (s *sms) Valid(content string) error {
	if len(content) > 63 {
		return fmt.Errorf("content is too long")
	}
	return nil
}

func (s *sms) Send(content string, phone int) error {
	if err := s.Valid(content); err != nil {
		return err
	}
	return s.Send(content, phone)
}

type TelecomSms struct {
	*sms
}

func NewTelecomSms() *TelecomSms {
	tel := &TelecomSms{}
	tel.sms = &sms{ISMS: tel}
	return tel
}

func (tel *TelecomSms) Send(content string, phone int) error {
	fmt.Println("send by telecom success")
	return nil
}
