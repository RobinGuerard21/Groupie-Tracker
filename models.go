package main

type Artist struct {
	Id				int
	Image			string
	Name			string
	Members			[]string
	MembersNbr		int
	CreationDate	int
	FirstAlbum		string
	Locations		[]string
	Dates			[]string
	Relation		map[string][]string
}

type ArtistJSON struct {
	Id				int							`json:"id"`
	Image			string						`json:"image"`
	Name			string						`json:"name"`
	Members			[]string					`json:"members"`
	CreationDate	int							`json:"creationDate"`
	FirstAlbum		string						`json:"firstAlbum"`
}

type DatesJSON struct {
	Index 			[]struct {
		Id			int							`json:"id"`
		Dates		[]string					`json:"dates"`
	}											`json:"index"`

}

type LocationsJSON struct {
	Index 			[]struct {
		Id			int							`json:"id"`
		Locations	[]string					`json:"locations"`
	}											`json:"index"`
}

type RelationsJSON struct {
	Index 			[]struct {
		Id				int						`json:"id"`
		DatesLocation	map[string][]string		`json:"datesLocations"`
	}											`json:"index"`
}