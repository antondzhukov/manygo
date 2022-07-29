## ManyChat API SDK for Golang

SDK based on the ManyChat public API. Supporting account, subscribers and sending.

To get entire information about endpoints, please read the official ManyChat documentation.

[ManyChat API Swagger](https://api.manychat.com/swagger#/)

> Please keep in your mind that ManyChat API has rate limits on each endpoint.
You can find current limits in [rate_limits.go](rate_limit.go) and control requests sending.

#### Features

* Page (info, tags, bot fields, custom fields, flows, growth tools, OTN topics)
* Subscriber (search, create and update subscribers, control custom fields, system fields, tags)
* Sending (send content, nodes and flows)
* Ready builders to construct messages

#### Installation

```
> go get github.com/antondzhukov/manygo
```

```go
package main

import (
	"fmt"
	"github.com/antondzhukov/manygo"
	"net/http"
)
```

#### Client initialization and create first request

```go
client := manygo.NewClient(http.DefaultClient, "YOUR_API_KEY")

page, httpResponse, err := client.Page.GetInfo()
```

### Sending
#### Send a Message

```go
message := client.DynamicBlock.Messages.BuildTextMessage("Hello world!")
dynamicBlock := client.DynamicBlock.BuildDynamicBlock()
dynamicBlock.AddMessage(message)

sendContentRequest := manygo.SendingSendContentRequest{
    SubscriberId: 12345,
    Data: dynamicBlock,
    MessageTag: manygo.MessageTagAccountUpdate,
}

ok, httpResponse, err := client.Sending.SendContent(sendContentRequest)
```

#### Send several Messages

```go
dynamicBlock := client.DynamicBlock.BuildDynamicBlock()

messageOne := client.DynamicBlock.Messages.BuildTextMessage("Hello world one!")
messageTwo := client.DynamicBlock.Messages.BuildTextMessage("Hello world two!")

dynamicBlock.AddMessage(messageOne)
dynamicBlock.AddMessage(messageTwo)

// OR

messages := make([]*Message, 0)
messages = append(messages, messageOne)
messages = append(messages, messagetwo)

dynamicBlock.AddMessages(messages)

sendContentRequest := manygo.SendingSendContentRequest{
    SubscriberId: 12345,
    Data: dynamicBlock,
    MessageTag: manygo.MessageTagAccountUpdate,
}

ok, httpResponse, err := client.Sending.SendContent(sendContentRequest)
```

### Build Messages

```go
// Text
message := client.DynamicBlock.Messages.BuildTextMessage("Hello world!")

// Image
message := client.DynamicBlock.Messages.BuildImageMessage("https://domain.com/image.png")

// Video
message := client.DynamicBlock.Messages.BuildVideoMessage("https://domain.com/video.mov")

// Audio
message := client.DynamicBlock.Messages.BuildAudioMessage("https://domain.com/video.wav")

// File
message := client.DynamicBlock.Messages.BuildVideoMessage("https://domain.com/doc.doc")

// Gallery
gallery := client.DynamicBlock.Gallery.NewGallery()
gallery.SetImageAspectRation(ImageAspectRationSquare) // default: ImageAspectRatioHorizontal
gallery.AddGalleryCard(
    "Title",
    "subtitle",
    "https://domain.com/image.jpg",
    "https://domain.com/actionPage", // use empty string ("") if you don't need an action
    make([]Button, 0), // add buttons if you want, or use an empty slice
)

message := client.DynamicBlock.Messages.BuildGalleryMessage(gallery)
```

#### Add Buttons to a Message

```go
button := client.DynamicBlock.Buttons.BuildUrlButton("Click", "https://domain.com", WebViewSizeFull)
message.AddButton(button)
```

see how to create other [Buttons](button.go)

#### Add Actions to a DynamicBlock

```go
action := client.DynamicBlock.Actions.BuildAddTagAction("My Tag")
dynamicBlock.AddAction(action)
```

see how to create other [Actions](action.go)

#### Add QuickReplies to a DynamicBlock

```go
quickReply := client.DynamicBlock.QuickReply.BuildFlowQuickReply("reply quickly", "ns_123abc...")
dynamicBlock.AddQuickReply(quickReply)
```

see how to create other [QuickReplies](quick_reply.go)

