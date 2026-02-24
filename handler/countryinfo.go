package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"assignment-1/models"
)

func CountryInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	code := strings.ToLower(strings.TrimSpace(r.PathValue("code")))
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	up, status, err := fetchCountryByCode(code)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)

		msg := "could not retrieve country info"
		if status == http.StatusNotFound {
			msg = "country code not found"
		} else if status == http.StatusServiceUnavailable {
			msg = "upstream service unavailable"
		} else if status == http.StatusBadGateway {
			msg = "bad response from upstream"
		}

		_ = json.NewEncoder(w).Encode(map[string]string{
			"error": msg,
		})
		return
	}

	// Map upstream → downstream (strict schema)
	out := models.CountryInfoResponse{
		Name:       up.Name.Common,
		Continents: up.Continents,
		Population: up.Population,
		Area:       up.Area,
		Languages:  up.Languages,
		Borders:    up.Borders,
		Flag:       up.Flags.PNG,
		Capital:    "",
	}
	if len(up.Capital) > 0 {
		out.Capital = up.Capital[0]
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(out)
}
