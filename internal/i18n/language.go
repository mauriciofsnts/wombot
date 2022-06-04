package i18n

type LanguageEntry string

func (l LanguageEntry) Str(args ...interface{}) string {
	return Replace(string(l), args...)
}

type LanguageMetadata struct {
	Name      LanguageEntry
	ShortName LanguageEntry
	Author    LanguageEntry
}

type Errors struct {
	Title                LanguageEntry
	Generic              LanguageEntry
	GenericGif           LanguageEntry
	ToSave               LanguageEntry
	ToSaveGif            LanguageEntry
	NotATextChannel      LanguageEntry
	AlreadyRegistered    LanguageEntry
	AlreadyRegisteredGif LanguageEntry
}

type Command struct {
	Title    LanguageEntry
	Response LanguageEntry
	Gif      LanguageEntry
}

type Commands struct {
	Ping  Command
	Setup Command
	Join  Command
}

type Language struct {
	Lang       LanguageMetadata
	Errors     Errors
	Commands   Commands
	Challenges Challenges
}

type Challenges struct {
	Title       LanguageEntry
	Description LanguageEntry
	Footer      LanguageEntry
}
