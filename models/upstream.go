package models

// Used for /info/{code}
type RestCountriesCountry struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`

	Continents []string          `json:"continents"`
	Population int64             `json:"population"`
	Area       float64           `json:"area"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`

	Flags struct {
		PNG string `json:"png"`
	} `json:"flags"`

	Capital []string `json:"capital"`
}

// Used for /exchange/{code} (minimal fields)
type RestCountriesExchangeBase struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`

	Borders []string `json:"borders"`

	Currencies map[string]struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
}

// Minimal upstream model to extract currencies for neighbours
type RestCountriesCurrencyOnly struct {
	Currencies map[string]struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
}