package gpkg

type GeoData interface {
	getSQLQuery() string
}
