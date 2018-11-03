package gpkg

import (
	"database/sql"
	"log"
)

type GPSData struct {
	sqlfile string
}

func (g *GPSData) ExportToCSV(db *sql.DB, d GeoData) (int, error) {
	qrystr := d.getSQLQuery()

	results, err := db.Query(qrystr)
	if err != nil {
		log.Printf("%q: %s", err, qrystr)
	}

	count := 0

	for results.Next() {
		var obs = Observation{}
		if err := results.Scan(obs.ObsInterfaces()...); err != nil {
			log.Fatal(err)
		}
		count++
	}

	return count, nil
}

func (g *GPSData) ExtractObsPointsAndSat(db *sql.DB) error {

	return nil
}

