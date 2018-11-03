package gpkg

import (
	"database/sql"
	"log"
	"reflect"
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

	dType := reflect.New(reflect.TypeOf(d).Elem()).Interface().(GeoData)
	var row []GeoData

	for results.Next() {
		if err := results.Scan(dType.getInterfacePtrs()...); err != nil {
			log.Fatal(err)
		}
		row = append(row, dType)
	}

	return len(row), nil
}
