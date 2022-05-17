package commands

import (
	"github.com/bwmarrin/discordgo"
)

type Context struct {
	Session *discordgo.Session
	Message *discordgo.Message
	Args    []string
	Handler *CommandHandler
}

func (c *Context) SendEmbed(embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	msg, err := c.Session.ChannelMessageSendComplex(c.Message.ChannelID, &discordgo.MessageSend{
		Reference: &discordgo.MessageReference{
			MessageID: c.Message.ID,
			ChannelID: c.Message.ChannelID,
			GuildID:   c.Message.GuildID,
		},
		Embed: embed,
	})

	return msg, err
}

func (ctx *Context) SuccessEmbedReturning(embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	embed.Color = 0x50fa7b
	return ctx.SendEmbed(embed)
}

func (ctx *Context) Success(embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	message, err := ctx.SuccessEmbedReturning(embed)

	return message, err
}
