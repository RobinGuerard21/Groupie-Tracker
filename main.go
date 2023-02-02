package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	Artiste()
	Date()
	Location()
	Relation()
	BuildArtist()
	err := InitTemplates()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/index", index)
	http.HandleFunc("/artist", artist)
	http.HandleFunc("/map", maps)
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("./static/"))))



	fmt.Printf("Listening")
	http.ListenAndServe("127.0.0.1:800", nil)
}