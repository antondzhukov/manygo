package manygo

type DynamicBlockBuilder struct {
	Messages        *MessageBuilder
	Buttons         *ButtonsBuilder
	Actions         *ActionBuilder
	QuickReply      *QuickReplyBuilder
	Gallery         *GalleryBuilder
	ExternalMessage *ExternalMessageBuilder
}

func NewDynamicBlockBuilder() *DynamicBlockBuilder {
	return &DynamicBlockBuilder{
		Messages:        &MessageBuilder{},
		Buttons:         &ButtonsBuilder{},
		Actions:         &ActionBuilder{},
		QuickReply:      &QuickReplyBuilder{},
		Gallery:         &GalleryBuilder{},
		ExternalMessage: &ExternalMessageBuilder{},
	}
}

func (builder *DynamicBlockBuilder) BuildDynamicBlock() *DynamicBlock {
	return &DynamicBlock{
		Version: "v2",
	}
}

type DynamicBlock struct {
	Version string `json:"version"`
	Content struct {
		Messages     []*Message    `json:"messages"`
		Actions      []*Action     `json:"actions,omitempty"`
		QuickReplies []*QuickReply `json:"quick_replies,omitempty"`
	} `json:"content"`
}

func (dynamicBlock *DynamicBlock) AddMessage(message *Message) {
	dynamicBlock.Content.Messages = append(dynamicBlock.Content.Messages, message)
}

func (dynamicBlock *DynamicBlock) AddMessages(messages []*Message) {
	dynamicBlock.Content.Messages = append(dynamicBlock.Content.Messages, messages...)
}

func (dynamicBlock *DynamicBlock) AddAction(action *Action) {
	dynamicBlock.Content.Actions = append(dynamicBlock.Content.Actions, action)
}

func (dynamicBlock *DynamicBlock) AddActions(actions []*Action) {
	dynamicBlock.Content.Actions = append(dynamicBlock.Content.Actions, actions...)
}

func (dynamicBlock *DynamicBlock) AddQuickReply(quickReply *QuickReply) {
	dynamicBlock.Content.QuickReplies = append(dynamicBlock.Content.QuickReplies, quickReply)
}

func (dynamicBlock *DynamicBlock) AddQuickReplies(quickReplies []*QuickReply) {
	dynamicBlock.Content.QuickReplies = append(dynamicBlock.Content.QuickReplies, quickReplies...)
}
