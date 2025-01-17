package handlers

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

type gatewayHandlerMessageReactionAdd struct{}

func (h *gatewayHandlerMessageReactionAdd) EventType() discord.GatewayEventType {
	return discord.GatewayEventTypeMessageReactionAdd
}

func (h *gatewayHandlerMessageReactionAdd) New() any {
	return &discord.GatewayEventMessageReactionAdd{}
}

func (h *gatewayHandlerMessageReactionAdd) HandleGatewayEvent(client bot.Client, sequenceNumber int, shardID int, v any) {
	payload := *v.(*discord.GatewayEventMessageReactionAdd)

	genericEvent := events.NewGenericEvent(client, sequenceNumber, shardID)

	if payload.Member != nil {
		client.Caches().Members().Put(*payload.GuildID, payload.UserID, *payload.Member)
	}

	client.EventManager().DispatchEvent(&events.MessageReactionAdd{
		GenericReaction: &events.GenericReaction{
			GenericEvent: genericEvent,
			MessageID:    payload.MessageID,
			ChannelID:    payload.ChannelID,
			GuildID:      payload.GuildID,
			UserID:       payload.UserID,
			Emoji:        payload.Emoji,
		},
		Member: payload.Member,
	})

	if payload.GuildID == nil {
		client.EventManager().DispatchEvent(&events.DMMessageReactionAdd{
			GenericDMMessageReaction: &events.GenericDMMessageReaction{
				GenericEvent: genericEvent,
				MessageID:    payload.MessageID,
				ChannelID:    payload.ChannelID,
				UserID:       payload.UserID,
				Emoji:        payload.Emoji,
			},
		})
	} else {
		client.EventManager().DispatchEvent(&events.GuildMessageReactionAdd{
			GenericGuildMessageReaction: &events.GenericGuildMessageReaction{
				GenericEvent: genericEvent,
				MessageID:    payload.MessageID,
				ChannelID:    payload.ChannelID,
				GuildID:      *payload.GuildID,
				UserID:       payload.UserID,
				Emoji:        payload.Emoji,
			},
			Member: *payload.Member,
		})
	}
}
