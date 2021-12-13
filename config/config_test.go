package config

import (
	"bytes"
	"os"
	"testing"
)

import "github.com/stretchr/testify/assert"

func Test_Read_Config_Stream(t *testing.T) {
	var buffer bytes.Buffer
	payload := `
{
	"host": {
		"hostName": "mail.example.com",
		"port": 587
	},
	"credential": {
		"userName": "username@example.com",
		"password": "PassW0rd"
	},
	"sender": {
		"address": "username@example.com"
	},
	"receivers": [
		{"address": "receiver_01@example.com"},
		{"address": "receiver_02@example.com"},
		{"address": "receiver_03@example.com"}
	]
}
`
	buffer.WriteString(payload)
	c, err := readData(&buffer)
	expected := &Config{
		Host: Host{
			HostName: "mail.example.com",
			Port:     587,
		},
		Credential: Credential{
			UserName: "username@example.com",
			Password: "PassW0rd",
		},
		Sender: EmailAddress{
			Address: "username@example.com",
		},
		Receivers: []EmailAddress{
			{Address: "receiver_01@example.com"},
			{Address: "receiver_02@example.com"},
			{Address: "receiver_03@example.com"},
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, c)
}

func Test_Read_File_No_Exists(t *testing.T) {
	fileName := "/no/exists"
	_, err := Read(fileName)
	assert.ErrorIs(t, err, os.ErrNotExist)
}
