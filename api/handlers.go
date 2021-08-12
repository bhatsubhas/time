package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func getCurrentTime(rw http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	timeZones := strings.Split(r.URL.Query().Get("tz"), ",")

	if len(timeZones) <= 1 {
		currentTime, err := getTime(timeZones[0])
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			fmt.Fprint(rw, fmt.Sprintf("invalid timezone %s", timeZones[0]))
			return
		}
		response["current_time"] = currentTime
	} else {
		for _, timeZone := range timeZones {
			currentTime, err := getTime(timeZone)
			if err != nil {
				rw.WriteHeader(http.StatusNotFound)
				fmt.Fprint(rw, fmt.Sprintf("invalid timezone %s", timeZone))
				return
			}
			response[timeZone] = currentTime
		}

	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(response)
}

func getTime(zone string) (string, error) {
	loc, err := time.LoadLocation(zone)
	if err != nil {
		return "", err
	}
	return time.Now().In(loc).String(), nil
}
