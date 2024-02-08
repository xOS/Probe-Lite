package singleton

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"

	"github.com/xos/serverstatus/model"
)

var Localizer *i18n.Localizer

func InitLocalizer() {
	bundle := i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	if _, exists := model.Languages[Conf.Language]; !exists {
		log.Println("NG>> language not exists:", Conf.Language)
		Conf.Language = "zh-CN"
	} else {
		_, err := bundle.LoadMessageFileFS(resource.I18nFS, "l10n/"+Conf.Language+".toml")
		if err != nil {
			panic(err)
		}
	}

	if _, err := bundle.LoadMessageFileFS(resource.I18nFS, "l10n/zh-CN.toml"); err != nil {
		panic(err)
	}
	Localizer = i18n.NewLocalizer(bundle, Conf.Language)
}
