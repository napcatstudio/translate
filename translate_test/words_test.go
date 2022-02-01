// Test the words files.

package translate_test

import (
	"io/ioutil"
	"path"
	"testing"

	xlns "github.com/napcatstudio/translate"
)

const wordsDir = "testdata/words"
const enWordsName = "en.words"

var (
	wordsEnWords = path.Join(wordsDir, enWordsName)
)

func TestWords(t *testing.T) {
	fis, err := ioutil.ReadDir(wordsDir)
	if err != nil {
		t.Fatalf("%s directory problem (%s).", wordsDir, err.Error())
	}
	for _, fi := range fis {
		name := fi.Name()
		if name != enWordsName {
			w2 := path.Join(wordsDir, name)
			_, err = xlns.XlnsMapFromFiles(wordsEnWords, w2)
			if err != nil {
				t.Errorf("en.words to %s failed.", w2)
			}
			_, err = xlns.XlnsMapFromFiles(w2, wordsEnWords)
			if err != nil {
				t.Errorf("%s to en.words failed.", w2)
			}
		}
	}
}

func TestWordsHasLanguage(t *testing.T) {
	var tests = []struct {
		lang string
		has  bool
	}{
		{"en", true},
		{"fil", true},
		{"zh-CN", true},
		{"es-419", true},
		{"fr", true},
		{"xh", false}, // Xhosa
		{"pt", false}, // We have pt-BR and pt-PT but not pt.
	}
	for _, test := range tests {
		has, err := xlns.WordsHasLanguage(wordsDir, test.lang)
		if err != nil {
			t.Fatalf("%s directory problem (%s)", wordsDir, err.Error())
		}
		if has != test.has {
			t.Errorf("Got %t for %s expected %t", has, test.lang, test.has)
		}
	}
}

func TestES419(t *testing.T) {
	words, err := xlns.WordsGetWords(wordsDir, "es-419")
	if err != nil {
		t.Fatalf("es-419 WordsGetWords failed")
	}
	if len(words) == 0 {
		t.Errorf("es-419 has no words")
	}
	xmap, err := xlns.WordsXlnsMap(wordsDir, "en", "es-419")
	if err != nil {
		t.Fatalf("es-419 WordsXlnsMap failed")
	}
	if len(xmap) == 0 {
		t.Errorf("es-419 map has no words")
	}
}

func TestWordsLanguages(t *testing.T) {
	langs, _ := xlns.WordsLanguages(wordsDir)
	enWords, _ := xlns.WordsGetWords(wordsDir, "en")
	for _, lang := range langs {
		if lang == "en" {
			continue
		}
		xlnsMap, _ := xlns.WordsXlnsMap(wordsDir, "en", lang)
		if len(xlnsMap) != len(enWords) {
			t.Fatalf(
				"bad translation map len for %s %d!=%d",
				lang, len(enWords), len(xlnsMap))
		}
		for _, word := range enWords {
			xln, ok := xlnsMap[word]
			if !ok {
				t.Errorf("%s missing translation", lang)
			}
			if xln == "" {
				t.Errorf("%s empty translation", lang)
			}
		}
	}
}
