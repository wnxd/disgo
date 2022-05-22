package events

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/snowflake/v2"
)

type EmojisUpdateEvent struct {
	*GenericEvent
	discord.GatewayEventGuildEmojisUpdate
}

// GenericEmojiEvent is called upon receiving EmojiCreateEvent, EmojiUpdateEvent or EmojiDeleteEvent (requires discord.GatewayIntentGuildEmojisAndStickers)
type GenericEmojiEvent struct {
	*GenericEvent
	GuildID snowflake.ID
	Emoji   discord.Emoji
}

// EmojiCreateEvent indicates that a new discord.Emoji got created in a discord.Guild (requires discord.GatewayIntentGuildEmojisAndStickers)
type EmojiCreateEvent struct {
	*GenericEmojiEvent
}

// EmojiUpdateEvent indicates that a discord.Emoji got updated in a discord.Guild (requires discord.GatewayIntentGuildEmojisAndStickers)
type EmojiUpdateEvent struct {
	*GenericEmojiEvent
	OldEmoji discord.Emoji
}

// EmojiDeleteEvent indicates that a discord.Emoji got deleted in a discord.Guild (requires discord.GatewayIntentGuildEmojisAndStickers)
type EmojiDeleteEvent struct {
	*GenericEmojiEvent
}
