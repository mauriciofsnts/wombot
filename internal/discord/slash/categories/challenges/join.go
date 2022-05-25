package challenges

import (
	"code.db.cafe/wombot/internal/discord/slash"
	"github.com/bwmarrin/discordgo"
)

func init() {
	slash.RegisterSlashCommand(
		&slash.SlashCommand{
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "join",
				Description: "Join a challenge",
			},
			Handler: func(s *discordgo.Session, i *discordgo.InteractionCreate) {

				if i.GuildID == "" {
					
				}

			},
		},
	)
}
