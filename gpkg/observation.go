package gpkg

import (
	"database/sql"
	"log"
	"time"
)

// observation contains gps observations and satellite data

type GPSData struct {
	sqlfile    string
	gpsSatData [][]int32
	data       []*GPSObservation
	gpsobs     []GPSObservation
	satobs     []SatObs
}

type GPSObservation struct {
	id                  uint32    `json:"id, omitempty"`
	systime             time.Time `json:"systime_nanos, omitempty"`
	lat                 float64   `json:"latitude, omitempty"`
	lon                 float64   `json:"longitude, omitempty"`
	alt                 float64   `json:"altitude_ft, omitempty"`
	provider            string    `json:"provider, omitempty"`
	gpsTime             uint64    `json:"gps_time, omitempty"`
	fixedSatCount       int8      `json:"sat_count, omitempty"`
	satobs              []*SatObs `json:"sats, omitempty"`
	hasRadialAccuracy   bool      `json:"has_radial_accuracy, omitempty"`
	hasVerticalAccuracy bool      `json:"has_vertical_accuracy, omitempty"`
	radialAccuracy      float64   `json:"radial_accuracy, omitempty"`
	verticalAccuracy    float64   `json:"vertical_accuracy, omitempty"`
	dataDump            string    `json:"data_dump, omitempty"`
	speed               float64   `json:"speed, omitempty"`
	speedAccuracy       float64   `json:"speed_accuracy, omitempty"`
}

type SatObs struct {
	id                    uint32  `json:"id, omitempty"`
	localtime             int64   `json:"local_time_nanos, omitempty"`
	svid                  uint8   `json:"satellite_id, omitempty"`
	constellation         string  `json:"constellation, omitempty"`
	cn0                   float32 `json:"cn0, omitempty"`
	agc                   float32 `json:"automatic_gain_control, omitempty"`
	hasAgc                bool    `json:"has_agc, omitempty"`
	infix                 bool    `json:"in_fix, omitempty"`
	satTimeNanos          float64 `json:"satellite_time_nanos, omitempty"`
	satTime1signmaNanos   float64 `json:"satellite_sigma_nanos, omitempty"`
	hasCarrierFreq        bool    `json:"has_carrier_freq, omitempty"`
	psuedorangeRateMps    float64 `json:"psuedo_range_rate_mps, omitempty"`
	psuedorangeRate1Sigma float64 `json:"psuedo_range_rate_sigma, omitempty"`
	hasEphemeris          bool    `json:"has_ephemeris, omitempty"`
	azimuthDeg            int32   `json:"azimuth_deg, omitempty"`
	elevationDeg          int32   `json:"elevation_deg, omitempty"`
	dataDump              string  `json:"data_dump, omitempty"`
}

func (g *GPSData) ExportToCSV() (int64, error) {
	return 0, nil
}

func (g *GPSData) ExtractObsPointsAndSat(db *sql.DB) error {

	// get gps observation points and sat data pairings
	qryStmt := `SELECT base_id, related_id FROM gps_observation_points_sat_data`

	results, err := db.Query(qryStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, qryStmt)
		return err
	}

	data := []int32{0, 0}
	// extract the base and related ids for gps observations and sat data
	for results.Next() {
		var baseid, relatedid int32
		if err := results.Scan(&baseid, &relatedid); err != nil {
			log.Fatal(err)
		}

		data[0] = baseid
		data[1] = relatedid
		g.gpsSatData = append(g.gpsSatData, data)
	}

	// get satellite observation data
	qryStmt = `SELECT id, local_time, svid, constellation,
       cn0, agc, has_agc, in_fix, sat_time_nanos,
       sat_time_1sigma_nanos, has_carrier_freq,
       pseudorange_rate_mps, pseudorange_rate_1sigma,
       has_ephemeris, azimuth_deg, elevation_deg,
       data_dump
	   FROM sat_data`

	results, err = db.Query(qryStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, qryStmt)
		return err
	}

	for results.Next() {
		var sat SatObs
		if err := results.Scan(
			&sat.id,
			&sat.localtime,
			&sat.svid,
			&sat.constellation,
			&sat.cn0,
			&sat.agc,
			&sat.hasAgc,
			&sat.infix,
			&sat.satTimeNanos,
			&sat.satTime1signmaNanos,
			&sat.hasCarrierFreq,
			&sat.psuedorangeRateMps,
			&sat.psuedorangeRate1Sigma,
			&sat.hasEphemeris,
			&sat.azimuthDeg,
			&sat.elevationDeg,
			&sat.dataDump); err != nil {
			log.Fatal(err)
		}

		g.satobs = append(g.satobs, sat)
	}

	return nil
}
