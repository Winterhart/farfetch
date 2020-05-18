package cmd

import (
	"fmt"
	"os"

	"github.com/winterhart/farfetch"
)

const (
	sendCommand   = `send`
	uploadCommand = `upload`
	hookEnv       = `SLACK_HOOK`
	tokenEnv      = `SLACK_TOKEN`
	channelEnv    = `SLACK_CH`
)

//main - Sample usage of the farfetch package
func main() {

	hook := os.Getenv(hookEnv)
	token := os.Getenv(tokenEnv)
	channel := os.Getenv(channelEnv)

	if hook == "" || token == "" || channel == "" {
		fmt.Print("One of the config is missing...")
		return
	}
	farfetch := farfetch.NewFarfetchImpl(hook, token, channel)
	command := os.Args[1]

	switch command {
	case sendCommand:
		err := farfetch.SendMessage(os.Args[2])
		if err != nil {
			fmt.Printf("\n Can't send message error: %+v ", err)
		}
	case uploadCommand:
		err := farfetch.UploadFile(os.Args[2])
		if err != nil {
			fmt.Printf("\n Can't upload file error:%+v", err)
		}
	default:
		fmt.Println("Unrecognized command")
	}

}
