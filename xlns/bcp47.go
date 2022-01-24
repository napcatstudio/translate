// bcp47 contains:
//   information correlating ISO 639 and ISO 3166 codes,
//   Google BCP47 places.
//
// see http://www.rfc-editor.org/rfc/bcp/bcp47.txt
package xlns

import (
	"fmt"
	"strings"
)

// GoogleBCP47 is a list of BCP-47 places supported by Google.
var GoogleBCP47 = []string{
	"af",
	"am",
	"ar",
	"hy-AM",
	"az-AZ",
	"eu-ES",
	"be",
	"bn-BD",
	"bg",
	"my-MM",
	"ca",
	"zh-CN",
	"zh-TW",
	"zh-HK",
	"hr",
	"cs-CZ",
	"da-DK",
	"nl-NL",
	"en-US", // Not checked in all tools.
	"en-AU",
	"en-GB",
	"en-IN",
	"en-CA",
	"en-ZA",
	"en-SG",
	"et",
	//"fil", Filipino - not ISO-639 2
	"fi-FI",
	"fr-FR",
	"fr-CA",
	"gl-ES",
	"ka-GE",
	"de-DE",
	"el-GR",
	//"iw-IL", iw is not ISO-639 - possible old code for Hebrew (he)?
	"hi-IN",
	"hu-HU",
	"is-IS",
	"id",
	"it-IT",
	"ja-JP",
	"kn-IN",
	"km-KH",
	"ko-KR",
	"ky-KG",
	"lo-LA",
	"lv",
	"lt",
	"mk-MK",
	"ms",
	"ml-IN",
	"mr-IN",
	"mn-MN",
	"ne-NP",
	"no-NO",
	"fa",
	"pl-PL",
	"pt-BR",
	"pt-PT",
	"ro",
	"rm",
	"ru-RU",
	"sr",
	"si-LK",
	"sk",
	"sl",
	"es-419", // Latin America special code
	"es-ES",
	"es-US",
	"sw",
	"sv-SE",
	"ta-IN",
	"te-IN",
	"th",
	"tr-TR",
	"uk",
	"vi",
	"zu",
}

// TranslateableGoogleLocales returns a list of BCP-47 locales we have
// languages for.  It excludes the defLang locale.
func TranslateableGoogleLocales(wordsDir, defLang string) ([]string, error) {
	var translateable []string
	for _, bcp47 := range GoogleBCP47 {
		if bcp47 == defLang {
			continue
		}
		iso639 := Iso639FromBcp47(bcp47)
		has, err := WordsHasLanguage(wordsDir, iso639)
		if err != nil {
			return nil, fmt.Errorf("Is %s a proper words directory?  Error %v",
				wordsDir, err)
		}
		if has {
			translateable = append(translateable, bcp47)
		}
	}
	return translateable, nil
}

// UntranslateableGoogleLocales returns a list of BCP-47 locales we don't
// have languages for.  It excludes the defLang locale.
func UntranslateableGoogleLocales(wordsDir, defLang string) ([]string, error) {
	var un []string
	for _, bcp47 := range GoogleBCP47 {
		if bcp47 == defLang {
			continue
		}
		iso639 := Iso639FromBcp47(bcp47)
		has, err := WordsHasLanguage(wordsDir, iso639)
		if err != nil {
			return nil, fmt.Errorf("Is %s a proper words directory?  Error %v",
				wordsDir, err)
		}
		if !has {
			un = append(un, bcp47)
		}
	}
	return un, nil
}

// GoogleLocaleForLang trys to find a Google supported locale for the given
// language.
func GoogleLocaleForLang(lang string) (string, error) {
	trialBcp47 := fmt.Sprintf("%s-%s", lang, strings.ToUpper(lang))
	for _, bcp47 := range GoogleBCP47 {
		if trialBcp47 == bcp47 {
			return bcp47, nil
		}
	}
	for _, bcp47 := range GoogleBCP47 {
		if lang == Iso639FromBcp47(bcp47) {
			return bcp47, nil
		}
	}
	return "", fmt.Errorf("%s is not a language in a Google locale", lang)
}

// Iso639FromBcp47 extracts the ISO 639 language code from the BCP-47 code.
func Iso639FromBcp47(bcp47 string) string {
	i := strings.Index(bcp47, "-")
	if i == -1 {
		return bcp47 // It is an ISO 639 code.
	}
	return bcp47[0:i]
}

// Iso3166FromBcp47 extracts the ISO 3166 country code from the BCP-47 code
// if it exists.  Note that in BCP-47 there are special codes that are not ISO
// 3166.
func Iso3166FromBcp47(bcp47 string) string {
	i := strings.Index(bcp47, "-")
	if i == -1 {
		return "" // It has no ISO 3166 part.
	}
	return bcp47[i+1:]
}
