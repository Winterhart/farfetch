# Farfetch

A tiny piece software to send files and messages using slack API


![](http://pixelartmaker.com/art/f740b9b40b4942d.png)

Image from: http://pixelartmaker.com/art/f740b9b40b4942d

## Overview

- Library to send message and upload file to Slack
- You can compile and use as binary 
- You can also import and use it in your code

## Setup

This library needs three variables to work:
- A slack-web hook (`hook`)
- A slack token (`token`)
- You also need to choose a slack channel (`channel id`)

### How to get the hook?

- https://api.slack.com/messaging/webhooks

### How to get the token?

- 

### How to get the channel id?

- https://www.wikihow.com/Find-a-Channel-ID-on-Slack-on-PC-or-Mac

## Usage

You can use this project in two distinct ways.

#### as a binary

Append your `env` with the following variable: `SLACK_HOOK` your web-hook, `SLACK_TOKEN` your token 
and `SLACK_CH` your channel id.


Then build the binary

```bash

go build cmd/main.go

```
Make the generated `farfetch` binary file executable
```bash

chmod +x ./farfetch

```
Then you can use the binary:

```bash

./farfetch send hello

```

Will send `hello` using your slack-bot to your designated slack channel. 

```bash

./farfetch upload ~/pictures/photo.png

```

Will upload the file `photo.png` to your designated slack channel.

#### as a package 

You can use this project as a package in your golang project.

```golang
package main

import (
	"os"
	"github.com/winterhart/farfetch"
)

func main() {
	webHook := os.Getenv(`SLACK_HOOK`)
	slack := farfetch.NewFarfetchImpl(webHook, "", "")
	slack.SendMessage("This message is from another project using Farfetch! ")

}

```



