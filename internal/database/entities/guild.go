package entities

type Guild struct {
	GuildID    string `gorm:"primary_key"`
	ChannelID  string `gorm:"not null"`
	CurrentDay int    `gorm:"not null"`
	HourOfDay  int    `gorm:"not null"`
	Language   string `gorm:"not null"`
	DayGoal    int    `gorm:"not null"`
}
