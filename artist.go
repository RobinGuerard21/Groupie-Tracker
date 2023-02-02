package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var artistejson []ArtistJSON
var artiste []Artist

func Artiste() {
	if artiste == nil {
		url := "https://groupietrackers.herokuapp.com/api/artists"

		robClient := http.Client{
			Timeout: time.Second * 20, // Timeout after 20 seconds
		}
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
			return
		}
		req.Header.Set("User-Agent", "Groupie Tracker Task")
		res, err := robClient.Do(req)
		if err != nil {
			log.Fatal(err)
			return
		}
		if res.Body != nil {
			defer res.Body.Close()
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = json.Unmarshal(body, &artistejson)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(artistejson[0].Name)
	}
	return
}

var datejson DatesJSON

func Date() {
	if artiste == nil {
		url := "https://groupietrackers.herokuapp.com/api/dates"

		robClient := http.Client{
			Timeout: time.Second * 20, // Timeout after 20 seconds
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
			return
		}
		req.Header.Set("User-Agent", "Groupie Tracker Task")
		res, err := robClient.Do(req)
		if err != nil {
			log.Fatal(err)
			return
		}
		if res.Body != nil {
			defer res.Body.Close()
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = json.Unmarshal(body, &datejson)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	return
}

var locationjson LocationsJSON

func Location() {
	if artiste == nil {
		url := "https://groupietrackers.herokuapp.com/api/locations"

		robClient := http.Client{
			Timeout: time.Second * 20, // Timeout after 20 seconds
		}
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
			return
		}
		req.Header.Set("User-Agent", "Groupie Tracker Task")
		res, err := robClient.Do(req)
		if err != nil {
			log.Fatal(err)
			return
		}
		if res.Body != nil {
			defer res.Body.Close()
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = json.Unmarshal(body, &locationjson)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	return
}

var relationjson RelationsJSON

func Relation() {
	if artiste == nil {
		url := "https://groupietrackers.herokuapp.com/api/relation"

		robClient := http.Client{
			Timeout: time.Second * 20, // Timeout after 20 seconds
		}
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			log.Fatal(err)
			return
		}
		req.Header.Set("User-Agent", "Groupie Tracker Task")
		res, err := robClient.Do(req)
		if err != nil {
			log.Fatal(err)
			return
		}
		if res.Body != nil {
			defer res.Body.Close()
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = json.Unmarshal(body, &relationjson)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	return
}

func BuildArtist() {
	if artiste == nil {
		for _, r := range artistejson {
			g := len(r.Members)
			artiste = append(artiste, Artist{r.Id, r.Image, r.Name, r.Members, g, r.CreationDate, r.FirstAlbum, nil, nil, nil})
		}
		for i, r := range locationjson.Index {
			artiste[i].Locations = r.Locations
		}
		for i, r := range datejson.Index {
			artiste[i].Dates = r.Dates
		}
		for i, r := range relationjson.Index {
			artiste[i].Relation = r.DatesLocation
		}
	}
}
