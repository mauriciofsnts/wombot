package utils

import "code.db.cafe/wombot/internal/discord/slash"

func init() {
	slash.RegisterSlashCommand(Ping, Setup)
}
