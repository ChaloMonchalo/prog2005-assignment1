package handler

import "net/http"

func ExchangeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Use: /countryinfo/v1/exchange/{two_letter_country_code} e.g., /countryinfo/v1/exchange/no"))
}