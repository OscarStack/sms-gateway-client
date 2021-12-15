package client

import (
	"fmt"
	"log"
	"strings"

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
