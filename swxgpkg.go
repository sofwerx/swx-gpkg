package main

import (
	"database/sql"
	"github.com/edwardfward/swx-gpkg/gpkg"
	"github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {

	var sqlite3conn *sqlite3.SQLiteConn

	sql.Register("sqlite3_test", &sqlite3.SQLiteDriver{
		ConnectHook: func(conn *sqlite3.SQLiteConn) error {
			sqlite3conn = conn
			return nil
		},
	})

	gpkgfilename := os.Args[1]
	// todo add file and permissions checking

	db, err := sql.Open("sqlite3_test", gpkgfilename)
	if err != nil {
		log.Fatal(err)
	}

	gpsdata := gpkg.GPSData{}

	gpsdata.ExtractObsPointsAndSat(db)

	defer db.Close()

}
