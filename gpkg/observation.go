package gpkg

import (
	"database/sql"
	"log"
	"time"
)

type Observation struct {
	id                  int32
	systime             time.Time
	lat                 float64
	lon                 float64
	alt                 float64
	provider            string
	gpsTime             int64
	fixsatcount         int8
	hasradialaccruacy   bool
	hasverticalaccuracy bool
	radialaccuracy      float64
	verticalaccuracy    float64
	obsdatadump         string
	speed               float64
	speedAccuracy       float64
	satId               int32
	satLocalTime        int64
	satSvid             int8
	satConstellation    string
	satCn0              float64
	satAgc              float64
	satHasAgc           bool
	satInFix            bool
	satTimeNanos        float64
	satTime1SigmaNanos  float64
	satHasCarrier       bool
	satPseudoRateMps    float64
	satPseudoRate1Sigma float64
	satHasEphemeris     bool
	satAzimuth          float64
	satElevation        float64
	satDataDump         string
}

type GPSData struct {
	sqlfile string
}

func (g *GPSData) ExportToCSV(db *sql.DB) (int, error) {
	qrystr := `
	SELECT P.id,
       P.SysTime,
       P.Lat,
       P.Lon,
       P.Alt,
       P.Provider,
       P.GPSTime,
       P.FixSatCount,
       P.HasRadialAccuracy,
       P.HasVerticalAccuracy,
       P.RadialAccuracy,
       P.VerticalAccuracy,
       P.data_dump,
       P.Speed,
       P.SpeedAccuracy,
       S.id as sat_id,
       S.local_time as sat_local_time,
       S.svid as sat_svid,
       S.constellation as sat_constellation,
       S.cn0 as sat_cn0,
       S.agc as sat_agc,
       S.has_agc as sat_has_agc,
       S.in_fix as sat_in_fix,
       S.sat_time_nanos,
       S.sat_time_1sigma_nanos,
       S.has_carrier_freq as sat_has_carrier_freq,
       S.pseudorange_rate_mps,
       S.pseudorange_rate_1sigma,
       S.has_ephemeris as sat_has_ephemeris,
       S.azimuth_deg as sat_azimuth_deg,
       S.elevation_deg as sat_elevation_deg,
       S.data_dump as sat_data_dump
	FROM gps_observation_points_sat_data AS G
       	 LEFT JOIN
     	 gps_observation_points AS P
     	 ON G.base_id = P.id
         LEFT JOIN
         sat_data as S
         ON G.related_id = S.id
	ORDER BY G.base_id`

	results, err := db.Query(qrystr)
	if err != nil {
		log.Printf("%q: %s", err, qrystr)
	}

	count := 1

	for results.Next() {
		var obs = Observation{}
		if err := results.Scan(
			&obs.id,
			&obs.systime,
			&obs.lat,
			&obs.lon,
			&obs.alt,
			&obs.provider,
			&obs.gpsTime,
			&obs.fixsatcount,
			&obs.hasradialaccruacy,
			&obs.hasverticalaccuracy,
			&obs.radialaccuracy,
			&obs.verticalaccuracy,
			&obs.obsdatadump,
			&obs.speed,
			&obs.speedAccuracy,
			&obs.satId,
			&obs.satLocalTime,
			&obs.satSvid,
			&obs.satConstellation,
			&obs.satCn0,
			&obs.satAgc,
			&obs.satHasAgc,
			&obs.satInFix,
			&obs.satTimeNanos,
			&obs.satTime1SigmaNanos,
			&obs.satHasCarrier,
			&obs.satPseudoRateMps,
			&obs.satPseudoRate1Sigma,
			&obs.satHasEphemeris,
			&obs.satAzimuth,
			&obs.satElevation,
			&obs.satDataDump); err != nil {
			log.Fatal(err)
		}
		count++
	}

	return count, nil
}

func (g *GPSData) ExtractObsPointsAndSat(db *sql.DB) error {

	return nil
}
