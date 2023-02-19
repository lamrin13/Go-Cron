package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type FlightPrice struct {
	Success bool
	Data    Data
}

type Data struct {
	Buckets []Bucket
}

type Bucket struct {
	Items []Item
}

type Item struct {
	Price Price
	Legs  []Leg
	Id    string
}

type Leg struct {
	Segments  []Segment
	Departure string
	Arrival   string
}

type Airport struct {
	Name string
}

type Segment struct {
	Origin           Airport
	Destination      Airport
	OperatingCarrier Operator
}

type Operator struct {
	Name string
}

type Price struct {
	Raw       float32
	Formatted string
}

type MailData struct {
	Origins      []string
	Destinations []string
	Airlines     []string
	Stops        int
	Departure    string
	Arrival      string
	Price        string
	PriceRaw     float32
}

const url = "https://app.goflightlabs.com/search-best-flights?access_key=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiI0IiwianRpIjoiZTcyMDBjMTM4MzNjMzg4OGYzMTdlZDU4YzhkYmFjZGU4MjZmYWY5NjQ0ZTk3MmQ0YTc4ODUzYjIzYmI4MDEwZmNhMjljOWQxZTc3NTZlOTIiLCJpYXQiOjE2NzYyMjE2OTIsIm5iZiI6MTY3NjIyMTY5MiwiZXhwIjoxNzA3NzU3NjkyLCJzdWIiOiIyMDA0OSIsInNjb3BlcyI6W119.jnYTr0VBt-yBHEYGUkv5kqShnbCq2cfe8IhUdSlNNxFv_QWkxK8mQ1d_mNLK9rMhTWtsMDWSOZQZXIQDzDgNFg&adults=1&origin=YYZ&destination=AMD&departureDate=2023-11-03&currency=CAD"

func ParseJSON() (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", fmt.Errorf("error while calling API: %s", err)
	}
	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("error while reading response: %s", err)
	}

	var obj FlightPrice

	err = json.Unmarshal(data, &obj)

	if err != nil {
		return "", fmt.Errorf("error while parsing response: %s", err)
	}

	unique := map[string]bool{}

	mailData := []MailData{}
	for _, bucket := range obj.Data.Buckets {
		for _, item := range bucket.Items {
			if _, ok := unique[item.Id]; !ok {
				temp := MailData{}
				processData(item, &temp, unique)
				mailData = append(mailData, temp)
			}
		}
	}
	sort.Slice(mailData, func(i, j int) bool {
		return mailData[i].Price < mailData[j].Price
	})

	return writeFormattedData(mailData), nil
}

func processData(item Item, temp *MailData, unique map[string]bool) {
	for _, leg := range item.Legs {
		for _, v := range leg.Segments {
			temp.Airlines = append(temp.Airlines, v.OperatingCarrier.Name)
			temp.Origins = append(temp.Origins, v.Origin.Name)
			temp.Destinations = append(temp.Destinations, v.Destination.Name)

		}
		temp.Stops = len(leg.Segments) - 1
		temp.Departure = leg.Departure
		temp.Arrival = leg.Arrival

	}
	unique[item.Id] = true
	temp.Price = item.Price.Formatted
	temp.PriceRaw = item.Price.Raw
}

func writeFormattedData(mailData []MailData) string {
	var sb strings.Builder
	for _, v := range mailData {
		sb.WriteString("Price: " + v.Price + "\n")
		sb.WriteString("Departure Time: " + v.Departure + "\n")
		sb.WriteString("Arrival Time: " + v.Arrival + "\n")
		sb.WriteString("Stops: " + strconv.Itoa(v.Stops) + "\n")

		for i := 0; i < len(v.Origins); i++ {
			sb.WriteString("----------------------------\n")
			sb.WriteString("Origins: " + v.Origins[i] + "\n")
			sb.WriteString("Destination: " + v.Destinations[i] + "\n")
			sb.WriteString("Airline: " + v.Airlines[i] + "\n")

		}

		sb.WriteString("**********************************************************************\n")
	}
	return sb.String()
}
