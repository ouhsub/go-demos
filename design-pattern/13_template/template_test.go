package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sms_send(t *testing.T) {
	tel := NewTelecomSms()
	err := tel.Send("test", 9999999)
	assert.NoError(t, err)
}
