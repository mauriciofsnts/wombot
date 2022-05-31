package i18n

import (
	"fmt"
	"os"
	"strings"

	"github.com/Pauloo27/logger"
	"github.com/ghodss/yaml"
)

type EnumLanguage string

var (
	Languages   []EnumLanguage
	languageMap map[EnumLanguage]*Language = make(map[EnumLanguage]*Language)
)

func newLanguageEnum(key EnumLanguage) EnumLanguage {
	Languages = append(Languages, key)
	return EnumLanguage(key)
}

var (
	LanguageEnglish   = newLanguageEnum("en_US")
	LanguageBrazilian = newLanguageEnum("pt_BR")
	LanguageCria      = newLanguageEnum("pt_CRIA")
)

func loadLanguage(lang EnumLanguage) error {

	fileName := "internal/i18n/languages/" + string(lang) + ".yml"

	data, err := os.ReadFile(fileName)

	if err != nil {
		return err
	}

	var language Language

	err = yaml.Unmarshal(data, &language)

	if err != nil {
		return err
	}

	languageMap[lang] = &language

	return nil
}

func init() {

	for _, lang := range Languages {
		if err := loadLanguage(lang); err != nil {
			logger.Fatal("Failed to load languages snif snif")
		}
	}

}

func GetLanguage(lang EnumLanguage) *Language {
	language, found := languageMap[lang]

	if !found {
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
