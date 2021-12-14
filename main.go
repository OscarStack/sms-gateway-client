package main

import (
	"fmt"
	"os"

	"github.com/OscarStack/sms-gateway-client/client"
)

// func main() {

// 	client, err := sshclient.DialWithPasswd("192.168.1.1:22", "root", "Doodle01")
// 	if err != nil {

// 		fmt.Println(err)
// 		return
// 	}
// 	defer client.Close()

// 	version, err := client.Cmd(`gsmctl --version`).Output()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	if !strings.HasPrefix(string(version), "GSMCTL version") {
// 		fmt.Println("INVALID VERSION")
// 		return
// 	}
// 	status, err := client.Cmd(`gsmctl --sms --send "+31648923273 test from a Go Client"`).Output()
// 	fmt.Println(string(status))

// }

func main() {
	tc := client.NewTeltonikaClient(client.TeltonikaHost{
		Host:     "teltonika.stackless.nl",
		Port:     "22",
		User:     "root",
		Password: os.Getenv("PASSWORD"),
	})

	if err := tc.SendSms("0031648923273", "Message from a doodle"); err != nil {
		fmt.Println(err)
	}

	tc.CloseClient()
}

// 0031648923273
// thomas 0031634871262
