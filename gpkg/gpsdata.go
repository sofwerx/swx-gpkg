package gpkg

import (
	"database/sql"
	csv2 "encoding/csv"
	"log"
	os2 "os"
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

	headers, err := results.Columns()
	if err != nil {
		log.Printf("%q: ", err)
	}

	csv := csv2.NewWriter(os2.Stdout)
	csv.Write(headers)
	csv.Flush()

	recordsOut := 0

	for results.Next() {
		if err := results.Scan(dType.getInterfacePtrs()...); err != nil {
			log.Fatal(err)
		}
		csv.Write(dType.String())
		csv.Flush()
		recordsOut++
	}

	return recordsOut, nil
}
