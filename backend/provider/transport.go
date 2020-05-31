// Package provider contains methods for receiving data from providers
package provider

// Transport - the main transport model
type Transport struct {
	VehicleType   int
	LineNumber    string
	Latitude      float32
	Longitude     float32
	VehicleNumber int
}
