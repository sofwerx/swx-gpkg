package gpkg

import "time"

// observation contains gps observations and satellite data

type GPSData struct {
	sqlfile    string
	gpsSatData [][]int32
	gpsObs     []GPSObservation
}

type GPSObservation struct {
	id                  uint32
	systime             time.Time
	lat                 float64
	lon                 float64
	alt                 float64
	provider            string
	gpsTime             uint64
	fixedSatCount       int8
	hasRadialAccuracy   bool
	hasVerticalAccuracy bool
	radialAccuracy      float64
	verticalAccuracy    float64
	dataDump            string
	speed               float64
	speedAccuracy       float64
}

type SatObservation struct {
	id                    uint32
	localtime             time.Time
	svid                  uint8
	constellation         string
	cn0                   float32
	agc                   float32
	hasAgc                bool
	infix                 bool
	satTimeNanos          time.Time
	satTime1signmaNanos   uint64
	hasCarrierFreq        bool
	psuedorangeRateMps    float64
	psuedorangeRate1Sigma float64
	hasEphemeris          bool
	azimuthDeg            int32
	elevationDeg          int32
	dataDump              string
}

func (g *GPSData) ExportToCSV() (int64, error) {
	return 0, nil
}

func (g *GPSData) ExtractData() error {
	return nil
}
