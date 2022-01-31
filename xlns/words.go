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

// WordsLanguages returns a list of ISO-639 or BCP-47 codes from the
// available words files.
func WordsLanguages(wordsDir string) ([]string, error) {
	fis, err := ioutil.ReadDir(wordsDir)
	if err != nil {
		return nil, err
	}
	langs := make([]string, 0, len(fis))
	for _, fi := range fis {
		name := fi.Name()
		if strings.HasSuffix(name, WORDS_SUFFIX) {
			langs = append(langs, strings.TrimSuffix(name, WORDS_SUFFIX))
		}
	}
	return langs, nil
}

// WordsLanguagesMap returns a map of language to bool for the languages
// in wordsDir.
func WordsLanguagesMap(wordsDir string) (map[string]bool, error) {
	langs, err := WordsLanguages(wordsDir)
	if err != nil {
		return nil, err
	}
	langMap := make(map[string]bool)
	for _, lang := range langs {
		langMap[lang] = true
	}
	return langMap, nil
}

// WordsHasLanguage returns whether or not we have a translation words file
// for the given code.
// This used to check for strict ISO-639 languages and not be case sensitive.
func WordsHasLanguage(wordsDir, lang string) (bool, error) {
	langs, err := WordsLanguages(wordsDir)
	if err != nil {
		return false, err
	}
	for _, wordsLang := range langs {
		if lang == wordsLang {
			return true, nil
		}
	}
	return false, nil
}

// WordsXlnsMap creates an XlnsMap object for the given languages.
func WordsXlnsMap(wordsDir, sourceIso639, targetIso639 string) (XlnsMap, error) {
	source := WordsFilename(wordsDir, sourceIso639)
	target := WordsFilename(wordsDir, targetIso639)
	return XlnsMapFromFiles(source, target)
}

// WordsGetWords returns the words for the given language.
func WordsGetWords(wordsDir, iso639 string) ([]string, error) {
	filename := WordsFilename(wordsDir, iso639)
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
	filename := WordsFilename(wordsDir, iso639)
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

// WordsFilename returns the path of the words file for the given language.
func WordsFilename(wordsDir, iso639 string) string {
	return path.Join(wordsDir, iso639+WORDS_SUFFIX)
}

// WordsMerge adds the words in fromWordsDir to toWordsDir.  If toWordsDir
// has more languages than fromWordsDir it bails.
func WordsMerge(toWordsDir, fromWordsDir string) error {
	toLangs, err := WordsLanguages(toWordsDir)
	if err != nil {
		return err
	}
	fromLangs, err := WordsLanguages(fromWordsDir)
	if err != nil {
		return err
	}
	// Verify toWordsDir contains all needed languages.
	if len(fromLangs) < len(toLangs) {
		return fmt.Errorf("not enough languages for update")
	}
	for _, toLang := range toLangs {
		found := false
		for _, fromLang := range fromLangs {
			if fromLang == toLang {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("fromWordsDir missing %s", toLang)
		}
	}
	// Merge the files.
	xlnss := make(map[string]XlnsMap)
	enToPath := WordsFilename(toWordsDir, "en")
	enFromPath := WordsFilename(fromWordsDir, "en")
	for _, toLang := range fromLangs {
		toPath := WordsFilename(toWordsDir, toLang)
		fromPath := WordsFilename(fromWordsDir, toLang)
		to, err := XlnsMapFromFiles(enToPath, toPath)
		if err != nil {
			return err
		}
		from, err := XlnsMapFromFiles(enFromPath, fromPath)
		if err != nil {
			return err
		}
		// Do the actual merge for this language.
		for en, other := range from {
			to[en] = other
		}
		xlnss[toLang] = to
	}
	// Now that we have the merged xlns maps write them.
	key := xlnss["en"].Key() // Keep them ordered!
	for lang, xlns := range xlnss {
		toPath := WordsFilename(toWordsDir, lang)
		toF, err := os.Create(toPath)
		if err != nil {
			return err
		}
		defer toF.Close()
		for _, keyWord := range key {
			_, err = fmt.Fprintln(toF, xlns[keyWord])
			if err != nil {
				return fmt.Errorf("writing %s got %v", lang, err)
			}
		}
	}
	return nil
}

// WordsCheck does very simplistic verification that a wordsDir is
// consistent.
func WordsCheck(wordsDir string) error {
	// Getting a list of languages will find any we don't know about.
	langs, err := WordsLanguages(wordsDir)
	if err != nil {
		return fmt.Errorf("check got %v", err)
	}
	// Checking each file to the last will make sure that the files have the
	// same number of words (not that the words are good translations).
	last := ""
	for _, lang := range langs {
		current := path.Join(wordsDir, fmt.Sprintf("%s.words", lang))
		if last == "" {
			last = current
			continue
		}
		_, err = XlnsMapFromFiles(last, current)
		if err != nil {
			return fmt.Errorf("check maps got %v", err)
		}
		last = current
	}
	return nil
}
