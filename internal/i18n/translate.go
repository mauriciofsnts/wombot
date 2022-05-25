package i18n

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type BotLanguage string

var (
	Languages   []BotLanguage
	languageMap map[BotLanguage]*Language = make(map[BotLanguage]*Language)
)

func newLanguage(key BotLanguage) BotLanguage {
	Languages = append(Languages, key)
	return BotLanguage(key)
}

var (
	LanguageEnglish   = newLanguage("en_US")
	LanguageBrazilian = newLanguage("pt_BR")
	LanguageCria      = newLanguage("pt_CRIA")
)

func loadLanguage(lang BotLanguage) error {

	fileName := "internal/i18n/languages/" + string(lang) + ".json"

	data, err := os.ReadFile(fileName)

	if err != nil {
		return err
	}

	var language Language

	err = json.Unmarshal(data, &language)

	if err != nil {
		return err
	}

	languageMap[lang] = &language

	return nil
}

func Start() error {

	for _, lang := range Languages {
		if err := loadLanguage(lang); err != nil {
			return err
		}
	}

	return nil
}

func Translate(lang BotLanguage) *Language {
	language, found := languageMap[lang]

	//  tá enviando um bgl erradao irmao
	if !found {
		//fallback de cria?
		language = languageMap[LanguageCria]
	}

	return language
}

func Replace(str string, args ...interface{}) string {

	for i, value := range args {
		str = strings.Replace(str, fmt.Sprintf("${%d}", i), fmt.Sprintf("%v", value), 1)
	}

	return str

}
