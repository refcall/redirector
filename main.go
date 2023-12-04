package main

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/kelseyhightower/envconfig"
)

//go:embed page.tpl
var tmpl string

type Config struct {
	Name                    string `default:"Automatic"`
	Destination             string `required:"true"`
	Logo                    string
	BackgroundColor         string `default:"white"`
	CardColor               string `default:"white"`
	ProgressBackgroundColor string `default:"gray"`
	ProgressColor           string `default:"black"`
	ProgressSeconds         uint   `default:"4"`
	RedirectSeconds         uint   `default:"2"`
}

func main() {
	var c Config
	err := envconfig.Process("redirector", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	tmpl, err := template.New("page").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	myHandler := MyHandler{
		Template: tmpl,
		Config:   c,
	}

	s := &http.Server{
		Addr:         ":8096",
		Handler:      myHandler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}

type MyHandler struct {
	http.Handler
	Template *template.Template
	Config   Config
}

type Template struct {
	Config      Config
	Translation Translation
}

func (m MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/redirect" {
		ref := r.Header.Get("Referer")
		log.Println("user redirect with referer", ref)

		if ref == "" {
			http.Redirect(w, r, m.Config.Destination, http.StatusTemporaryRedirect)
			return
		}

		url, err := url.Parse(ref)
		if err != nil {
			http.Redirect(w, r, m.Config.Destination, http.StatusTemporaryRedirect)
			return
		}

		redirect := m.Config.Destination + url.Path + "?" + url.RawQuery
		http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
		return
	}

	lngs := ParseAcceptLanguage(r.Header.Get("Accept-Language"))
	var translation *Translation
	for _, l := range lngs {
		if Translations[l.Lang] != nil {
			translation = Translations[l.Lang]
			break
		}
	}
	if translation == nil {
		translation = Translations["en"]
	}

	m.Template.Execute(w, &Template{
		Config:      m.Config,
		Translation: *translation,
	})
}
