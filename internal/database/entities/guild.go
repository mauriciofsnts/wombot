package entities

type Guild struct {
	GuildID    string `gorm:"primary_key"`
	ChannelID  string `gorm:"not null"`
	CurrentDay int64  `gorm:"not null"`
	HourOfDay  int64  `gorm:"not null"`
}
