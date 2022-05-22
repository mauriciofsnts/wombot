package slash

import (
	"github.com/Pauloo27/logger"
	"github.com/bwmarrin/discordgo"
)

type SlashCommand struct {
	*discordgo.ApplicationCommand
	Handler func(s *discordgo.Session, m *discordgo.InteractionCreate)
}

var commands = make(map[string]*SlashCommand)

func RegisterSlashCommand(cmds ...*SlashCommand) {
	for _, command := range cmds {
		commands[command.Name] = command
	}
}

// cringe
func Start(s *discordgo.Session) error {

	applicationCommands := make([]*discordgo.ApplicationCommand, len(commands))

	i := 0
	for _, c := range commands {
		logger.Debugf("Registering command: %s", c.Name)

		applicationCommands[i] = c.ApplicationCommand
		i++
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		commandName := i.ApplicationCommandData().Name

		if command, ok := commands[commandName]; ok {
			command.Handler(s, i)
		}

	})

	_, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", applicationCommands)

	return err
}
