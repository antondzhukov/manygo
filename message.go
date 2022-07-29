package manygo

const (
	MessageTypeText  string = "text"
	MessageTypeImage string = "image"
	MessageTypeVideo string = "video"
	MessageTypeAudio string = "audio"
	MessageTypeFile  string = "file"
	MessageTypeCards string = "cards"
)

type MessageBuilder struct{}

type Message struct {
	Type             string         `json:"type"`
	Text             string         `json:"text,omitempty"`
	Url              string         `json:"url,omitempty"`
	Buttons          []*Button      `json:"buttons,omitempty"`
	Elements         []*GalleryCard `json:"elements,omitempty"`
	ImageAspectRatio string         `json:"image_aspect_ratio"`
}

func (message *Message) AddButton(button *Button) {
	message.Buttons = append(message.Buttons, button)
}

func (message *Message) AddButtons(buttons []*Button) {
	message.Buttons = append(message.Buttons, buttons...)
}

func (builder *MessageBuilder) BuildTextMessage(text string) *Message {
	message := &Message{
		Type: MessageTypeText,
		Text: text,
	}

	return message
}

func (builder *MessageBuilder) BuildImageMessage(url string) *Message {
	message := &Message{
		Type: MessageTypeImage,
		Url:  url,
	}

	return message
}

func (builder *MessageBuilder) BuildVideoMessage(url string) *Message {
	message := &Message{
		Type: MessageTypeVideo,
		Url:  url,
	}

	return message
}

func (builder *MessageBuilder) BuildAudioMessage(url string) *Message {
	message := &Message{
		Type: MessageTypeAudio,
		Url:  url,
	}

	return message
}

func (builder *MessageBuilder) BuildFileMessage(url string) *Message {
	message := &Message{
		Type: MessageTypeFile,
		Url:  url,
	}

	return message
}

func (builder *MessageBuilder) BuildGalleryMessage(gallery *Gallery) *Message {
	message := &Message{
		Type:             MessageTypeCards,
		Elements:         gallery.Elements,
		ImageAspectRatio: gallery.GetImageAspectRation(),
	}

	return message
}
