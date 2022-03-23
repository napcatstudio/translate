// Common translation tools.
package translate

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type XlnsMap map[string]string

// TranslationMap takes two lists of words ordered by meaning and creates a
// map from the first to the second.
func NewXlnsMap(base1, base2 io.Reader) (XlnsMap, error) {
	b1sc := bufio.NewScanner(base1)
	b2sc := bufio.NewScanner(base2)
	t := make(XlnsMap)
	for b1sc.Scan() {
		read := b2sc.Scan()
		if !read {
			err := b2sc.Err()
			if err == nil {
				err = fmt.Errorf("mismatched word files")
			}
			return nil, err
		}
		t[strings.TrimSpace(b1sc.Text())] = strings.TrimSpace(b2sc.Text())
	}
	if err := b1sc.Err(); err != nil {
		return nil, err
	}
	if err := b2sc.Err(); err != nil {
		return nil, err
	}
	return t, nil
}

// XlnsMapFromFiles creates a XlnsMap from the words files w1 and w2.
func XlnsMapFromFiles(w1, w2 string) (XlnsMap, error) {
	w1r, err := os.Open(w1)
	if err != nil {
		return nil, err
	}
	defer w1r.Close()
	w2r, err := os.Open(w2)
	if err != nil {
		return nil, err
	}
	defer w2r.Close()
	return NewXlnsMap(w1r, w2r)
}

// Translate a string using the translation map.  Use target if there is no
// other tranlation.
func (xm XlnsMap) Translate(source, target string) string {
	// Straight translation.
	newTarget, ok := xm[source]
	if ok {
		return cleanUp(source, newTarget)
	}
	if strings.HasPrefix(source, "$(") && strings.HasSuffix(source, ")") {
		envKey := source[2 : len(source)-1]
		envVal := os.Getenv(envKey)
		if envVal != "" {
			return xm.Translate(envVal, target)
		}
		return target
	}
	// Uppercase version of known word?
	newTarget, ok = xm[strings.ToLower(source)]
	if ok {
		return cleanUp(source, newTarget)
	}
	if IsCamelCase(source) {
		pieces := camelCasePieces(source)
		for ip, p := range pieces {
			pieces[ip] = xm.Translate(p, "")
		}
		result := strings.Join(pieces, "")
		if len(result) != 0 {
			return result
		}
	}
	// Is it a number?
	_, err := strconv.ParseFloat(source, 32)
	if err == nil {
		return source
	}
	return target
}

// TranslateText by word.
func (xm XlnsMap) TranslateText(text string) string {
	result := ""
	word := ""
	for _, rv := range text {
		if unicode.IsLetter(rv) {
			word = word + string(rv)
		} else {
			if word != "" {
				result = result + xm.Translate(word, word)
				word = ""
			}
			result = result + string(rv)
		}
	}
	if word != "" {
		result = result + xm.Translate(word, word)
	}
	return result
}

// TranslateByLine translates source looking for entire lines.  If that
// fails it will translate by word (see TranslateText).
func (xm XlnsMap) TranslateByLine(source string) string {
	lines := strings.Split(source, "\n")
	xlines := make([]string, len(lines), len(lines))
	for il, line := range lines {
		xline, ok := xm[line]
		if !ok {
			xline = xm.TranslateText(line)
		}
		xlines[il] = xline
	}
	return strings.Join(xlines, "\n")
}

// TranslateByLineWithAlternate translates like TranslateByLine but if the
// translation is too long translates alternative text instead.
func (xm XlnsMap) TranslateByLineWithAlternate(source, altSource string, limit int) string {
	xlns := xm.TranslateByLine(source)
	if len(xlns) > limit && altSource != "" {
		xlns = xm.TranslateByLine(altSource)
	}
	return xlns
}

// TranslateWords translates a list of words.  If there is no translation
// for a word it will be skipped.
func (xm XlnsMap) TranslateWords(words []string) []string {
	result := make([]string, 0, len(words))
	for _, word := range words {
		xword := xm.Translate(word, "")
		if xword != "" {
			result = append(result, xword)
		}
	}
	return result
}

// Key returns the "key" words as a sorted slice.
func (xm XlnsMap) Key() []string {
	var key []string
	for k := range xm {
		key = append(key, k)
	}
	sort.Sort(sort.StringSlice(key))
	return key
}

// IsCamelCase return true if the string is camelCase.
func IsCamelCase(s string) bool {
	first := true
	lastLower := false
	for _, rv := range s {
		if !unicode.IsLetter(rv) {
			return false
		}
		upper := unicode.IsUpper(rv)
		if first {
			first = false
		} else if lastLower && upper {
			return true
		}
		lastLower = !upper
	}
	return false
}

func camelCasePieces(s string) []string {
	pieces := make([]string, 0, 2)
	lastLower := false
	start := 0
	for i := 0; i < len(s); {
		rv, nb := utf8.DecodeRuneInString(s[i:])
		isUpper := unicode.IsUpper(rv)
		if isUpper && lastLower && i > 0 {
			pieces = append(pieces, s[start:i])
			start = i
		}
		lastLower = !isUpper
		i += nb
	}
	if start < len(s) {
		pieces = append(pieces, s[start:])
	}
	//fmt.Println(s, ":", pieces)
	return pieces
}

func firstLetterMatchCase(orig, s string) string {
	origRune, _ := utf8.DecodeRuneInString(orig)
	if unicode.IsUpper(origRune) {
		sRune, width := utf8.DecodeRuneInString(s)
		//fmt.Println(orig, s, strings.ToUpper(string(sRune)) + s[width:])
		return strings.ToUpper(string(sRune)) + s[width:]
	}
	return s
}

func cleanUp(orig, s string) string {
	s = firstLetterMatchCase(orig, s)
	if strings.HasSuffix(s, "-") {
		s = s[:len(s)-1]
	}
	return s
}
