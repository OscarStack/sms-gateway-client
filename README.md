# sms-gateway-client
Teltonika SMS Gateway Client 


This library is a integration for the Teltonika gsmctl package. You can setup a ssh connection to your teltonika router and run certain commands. 


## How it works
### Create a client
```go
tc := client.NewTeltonikaClient(client.TeltonikaHost{
	Host:     "192.168.1.1",
	Port:     "22",
	User:     "root",
	Password: os.Getenv("PASSWORD"),
})

// When done
tc.CloseClient()

```

### Send a message 
```go
if err := tc.SendSms("003164823423423", "just a message"); err != nil {
	fmt.Println(err)
}

```
### List messages
```go
list, err := tc.ReadSmsList(models.ALL_MESSAGES)
if err != nil {
	fmt.Println(err)
}
```

### Read message by index 
```go
message, err := tc.ReadSmsByIndex("1")
if err != nil {
	fmt.Println(err)
}
```
### Read latest
```go
latestMessage, err := tc.ReadLatest()
if err != nil {
	fmt.Println(err)
}
```
### Delete message by index
```go
if err := tc.DeleteSmsByIndex("1"); err != nil {
	fmt.Println(err)
}
```

