// bcp47 contains:
//   information correlating ISO 639 and ISO 3166 codes,
//   see http://www.rfc-editor.org/rfc/bcp/bcp47.txt
package translate

import (
	"strings"
)

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
