package main

import (
	"database/sql"
	"fmt"
	"github.com/edwardfward/swx-gpkg/gpkg"
	"github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {

	sql.Register("sqlite3_test", &sqlite3.SQLiteDriver{
		ConnectHook: func(conn *sqlite3.SQLiteConn) error {
			return nil
		},
	})

	gpkgfilename := os.Args[1]

	db, err := sql.Open("sqlite3_test", gpkgfilename)
	if err != nil {
		log.Fatal(err)
	}

	gpsdata := gpkg.GPSData{}
	schema := gpkg.Observation{}

	c, err := gpsdata.ExportToCSV(db, &schema)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Parsed %d records", c)

	defer db.Close()

}
