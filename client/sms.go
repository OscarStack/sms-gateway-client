package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/OscarStack/sms-gateway-client/models"
	"github.com/OscarStack/sms-gateway-client/util"
)

func (tc *TeltonikaClient) SendSms(mobileNumber models.SmsNumber, message string) error {

	if !mobileNumber.Validate() {
		return errors.New("invalid mobile number")
	}
	if len(message) >= 150 {
		return errors.New("message is to long, only 150 characters are allowed")
	}

	status, err := tc.sshClient.Cmd(fmt.Sprintf(`gsmctl --sms --send "%s %s"`, mobileNumber, message)).Output()
	if err != nil {
		return err
	}
	fmt.Println(string(status))
	switch strings.TrimSpace((string(status))) {
	case "OK":
		return nil
	default:
		return errors.New("message not send")
	}
}

func (tc *TeltonikaClient) ReadSmsList() (smsList []models.ReceivedMessage, err error) {
	status, err := tc.sshClient.Cmd(fmt.Sprintf(`gsmctl --sms --list all`)).Output()
	if err != nil {
		return
	}

	lines := string(status)
	linesSplitted := strings.Split(strings.TrimSpace(lines), util.LineSeperator)
	for _, v := range linesSplitted {
		x := models.ReceivedMessage{}
		err := x.FillStruct(util.ToMap(v))
		if err != nil {
			fmt.Println(err)
		}
		smsList = append(smsList, x)
	}

	return
}

func (tc *TeltonikaClient) ReadSmsByIndex(index string) (smsList models.ReceivedMessage, err error) {
	status, err := tc.sshClient.Cmd(fmt.Sprintf(`gsmctl --sms --read %s`, index)).Output()
	if err != nil {
		return
	}
	output := string(status)
	if strings.HasPrefix(output, "N/A") {
		err = errors.New("record not found")
		return
	}

	linesSplitted := strings.Split(strings.TrimSpace(output), util.LineSeperator)
	for _, v := range linesSplitted {
		x := models.ReceivedMessage{}
		err = x.FillStruct(util.ToMap(v))
		if err != nil {
			break
		}
		smsList = x
	}
	return
}

func (tc *TeltonikaClient) DeleteSmsByIndex(index string) (err error) {
	status, err := tc.sshClient.Cmd(fmt.Sprintf(`gsmctl --sms --delete %s`, index)).Output()
	if err != nil {
		return
	}
	output := string(status)

	if strings.HasPrefix(output, "OK") {
		err = nil
		return
	}

	return
}
