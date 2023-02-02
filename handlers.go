package main

import (
	"fmt"
	htmlTemplate "html/template"
	"net/http"
)

type templateArtist struct {
	Artist []Artist
}

func index(w http.ResponseWriter, r *http.Request) {
	templ, err := htmlTemplate.ParseFiles("html/index.page.tmpl")
	templ.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func artist(w http.ResponseWriter, r *http.Request) {
	templ, err := htmlTemplate.ParseFiles("html/artist.page.tmpl")
	templ.Execute(w, templateArtist{
		Artist: artiste,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func maps(w http.ResponseWriter, r *http.Request) {
	templ, err := htmlTemplate.ParseFiles("html/maps.page.tmpl")
	templ.Execute(w, templateArtist{
		Artist: artiste,
	})
	if err != nil {
		fmt.Println(err)
	}
}
