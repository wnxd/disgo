package handlers

import (
	"github.com/DisgoOrg/disgo/core"
	"github.com/DisgoOrg/disgo/core/events"
	"github.com/DisgoOrg/disgo/discord"
	"github.com/DisgoOrg/disgo/gateway"
)

// InteractionCreateHandler handles api.InteractionCreateGatewayEvent
type InteractionCreateHandler struct{}

// EventType returns the api.GatewayEventType
func (h *InteractionCreateHandler) EventType() gateway.EventType {
	return gateway.EventTypeInteractionCreate
}

// New constructs a new payload receiver for the raw gateway event
func (h *InteractionCreateHandler) New() interface{} {
	return discord.UnmarshalInteraction{}
}

// HandleGatewayEvent handles the specific raw gateway event
func (h *InteractionCreateHandler) HandleGatewayEvent(disgo core.Disgo, eventManager core.EventManager, sequenceNumber int, i interface{}) {
	unmarshalInteraction, ok := i.(discord.UnmarshalInteraction)
	if !ok {
		return
	}
	HandleInteraction(disgo, eventManager, sequenceNumber, nil, unmarshalInteraction)
}

func HandleInteraction(disgo core.Disgo, eventManager core.EventManager, sequenceNumber int, c chan discord.InteractionResponse, unmarshalInteraction discord.UnmarshalInteraction) {
	interaction := disgo.EntityBuilder().CreateInteraction(unmarshalInteraction, c, core.CacheStrategyYes)

	genericInteractionEvent := &events.GenericInteractionEvent{
		GenericEvent: events.NewGenericEvent(disgo, sequenceNumber),
		Interaction:  interaction,
	}

	switch unmarshalInteraction.Type {
	case discord.InteractionTypeCommand:
		eventManager.Dispatch(&events.CommandEvent{
			GenericInteractionEvent: genericInteractionEvent,
			CommandInteraction:      disgo.EntityBuilder().CreateCommandInteraction(interaction, core.CacheStrategyYes),
		})

	case discord.InteractionTypeComponent:
		componentInteraction := disgo.EntityBuilder().CreateComponentInteraction(interaction, core.CacheStrategyYes)

		genericComponentEvent := &events.GenericComponentEvent{
			GenericInteractionEvent: genericInteractionEvent,
			ComponentInteraction:    componentInteraction,
		}

		switch componentInteraction.Data.ComponentType {
		case discord.ComponentTypeButton:
			eventManager.Dispatch(&events.ButtonClickEvent{
				GenericComponentEvent: genericComponentEvent,
				ButtonInteraction:     disgo.EntityBuilder().CreateButtonInteraction(componentInteraction),
			})

		case discord.ComponentTypeSelectMenu:
			eventManager.Dispatch(&events.SelectMenuSubmitEvent{
				GenericComponentEvent: genericComponentEvent,
				SelectMenuInteraction: disgo.EntityBuilder().CreateSelectMenuInteraction(componentInteraction),
			})
		}

	}
}
