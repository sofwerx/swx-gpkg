package gpkg

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"
)

type Observation struct {
	Id                  int32     `json:"id"`
	SysTime             time.Time `json:"sys_time"`
	Lat                 float64   `json:"lat"`
	Lon                 float64   `json:"lon"`
	Alt                 float64   `json:"alt"`
	Provider            string    `json:"provider"`
	GpsTime             int64     `json:"gps_time"`
	FixSatCount         int8      `json:"fix_sat_count"`
	HasRadialAccuracy   bool      `json:"has_radial_accuracy"`
	HasVerticalAccuracy bool      `json:"has_vertical_accuracy"`
	RadialAccuracy      float64   `json:"radial_accuracy"`
	VerticalAccuracy    float64   `json:"vertical_accuracy"`
	Speed               float64   `json:"speed"`
	SpeedAccuracy       float64   `json:"speed_accuracy"`
	SatId               int32     `json:"sat_id"`
	SatLocalTime        int64     `json:"sat_local_time"`
	SatVehicleId        int8      `json:"sat_vehicle_id"`
	SatConstellation    string    `json:"sat_constellation"`
	SatCn0              float64   `json:"sat_cn_0"`
	SatAgc              float64   `json:"sat_agc"`
	SatHasAgc           bool      `json:"sat_has_agc"`
	SatInFix            bool      `json:"sat_in_fix"`
	SatTimeNanos        float64   `json:"sat_time_nanos"`
	SatTime1SigmaNanos  float64   `json:"sat_time_1_sigma_nanos"`
	SatHasCarrier       bool      `json:"sat_has_carrier"`
	SatPseudoRateMps    float64   `json:"sat_pseudo_rate_mps"`
	SatPseudoRate1Sigma float64   `json:"sat_pseudo_rate_1_sigma"`
	SatHasEphemeris     bool      `json:"sat_has_ephemeris"`
	SatAzimuth          float64   `json:"sat_azimuth"`
	SatElevation        float64   `json:"sat_elevation"`
}

func (o *Observation) getInterfacePtrs() []interface{} {

	v := reflect.Indirect(reflect.ValueOf(o))
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Addr().Interface()
	}

	return values
}

func (o *Observation) Json() []byte {
	j, err := json.Marshal(o)
	if err != nil {
		log.Fatal("could not marshal json")
	}

	return j
}

func (o *Observation) String() []string {
	var s []string
	v := reflect.Indirect(reflect.ValueOf(o))
	s = make([]string, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface()
		s[i] = fmt.Sprintf("%v", value)
	}

	return s
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
				S.elevation_deg as sat_elevation_deg                               
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
