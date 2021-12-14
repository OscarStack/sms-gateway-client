package client

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/OscarStack/sms-gateway-client/models"
	"github.com/OscarStack/sms-gateway-client/sshclient"
)

type TeltonikaClient struct {
	sshClient sshclient.Client
}

type TeltonikaHost struct {
	Host     string
	Port     string
	User     string
	Password string
}

func NewTeltonikaClient(params TeltonikaHost) *TeltonikaClient {
	client, err := sshclient.DialWithPasswd(fmt.Sprintf("%s:%s", params.Host, params.Port), params.User, params.Password)
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Close()

	version, err := client.Cmd(`gsmctl --version`).Output()
	if err != nil {
		log.Fatal(err)
	}
	if !strings.HasPrefix(string(version), "GSMCTL version") {
		log.Fatal("Invalid version")
	}
	fmt.Println("OPENED SESSION")
	return &TeltonikaClient{
		sshClient: *client,
	}
}
func (tc *TeltonikaClient) CloseClient() error {
	fmt.Println("Closed client")
	return tc.sshClient.Close()
}

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
