package webhook

import (
	"fmt"
	"io"

	"github.com/DisgoOrg/disgo/discord"
)

// MessageUpdateBuilder helper to build MessageUpdate easier
type MessageUpdateBuilder struct {
	discord.WebhookMessageUpdate
}

// NewMessageUpdateBuilder creates a new MessageUpdateBuilder to be built later
//goland:noinspection GoUnusedExportedFunction
func NewMessageUpdateBuilder() *MessageUpdateBuilder {
	return &MessageUpdateBuilder{
		WebhookMessageUpdate: discord.WebhookMessageUpdate{
			AllowedMentions: &DefaultAllowedMentions,
		},
	}
}

// SetContent sets content of the Message
func (b *MessageUpdateBuilder) SetContent(content string) *MessageUpdateBuilder {
	b.Content = &content
	return b
}

// SetContentf sets content of the Message
func (b *MessageUpdateBuilder) SetContentf(content string, a ...interface{}) *MessageUpdateBuilder {
	return b.SetContent(fmt.Sprintf(content, a...))
}

// ClearContent removes content of the Message
func (b *MessageUpdateBuilder) ClearContent() *MessageUpdateBuilder {
	return b.SetContent("")
}

// SetEmbeds sets the Embed(s) of the Message
func (b *MessageUpdateBuilder) SetEmbeds(embeds ...discord.Embed) *MessageUpdateBuilder {
	if b.Embeds == nil {
		b.Embeds = new([]discord.Embed)
	}
	*b.Embeds = embeds
	return b
}

// SetEmbed sets the provided Embed at the index of the Message
func (b *MessageUpdateBuilder) SetEmbed(i int, embed discord.Embed) *MessageUpdateBuilder {
	if b.Embeds == nil {
		b.Embeds = new([]discord.Embed)
	}
	if len(*b.Embeds) > i {
		(*b.Embeds)[i] = embed
	}
	return b
}

// AddEmbeds adds multiple embeds to the Message
func (b *MessageUpdateBuilder) AddEmbeds(embeds ...discord.Embed) *MessageUpdateBuilder {
	if b.Embeds == nil {
		b.Embeds = new([]discord.Embed)
	}
	*b.Embeds = append(*b.Embeds, embeds...)
	return b
}

// ClearEmbeds removes all the embeds from the Message
func (b *MessageUpdateBuilder) ClearEmbeds() *MessageUpdateBuilder {
	b.Embeds = &[]discord.Embed{}
	return b
}

// RemoveEmbed removes an embed from the Message
func (b *MessageUpdateBuilder) RemoveEmbed(i int) *MessageUpdateBuilder {
	if b.Embeds == nil {
		b.Embeds = new([]discord.Embed)
	}
	if len(*b.Embeds) > i {
		*b.Embeds = append((*b.Embeds)[:i], (*b.Embeds)[i+1:]...)
	}
	return b
}

// SetActionRows sets the ActionRowComponent(s) of the Message
func (b *MessageUpdateBuilder) SetActionRows(actionRows ...discord.ActionRowComponent) *MessageUpdateBuilder {
	if b.Components == nil {
		b.Components = new([]discord.Component)
	}
	*b.Components = actionRowsToComponents(actionRows)
	return b
}

// SetActionRow sets the provided ActionRowComponent at the index of Component(s)
func (b *MessageUpdateBuilder) SetActionRow(i int, actionRow discord.ActionRowComponent) *MessageUpdateBuilder {
	if b.Components == nil {
		b.Components = new([]discord.Component)
	}
	if len(*b.Components) > i {
		(*b.Components)[i] = actionRow
	}
	return b
}

// AddActionRow adds a new ActionRowComponent with the provided Component(s) to the Message
func (b *MessageUpdateBuilder) AddActionRow(components ...discord.Component) *MessageUpdateBuilder {
	if b.Components == nil {
		b.Components = new([]discord.Component)
	}
	*b.Components = append(*b.Components, discord.NewActionRow(components...))
	return b
}

