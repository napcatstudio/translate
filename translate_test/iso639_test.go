// Test the iso639 functions.
package translate_test

import (
	"testing"

	xlns "github.com/napcatstudio/translate"
)

func TestIso639ForLanguage(t *testing.T) {
	var tests = []struct {
		lang, iso639 string
	}{
		{"zulu", "zu"},
		{"English", "en"},
		{"frigian", ""},
	}
	for _, test := range tests {
		iso639 := xlns.Iso639ForLanguage(test.lang)
		if iso639 != test.iso639 {
			t.Errorf("Expected ISO-639 code %s for language %s got %s",
				test.iso639, test.lang, iso639)
		}
	}
}

func TestLanguageForIso639(t *testing.T) {
	var tests = []struct {
		iso639, lang string
	}{
		{"zu", "Zulu"},
		{"en", "English"},
		{"xz", ""},
	}
	for _, test := range tests {
		lang := xlns.LanguageForIso639(test.iso639)
		if lang != test.lang {
			t.Errorf("Expected language %s for ISO-639 %s got %s",
				test.lang, test.iso639, lang)
		}
	}
}
