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

	sql.Register("geoSql", &sqlite3.SQLiteDriver{
		ConnectHook: func(conn *sqlite3.SQLiteConn) error {
			return nil
		},
	})

	// check command line argument count
	if len(os.Args) < 2 {
		fmt.Printf("command requires sqllite filename.gpkg to " +
			"convert...exiting\n")
		os.Exit(1)
	}

	sqlFile := os.Args[1]

	db, err := sql.Open("geoSql", sqlFile)
	if err != nil {
		log.Fatalf("error: could not open %s", sqlFile)
	}

	gpsData := gpkg.GPSData{}
	schema := gpkg.Observation{}

	c, err := gpsData.ExportToCSV(db, &schema)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Parsed %d records", c)

	defer db.Close()

}
