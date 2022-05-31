package entities

import (
	"code.db.cafe/wombot/internal/i18n"
)

type Guild struct {
	GuildID    string            `gorm:"primary_key"`
	ChannelID  string            `gorm:"not null"`
	CurrentDay int64             `gorm:"not null"`
	HourOfDay  int64             `gorm:"not null"`
	Language   i18n.EnumLanguage `gorm:"not null"`
}
