package provider

import (
	"encoding/csv"
	"log"
	"strconv"
	"strings"

	"github.com/lafin/tallinn-transport/rest"
)

type TallinnTransportResponse struct {
	VehicleType   int
	LineNumber    string
	Latitude      int
	Longitude     int
	Speed         int
	Sign          int
	VehicleNumber int
}

func parseTallinnTransportResponse(response []byte) ([]Transport, error) {
	r := csv.NewReader(strings.NewReader(string(response)))
	records, err := r.ReadAll()
	if err != nil {
		log.Printf("[ERROR] parse response, %s", err)
		return nil, err
	}

	var items []Transport
	for _, record := range records {
		item := Transport{}
		if record[0] != "" {
			value, err := strconv.Atoi(record[0])
			if err != nil {
				log.Printf("[ERROR] parse vehicleType, %s", err)
				return nil, err
			}
			item.VehicleType = value
		}
		if record[1] != "" {
			item.LineNumber = record[1]
		}
		if record[2] != "" {
			value, err := strconv.ParseFloat(record[2], 32)
			if err != nil {
				log.Printf("[ERROR] parse latitude, %s", err)
				return nil, err
			}
			item.Latitude = float32(value)
		}
		if record[3] != "" {
			value, err := strconv.ParseFloat(record[3], 32)
			if err != nil {
				log.Printf("[ERROR] parse longitude, %s", err)
				return nil, err
			}
			item.Longitude = float32(value)
		}
		items = append(items, item)
	}
	return items, nil
}

func GetTallinnTransport() ([]Transport, error) {
	response, err := rest.Get("https://transport.tallinn.ee/gps.txt")
	if err != nil {
		return nil, err
	}
	res, err := parseTallinnTransportResponse(response)
	if err != nil {
		return nil, err
	}
	return res, nil
}
