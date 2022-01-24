// words provides information on what languages we have translations
// (words) for.
package xlns

import (
    "fmt"
    "io/ioutil"
    "path"
    "strings"
)

const wordsSuffix = ".words"

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
        if strings.HasSuffix(name, wordsSuffix) {
            if len(name) != 8 {
                return nil, fmt.Errorf(
                        "%s is not a properly formed words file name", name)
            }
            iso639 := strings.TrimSuffix(name, wordsSuffix)
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
    source := path.Join(wordsDir, sourceIso639 + wordsSuffix)
    target := path.Join(wordsDir, targetIso639 + wordsSuffix)
    return XlnsMapFromFiles(source, target)
}
