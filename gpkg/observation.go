package gpkg

import (
	"reflect"
	"time"
)

type Observation struct {
	Id                  int32
	Systime             time.Time
	Lat                 float64
	Lon                 float64
	Alt                 float64
	Provider            string
	GpsTime             int64
	Fixsatcount         int8
	Hasradialaccruacy   bool
	Hasverticalaccuracy bool
	Radialaccuracy      float64
	Verticalaccuracy    float64
	Obsdatadump         string
	Speed               float64
	SpeedAccuracy       float64
	SatId               int32
	SatLocalTime        int64
	SatSvid             int8
	SatConstellation    string
	SatCn0              float64
	SatAgc              float64
	SatHasAgc           bool
	SatInFix            bool
	SatTimeNanos        float64
	SatTime1SigmaNanos  float64
	SatHasCarrier       bool
	SatPseudoRateMps    float64
	SatPseudoRate1Sigma float64
	SatHasEphemeris     bool
	SatAzimuth          float64
	SatElevation        float64
	SatDataDump         string
}

func (o *Observation) getInterfacePtrs() []interface{} {

	v := reflect.Indirect(reflect.ValueOf(o))
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Addr().Interface()
	}

	return values
}

func (o *Observation) getSQLQuery() string {
	return `
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
           	ORDER BY G.base_id
			`
}
