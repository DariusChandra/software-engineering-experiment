package main

import (
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"net/http"
)

var localizer *i18n.Localizer //1
var bundle *i18n.Bundle       //2
func init() { //3
	bundle = i18n.NewBundle(language.English)                                                  //4
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)                                       //5
	bundle.LoadMessageFile("resources/en.json")                                                //6
	bundle.LoadMessageFile("resources/fr.json")                                                //7
	localizer = i18n.NewLocalizer(bundle, language.English.String(), language.French.String()) //8
	http.HandleFunc("/setlang/", SetLangPreferences)
	http.HandleFunc("/localize/", Localize) //1
	http.ListenAndServe(":8080", nil)       //2
}

func SetLangPreferences(_ http.ResponseWriter, request *http.Request) {
	lang := request.FormValue("lang")                   //1
	accept := request.Header.Get("Accept-Language")     //2
	localizer = i18n.NewLocalizer(bundle, lang, accept) //3
}

func Localize(responseWriter http.ResponseWriter, request *http.Request) {
	valToLocalize := request.URL.Query().Get("msg") //1
	localizeConfig := i18n.LocalizeConfig{          //2
		MessageID: valToLocalize,
	}
	localization, _ := localizer.Localize(&localizeConfig) //3
	fmt.Fprintln(responseWriter, localization)             //4
}

func main() {
	withFiles()
}

func withFiles() {
	localizeConfigWelcome := i18n.LocalizeConfig{
		MessageID: "welcome", //1
	}
	localizationUsingJson, _ := localizer.Localize(&localizeConfigWelcome) //2
	fmt.Println(localizationUsingJson)
}

func withoutFiles() {
	messageEn := i18n.Message{ //1
		ID:    "hello",
		Other: "Hello!",
	}
	messageFr := i18n.Message{ //2
		ID:    "hello",
		Other: "Bonjour!",
	}

	bundle := i18n.NewBundle(language.English)       //1
	bundle.AddMessages(language.English, &messageEn) //2
	bundle.AddMessages(language.French, &messageFr)  //3
	localizer := i18n.NewLocalizer(bundle,           //4
		language.French.String(),
		language.English.String())
	localizeConfig := i18n.LocalizeConfig{ //5
		MessageID: "hello",
	}
	localization, _ := localizer.Localize(&localizeConfig)
	fmt.Println(localization)

	defaultmessageEn := i18n.Message{
		ID:    "welcome",
		Other: "Welcome to my app!",
	}

	localizeConfigWithDefault := i18n.LocalizeConfig{
		MessageID:      "welcome",         //1
		DefaultMessage: &defaultmessageEn, //2
	}
	localizationReturningDefault, _ := localizer.Localize(&localizeConfigWithDefault)
	fmt.Println(localizationReturningDefault)
}
