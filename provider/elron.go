package provider

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/lafin/tallinn-transport/rest"
)

type ElronResponse struct {
	Status int `json:"status"`
	Data   []struct {
		Reis             string `json:"reis"`
		Liin             string `json:"liin"`
		ReisiAlgusAeg    string `json:"reisi_algus_aeg"`
		ReisiLoppAeg     string `json:"reisi_lopp_aeg"`
		Kiirus           string `json:"kiirus"`
		Latitude         string `json:"latitude"`
		Longitude        string `json:"longitude"`
		RongiSuund       string `json:"rongi_suund"`
		ErinevusPlaanist string `json:"erinevus_plaanist"`
		Lisateade        string `json:"lisateade"`
		PohjusTeade      string `json:"pohjus_teade"`
		AvaldaKodulehel  string `json:"avalda_kodulehel"`
		AsukohaUuendus   string `json:"asukoha_uuendus"`
		ReisiStaatus     string `json:"reisi_staatus"`
		ViimanePeatus    string `json:"viimane_peatus"`
	} `json:"data"`
	Src string `json:"src"`
}

func parseElronResponse(response []byte) ([]Transport, error) {
	var data ElronResponse
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}
	var items []Transport
	for _, record := range (&data).Data {
		var item = Transport{
			VehicleType: 4,
			LineNumber:  record.Liin,
		}
		if record.Latitude != "" {
			value, err := strconv.ParseFloat(record.Latitude, 32)
			if err != nil {
				log.Printf("[ERROR] parse latitude, %s", err)
				return nil, err
			}
			item.Latitude = float32(value)
		}
		if record.Longitude != "" {
			value, err := strconv.ParseFloat(record.Longitude, 32)
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

func GetElronTransport() ([]Transport, error) {
	response, err := rest.Get("https://elron.ee/api/v1/map")
	if err != nil {
		return nil, err
	}
	res, err := parseElronResponse(response)
	if err != nil {
		return nil, err
	}
	return res, nil
}
