package main

import (
	"os"

	"github.com/OscarStack/sms-gateway-client/client"
)

func main() {
	tc := client.NewTeltonikaClient(client.TeltonikaHost{
		Host:     "95.97.89.20",
		Port:     "22",
		User:     "root",
		Password: os.Getenv("PASSWORD"),
	})

	// if err := tc.SendSms("0031648923273", "woef de honger"); err != nil {
	// 	fmt.Println(err)
	// }
	// for {
	// 	list, err := tc.ReadSmsList()
	// 	fmt.Print("\033[H\033[2J")
	// 	fmt.Printf("WATCHING\n")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	for _, v := range list {
	// 		fmt.Println(v)
	// 	}
	// 	time.Sleep(2 * time.Second)

	// }

	tc.CloseClient()
}

// 0031648923273
// thomas 0031634871262
