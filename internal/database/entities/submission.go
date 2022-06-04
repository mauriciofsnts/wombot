package entities

type Submission struct {
	UserId     string `gorm:"primary_key"`
	User       User
	GuildId    string
	Guild      Guild
	MessageId  string `gorm:"not null"`
	CurrentDay int64  `gorm:"not null"`
}
