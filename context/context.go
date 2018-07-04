package context

import (
	"github.com/Fogapod/KiwiGO/bot"
	"github.com/Fogapod/KiwiGO/command"
	"github.com/bwmarrin/discordgo"
)

type Context struct {
	Bot               *bot.Bot
	Message           *discordgo.Message
	Channel           *discordgo.Channel
	Guild             *discordgo.Guild
	Author            *discordgo.User
	Prefix            string
	RegisterResponses bool
	Commands          []*command.Command
}

// experimental, arguments may change
// TODO: ChannelMessageSendComplex implementation
func (ctx *Context) Send(channelID, content string) (*discordgo.Message, error) {
	if ctx.RegisterResponses {
		// register with redis
	}
	return ctx.Bot.Session.ChannelMessageSend(channelID, content)
}

// experimental, arguments may change
func (ctx *Context) React(emoji string) (*discordgo.MessageReaction, error) {
	if ctx.RegisterResponses {
		// register with redis
	}
	// add reaction
	return nil, nil
}

func MakeContext(b *bot.Bot, msg *discordgo.Message, prefix string) (*Context, error) {
	channel, err := b.Session.State.Channel(msg.ChannelID)
	if err != nil {
		b.Logger.Debug("Channel with id %d not found", msg.ChannelID)
		return nil, err
	}

	var guild *discordgo.Guild

	if msg.GuildID != "" {
		guild, err = b.Session.State.Guild(msg.GuildID)
		if err != nil {
			b.Logger.Debug("Guild with id %d not found", msg.ChannelID)
			return nil, err
		}
	}

	return &Context{
		b,
		msg,
		channel,
		guild,
		msg.Author,
		prefix,
		true,
		[]*command.Command{},
	}, nil
}