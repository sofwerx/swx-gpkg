package gpkg

import "time"

// referencing Android API's GnssStatus class
// https://developer.android.com/reference/android/location/GnssStatus.html

type Satellite struct {
	Id                      int64     `json:"id"`
	LocalTime               time.Time `json:"local_time"`
	VehicleId               int8      `json:"vehicle_id"`
	Constellation           string    `json:"constellation"`
	CarrierToNoiseRatioDbHz float64   `json:"carrier_to_noise_ratio"`
	AutomaticGainControl    float64   `json:"automatic_gain_control"`
	HasAutomaticGainControl bool      `json:"has_automatic_gain_control"`
	UsedInFix               bool      `json:"used_in_fix"`
	TimeNanoSeconds         float64   `json:"time_nano_seconds"`
	Time1SigmaSeconds       float64   `json:"time_1_sigma_seconds"`
	HasCarrierFrequency     bool      `json:"has_carrier_frequency"`
}
