// words provides information on what languages we have translations
// (words) for.
package xlns

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const WORDS_SUFFIX = ".words"

// WordsLanguages returns a list of ISO-639 codes from the available words
// files.
func WordsLanguages(wordsDir string) ([]string, error) {
	fis, err := ioutil.ReadDir(wordsDir)
	if err != nil {
		return nil, err
	}
	langs := make([]string, 0, len(fis))
	for _, fi := range fis {
		name := fi.Name()
		if strings.HasSuffix(name, WORDS_SUFFIX) {
			if len(name) != 8 {
				return nil, fmt.Errorf(
					"%s is not a properly formed words file name", name)
			}
			iso639 := strings.TrimSuffix(name, WORDS_SUFFIX)
			if "" == LanguageForIso639(iso639) {
				return nil, fmt.Errorf(
					"%s is not a ISO-639 code (%s %s)",
					iso639, wordsDir, name)
			}
			langs = append(langs, iso639)
		}
	}
	return langs, nil
}

// WordsHasLanguage returns whether or not we have a translation words file
// for the given ISO-639 code.
func WordsHasLanguage(wordsDir, iso639 string) (bool, error) {
	iso639s, err := WordsLanguages(wordsDir)
	if err != nil {
		return false, err
	}
	lower := strings.ToLower(iso639)
	for _, wordsIso639 := range iso639s {
		if lower == wordsIso639 {
			return true, nil
		}
	}
	if "" == LanguageForIso639(iso639) {
		return false, fmt.Errorf("%s is not a ISO-639 code", iso639)
	}
	return false, nil
}

// WordsXlnsMap creates an XlnsMap object for the given languages.
func WordsXlnsMap(wordsDir, sourceIso639, targetIso639 string) (XlnsMap, error) {
	source := wordsFilename(wordsDir, sourceIso639)
	target := wordsFilename(wordsDir, targetIso639)
	return XlnsMapFromFiles(source, target)
}

// WordsGetWords returns the words for the given language.
func WordsGetWords(wordsDir, iso639 string) ([]string, error) {
	filename := wordsFilename(wordsDir, iso639)
	r, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening %s got %v", filename, err)
	}
	defer r.Close()
	scanner := bufio.NewScanner(r)
	var ss []string
	for scanner.Scan() {
		ss = append(ss, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading %s got %v", filename, err)
	}
	return ss, nil
}

// WordsWriteWords writes the words for a language.
func WordsWriteWords(wordsDir, iso639 string, words []string) error {
	filename := wordsFilename(wordsDir, iso639)
	w, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("creating %s got %v", filename, err)
	}
	defer w.Close()
	for _, word := range words {
		_, err := fmt.Fprintf(w, "%s\n", word)
		if err != nil {
			return fmt.Errorf("writing %s got %v", filename, err)
		}
	}
	return nil
}

func wordsFilename(wordsDir, iso639 string) string {
	return path.Join(wordsDir, iso639+WORDS_SUFFIX)
}
