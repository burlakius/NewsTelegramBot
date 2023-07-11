package translator

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var languageTags map[string]language.Tag

func GetPrinter(lang string) *message.Printer {
	return message.NewPrinter(languageTags[lang])
}

func GetAllTranslations(text string) []string {
	result := make([]string, 0, len(languageTags))
	for _, languageTag := range languageTags {
		printer := message.NewPrinter(languageTag)
		result = append(result, printer.Sprintf(text))
	}

	return result
}

func SetupTranslations() {
	languageTags = make(map[string]language.Tag)

	languagesFile, err := os.OpenFile("languages.json", os.O_RDONLY, 444)
	if err != nil {
		logrus.Fatal(err)
	}

	data, err := ioutil.ReadAll(languagesFile)
	if err != nil {
		logrus.Fatal(err)
	}

	var languagesList []string
	err = json.Unmarshal(data, &languagesList)
	if err != nil {
		logrus.Fatal(err)
	}

	for _, lang := range languagesList {
		languageTags[lang] = language.MustParse(lang)
	}

	logrus.Infoln("Translations are set up:", strings.Join(languagesList, " "))
}
