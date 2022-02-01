// Test the languages for countries functions.
package translate_test

import (
	"testing"

	xlns "github.com/napcatstudio/translate"
)

func TestHasLanguagesForCountry(t *testing.T) {
	var tests = []struct {
		country string
		has     bool
	}{
		{"United States", true},
		{"United Kingdom", true},
		{"Frigdia", false},
	}
	for _, test := range tests {
		has := xlns.HasLanguagesForCountry(test.country)
		if has != test.has {
			t.Errorf("Expected %v for country %s", test.has, test.country)
		}
	}
}

func TestCountryHasLanguage(t *testing.T) {
	var tests = []struct {
		country, language string
		has               bool
	}{
		{"United States", "English", true},
		{"United Kingdom", "English", true},
		{"United States", "French", false},
		{"Brunei", "Malay", true},
		{"Brunei", "English", true},
		{"Frigdia", "Frigidian", false},
	}
	for _, test := range tests {
		has := xlns.CountryHasLanguage(test.country, test.language)
		if has != test.has {
			t.Errorf("Expected %v for country %s and language %s",
				test.has, test.country, test.language)
		}
	}
}

func TestLanguageForCountry(t *testing.T) {
	var tests = []struct {
		country, language string
	}{
		{"United States", "English"},
		{"United Kingdom", "English"},
		{"Frigdia", ""},
	}
	for _, test := range tests {
		lang := xlns.LanguageForCountry(test.country)
		if lang != test.language {
			t.Errorf("Expected %s for country %s got %s",
				test.language, test.country, lang)
		}
	}
}

func TestHasColonialLanguage(t *testing.T) {
	var tests = []struct {
		country string
		has     bool
	}{
		{"United States", true},
		{"United Kingdom", true},
		{"Andorra", true},
		{"Angola", true},
		{"Australia", true},
		{"Estonia", false},
		{"Frigdia", false},
	}
	for _, test := range tests {
		has := xlns.HasColonialLanguage(test.country)
		if has != test.has {
			t.Errorf("Expected %v for country %s got %v",
				test.has, test.country, has)
		}
	}
}
