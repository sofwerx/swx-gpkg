package gpkg

import (
	"database/sql"
	"fmt"
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

	count := 0

	for results.Next() {
		dType := reflect.New(reflect.TypeOf(d).Elem()).Interface().(GeoData)
		if err := results.Scan(dType.getInterfacePtrs()...); err != nil {
			log.Fatal(err)
		}
		fmt.Println(dType)
	}

	return count, nil
}
