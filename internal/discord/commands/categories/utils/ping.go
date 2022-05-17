package utils

import (
	"code.db.cafe/wombot/internal/discord/commands"
	"github.com/bwmarrin/discordgo"
)

type CmdPing struct{}

func (c *CmdPing) Invokes() []string {
	return []string{"ping", "p"}
}

func (c *CmdPing) Description() string {
	return "Pong!"
}

func (c *CmdPing) AdminRequired() bool {
	return true
}

func (c *CmdPing) Exec(ctx *commands.Context) (err error) {

	ctx.Success(&discordgo.MessageEmbed{
		Description: "Pong!",
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Os cria ta insano hein",
		},
	})

	return nil
}