// AddActionRows adds the ActionRowComponent(s) to the Message
func (b *MessageUpdateBuilder) AddActionRows(actionRows ...discord.ActionRowComponent) *MessageUpdateBuilder {
	if b.Components == nil {
		b.Components = new([]discord.Component)
	}
	*b.Components = append(*b.Components, actionRowsToComponents(actionRows)...)
	return b
}

// RemoveActionRow removes a ActionRowComponent from the Message
func (b *MessageUpdateBuilder) RemoveActionRow(i int) *MessageUpdateBuilder {
	if b.Components == nil {
		b.Components = new([]discord.Component)
	}
	if len(*b.Components) > i {
		*b.Components = append((*b.Components)[:i], (*b.Components)[i+1:]...)
	}
	return b
}

// ClearActionRows removes all the ActionRowComponent(s) of the Message
func (b *MessageUpdateBuilder) ClearActionRows() *MessageUpdateBuilder {
	b.Components = &[]discord.Component{}
	return b
}

// SetFiles sets the restclient.File(s) for this MessageCreate
func (b *MessageUpdateBuilder) SetFiles(files ...*discord.File) *MessageUpdateBuilder {
	b.Files = files
	return b
}

// SetFile sets the restclient.File at the index for this MessageCreate
func (b *MessageUpdateBuilder) SetFile(i int, file *discord.File) *MessageUpdateBuilder {
	if len(b.Files) > i {
		b.Files[i] = file
	}
	return b
}

// AddFiles adds the restclient.File(s) to the MessageCreate
func (b *MessageUpdateBuilder) AddFiles(files ...*discord.File) *MessageUpdateBuilder {
	b.Files = append(b.Files, files...)
	return b
}

// AddFile adds a restclient.File to the MessageCreate
func (b *MessageUpdateBuilder) AddFile(name string, reader io.Reader, flags ...discord.FileFlags) *MessageUpdateBuilder {
	b.Files = append(b.Files, discord.NewFile(name, reader, flags...))
	return b
}

// ClearFiles removes all files of this MessageCreate
func (b *MessageUpdateBuilder) ClearFiles() *MessageUpdateBuilder {
	b.Files = []*discord.File{}
	return b
}

// RemoveFiles removes the file at this index
func (b *MessageUpdateBuilder) RemoveFiles(i int) *MessageUpdateBuilder {
	if len(b.Files) > i {
		b.Files = append(b.Files[:i], b.Files[i+1:]...)
	}
	return b
}

// RetainAttachments removes all Attachment(s) from this Message except the ones provided
func (b *MessageUpdateBuilder) RetainAttachments(attachments ...discord.Attachment) *MessageUpdateBuilder {
	if b.Attachments == nil {
		b.Attachments = new([]discord.Attachment)
	}
	*b.Attachments = append(*b.Attachments, attachments...)
	return b
}

// RetainAttachmentsByID removes all Attachment(s) from this Message except the ones provided
func (b *MessageUpdateBuilder) RetainAttachmentsByID(attachmentIDs ...discord.Snowflake) *MessageUpdateBuilder {
	if b.Attachments == nil {
		b.Attachments = new([]discord.Attachment)
	}
	for _, attachmentID := range attachmentIDs {
		*b.Attachments = append(*b.Attachments, discord.Attachment{ID: attachmentID})
	}
	return b
}

// SetAllowedMentions sets the AllowedMentions of the Message
func (b *MessageUpdateBuilder) SetAllowedMentions(allowedMentions *discord.AllowedMentions) *MessageUpdateBuilder {
	b.AllowedMentions = allowedMentions
	return b
}

// ClearAllowedMentions clears the allowed mentions of the Message
func (b *MessageUpdateBuilder) ClearAllowedMentions() *MessageUpdateBuilder {
	return b.SetAllowedMentions(nil)
}

// Build builds the MessageUpdateBuilder to a MessageUpdate struct
func (b *MessageUpdateBuilder) Build() discord.WebhookMessageUpdate {
	return b.WebhookMessageUpdate
}
