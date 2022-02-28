// Test the iso3166 functions.
package translate_test

import (
	"testing"

	xlns "github.com/napcatstudio/translate"
)

func TestIso3166ForCountry(t *testing.T) {
	var tests = []struct {
		country, iso3166 string
	}{
		{"United States of America", "US"},
		{"United Kingdom of Great Britain and Northern Ireland", "GB"},
		{"Frigdia", ""},
	}
	for _, test := range tests {
		iso3166 := xlns.Iso3166ForCountry(test.country)
		if iso3166 != test.iso3166 {
			t.Errorf("Expected ISO-3166 code %s for country %s got %s",
				test.iso3166, test.country, iso3166)
		}
	}
}

func TestCountryForIso3166(t *testing.T) {
	var tests = []struct {
		iso3166, country string
	}{
		{"US", "United States of America"},
		{"gb", "United Kingdom of Great Britain and Northern Ireland"},
		{"fg", ""},
	}
	for _, test := range tests {
		country := xlns.CountryForIso3166(test.iso3166)
		if country != test.country {
			t.Errorf("Expected country %s for ISO-3166 code %s got %s",
				test.country, test.iso3166, country)
		}
	}
}
