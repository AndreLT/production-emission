package datastore

import (
	"log"
	"os"
	"production-emission/loader"
	"strings"
)

type Books struct {
	Store *[]*loader.EmissionYear `json:"store"`
}

func (e *Books) Initialize() {
	filename := "./assets/emission.csv"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	e.Store = loader.LoadData(file)
}

func (e *Books) SearchByIso(iso string, limit int) *[]*loader.EmissionYear {
	ret := Filter(e.Store, func(v *loader.EmissionYear) bool {
		return strings.Contains(strings.ToLower(v.IsoCode), strings.ToLower(iso))
	})

	if limit > len(*ret) {
		data := (*ret)
		return &data
	} else {
		if limit == 0 {
			limit = len(*ret)
		}
		data := (*ret)[len(*ret)-limit:]
		return &data
	}
}

func Filter(vs *[]*loader.EmissionYear, f func(*loader.EmissionYear) bool) *[]*loader.EmissionYear {
	vsf := make([]*loader.EmissionYear, 0)
	for _, v := range *vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return &vsf
}
