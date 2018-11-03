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
		if err := results.Scan(d.getInterfacePtrs()...); err != nil {
			log.Fatal(err)
		}
		count++
	}

	return count, nil
}
