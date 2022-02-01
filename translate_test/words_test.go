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
