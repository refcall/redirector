package main

import (
	"strconv"
	"strings"
)

type LangQ struct {
	Lang string
	Q    float64
}

func ParseAcceptLanguage(acptLang string) []LangQ {
	var lqs []LangQ

	langQStrs := strings.Split(acptLang, ",")
	for _, langQStr := range langQStrs {
		trimedLangQStr := strings.Trim(langQStr, " ")

		langQ := strings.Split(trimedLangQStr, ";")
		if len(langQ) == 1 {
			lq := LangQ{langQ[0], 1}
			lqs = append(lqs, lq)
		} else {
			qp := strings.Split(langQ[1], "=")
			q, err := strconv.ParseFloat(qp[1], 64)
			if err != nil {
				panic(err)
			}
			lq := LangQ{langQ[0], q}
			lqs = append(lqs, lq)
		}
	}
	return lqs
}

type Translation struct {
	Text     string
	Redirect string
}

var Translations = map[string]*Translation{
	"fr": &Translation{
		Text:     "Nous vous redirigeons vers la page demandée, veuillez patienter quelques secondes.",
		Redirect: "Vous n'êtes pas redirigé ?",
	},
	"en": &Translation{
		Text:     "We are redirecting you to the requested page, please wait a few seconds.",
		Redirect: "You're not redirected?",
	},
}
