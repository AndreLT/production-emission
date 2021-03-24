package datastore

import "production-emission/loader"

type Emission interface {
	Initialize()
	SearchByIso(iso string, limit int) *[]*loader.EmissionYear
}
