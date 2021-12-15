package client

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/OscarStack/sms-gateway-client/models"
	"github.com/OscarStack/sms-gateway-client/util"
)

// SendSms sends a message to a mobile number.
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

// ReadSmsList returns all messages in memory
// Filter -> New, All, Read
func (tc *TeltonikaClient) ReadSmsList(filter models.Filter) (smsList []models.ReceivedMessage, err error) {
	status, err := tc.sshClient.Cmd(fmt.Sprintf(`gsmctl --sms --list %s`, filter)).Output()
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

// ReadSmsByIndex returns the message corresponding certain index.
func (tc *TeltonikaClient) ReadSmsByIndex(index string) (message models.ReceivedMessage, err error) {
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
		message = x
	}
	return
}

// DeleteSmsByIndex deletes a message from the memory.
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

// ReadSmsTotal returns our total messages and our limit
func (tc *TeltonikaClient) ReadSmsTotal() (total models.TotalMessage, err error) {
	status, err := tc.sshClient.Cmd(`gsmctl --sms --total`).Output()
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
		x := models.TotalMessage{}
		err = x.FillStruct(util.ToMap(v))
		if err != nil {
			break
		}
		total = x
	}
	return
}

// ReadLatest returns the latest SMS in our memory
func (tc *TeltonikaClient) ReadLatest() (message models.ReceivedMessage, err error) {
	count, err := tc.ReadSmsTotal()
	if err != nil {
		return
	}
	t, err := strconv.Atoi(count.Used)
	if err != nil {
		return
	}
	if t <= 0 {
		err = errors.New("no messages found")
		return
	}
	fmt.Println("looking for ", t)
	message, err = tc.ReadSmsByIndex(fmt.Sprintf("%v", t))
	return
}
