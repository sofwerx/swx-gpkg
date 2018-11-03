package gpkg

type GeoData interface {
	getSQLQuery() string
	getInterfacePtrs() []interface{}
	String() []string
}
