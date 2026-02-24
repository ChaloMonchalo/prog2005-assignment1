package handler

import "net/http"

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Use: /countryinfo/v1/info/{two_letter_country_code} e.g., /countryinfo/v1/info/no"))
}