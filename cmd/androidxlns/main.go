// Translate an Android Studio strings file to other languages.
package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"translate/xlns"
)

func usage(why string) {
	text := `androidResDir baseLang wordsDir

    Reads an Android strings resource file from the values directory and
    creates values-XX directories, where XX is an ISO-639 two letter language
    code, containing a translation, from baseLang, to each language found in
    wordsDir.

Where:
    androidResDir Is a directory containing the Android resources including the
        values directory.
    baseLang Is the base language to use for translation.
    wordsDir Is the directory containing the meaning ordered words files.`
	fmt.Fprintf(
		os.Stderr,
		"ERROR: %s.\nUSAGE:\n    %s %s",
		why,
		os.Args[0],
		text)
	os.Exit(-1)
}

// isDir returns true if the given path name is a directory.
func isDir(name string) bool {
	fi, err := os.Stat(name)
	return err == nil && fi.IsDir()
}

// isFile returns true if the given path name is a file.
func isFile(name string) bool {
	fi, err := os.Stat(name)
	return err == nil && !fi.IsDir()
}

func main() {
	if len(os.Args) != 4 {
		usage("Wrong number of arguments")
	}
	androidResDir := os.Args[1]
	baseLang := strings.ToLower(os.Args[2])
	wordsDir := os.Args[3]

	// Check arguments.
	valuesDir := path.Join(androidResDir, "values")
	if !isDir(valuesDir) {
		usage(fmt.Sprintf("%s is not a directory", valuesDir))
	}
	stringsXml := path.Join(valuesDir, "strings.xml")
	if !isFile(stringsXml) {
		usage(fmt.Sprintf("%s not found", stringsXml))
	}
	lang := xlns.LanguageForIso639(baseLang)
	if lang == "" {
		usage(fmt.Sprintf("%s is not an ISO-639 language", baseLang))
	}
	if !isDir(wordsDir) {
		usage(fmt.Sprintf("%s is not a directory", wordsDir))
	}
	baseWordsFile := path.Join(wordsDir, baseLang+".words")
	if !isFile(baseWordsFile) {
		usage(fmt.Sprintf("%s not found", baseWordsFile))
	}

	// Read strings file.
	rss, err := xlns.ReadResourceStringsFile(stringsXml)
	if err != nil {
		usage(err.Error())
	}

	// Get the list of languages in wordsDir.
	iso639s, err := xlns.WordsLanguages(wordsDir)
	if err != nil {
		usage(err.Error())
	}
	if 0 == len(iso639s) {
		usage(fmt.Sprintf("%s contains no words files", wordsDir))
	}
	hasBaseLang, _ := xlns.WordsHasLanguage(wordsDir, baseLang)
	if !hasBaseLang {
		usage(fmt.Sprintf("Base language %s not in %s", baseLang, wordsDir))
	}

	// We want to make directories with the same Mode.
	valuesFi, _ := os.Stat(valuesDir)

	missing := make(map[string]string)

	// Translate each language.
	for _, iso639 := range iso639s {
		if iso639 == baseLang {
			continue
		}

		// Does a values-xx directory exist?
		langValDir := path.Join(androidResDir, "values-"+iso639)
		if !isDir(langValDir) {
			fmt.Printf("making directory %s\n", langValDir)
			err = os.Mkdir(langValDir, valuesFi.Mode())
			if err != nil {
				usage(err.Error())
			}
		}

		// Create translator.
		xlnsMap, err := xlns.WordsXlnsMap(wordsDir, baseLang, iso639)
		if err != nil {
			usage(err.Error())
		}

		// Translate.
		langRss := rss.Clone()
		for i, sr := range langRss.Strings {
			translated := xlnsMap.TranslateText(sr.Value)
			if translated == sr.Value {
				s, ok := missing[sr.Value]
				if ok {
					missing[sr.Value] = s + " " + iso639
				} else {
					missing[sr.Value] = iso639
				}
			} else {
				langRss.Strings[i].Value = translated
			}
		}

		// New file.
		langFile := path.Join(langValDir, "strings.xml")
		fmt.Printf("writing %s\n", langFile)
		err = langRss.WriteFile(langFile)
		if err != nil {
			usage(err.Error())
		}
	}

	// Report missing words.
	if len(missing) > 0 {
		fmt.Printf("Missing word summary:\n")
		for word, iso639s := range missing {
			fmt.Printf("%s: %s\n", word, iso639s)
		}
	}
}
