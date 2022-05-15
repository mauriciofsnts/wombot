package utils

import (
	"log"

	"code.db.cafe/wombot/internal/discord/commands"
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
	log.Println("Pong!")
	_, err = ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, "Pong!")
	return
}
