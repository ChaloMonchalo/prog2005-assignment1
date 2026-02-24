package models

type CountryInfoResponse struct {
	Name       string            `json:"name"`
	Continents []string          `json:"continents"`
	Population int64             `json:"population"`
	Area       float64           `json:"area"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capital    string            `json:"capital"`
}

type ExchangeResponse struct {
	Country      string             `json:"country"`
	BaseCurrency string             `json:"base-currency"`
	ExchangeRates map[string]float64 `json:"exchange-rates"`
}