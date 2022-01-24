// iso639 contains information about ISO 639 language and language group
// codes.
package xlns

import (
    "strings"
)

type iso639Code struct {
    language, iso639 string
}

var iso639Codes = []iso639Code {
    {"Abkhazian", "ab"},
    {"Afar", "aa"},
    {"Afrikaans", "af"},
    {"Albanian", "sq"},
    {"Amharic", "am"},
    {"Arabic", "ar"},
    {"Armenian", "hy"},
    {"Assamese", "as"},
    {"Aymara", "ay"},
    {"Azerbaijani", "az"},
    {"Bashkir", "ba"},
    {"Basque", "eu"},

    //{"Bengali (Bangla)", "bn"},
    {"Bengali", "bn"},
    {"Bangla", "bn"},

    {"Bhutani", "dz"},
    {"Bihari", "bh"},
    {"Bislama", "bi"},
    {"Breton", "br"},
    {"Bulgarian", "bg"},
    {"Burmese", "my"},

    //{"Byelorussian (Belarusian)", "be"},
    {"Byelorussian", "be"},
    {"Belarusian", "be"},

    {"Cambodian", "km"},
    {"Catalan", "ca"},

    //{"Chinese (Simplified)", "zh"},
    //{"Chinese (Traditional)", "zh"},
    {"Chinese", "zh"},

    {"Corsican", "co"},
    {"Croatian", "hr"},
    {"Czech", "cs"},
    {"Danish", "da"},
    {"Dutch", "nl"},
    {"English", "en"},
    {"Esperanto", "eo"},
    {"Estonian", "et"},
    {"Faeroese", "fo"},
    {"Farsi", "fa"},
    {"Fiji", "fj"},
    {"Finnish", "fi"},
    {"French", "fr"},
    {"Frisian", "fy"},
    {"Galician", "gl"},

    //{"Gaelic (Scottish)", "gd"},
    //{"Gaelic (Manx)", "gv"},
    {"Gaelic", "gd"},
    {"Scottish", "gd"},
    {"Manx", "gv"},

    {"Georgian", "ka"},
    {"German", "de"},
    {"Greek", "el"},
    {"Greenlandic", "kl"},
    {"Guarani", "gn"},
    {"Gujarati", "gu"},

    {"Haitian", "ht"}, // Added by Graham - aka Creole

    {"Hausa", "ha"},
    {"Hebrew", "he"},
    {"Hindi", "hi"},
    {"Hungarian", "hu"},
    {"Icelandic", "is"},
    {"Indonesian", "id"},
    {"Interlingua", "ia"},
    {"Interlingue", "ie"},
    {"Inuktitut", "iu"},
    {"Inupiak", "ik"},
    {"Irish", "ga"},
    {"Italian", "it"},
    {"Japanese", "ja"},
    {"Javanese", "ja"},
    {"Kannada", "kn"},
    {"Kashmiri", "ks"},
    {"Kazakh", "kk"},

    //{"Kinyarwanda (Ruanda)", "rw"},
    {"Kinyarwanda", "rw"},
    {"Ruanda", "rw"},

    {"Kirghiz", "ky"},

    //{"Kirundi (Rundi)", "rn"},
    {"Kirundi", "rn"},
    {"Rundi", "rn"},

    {"Korean", "ko"},
    {"Kurdish", "ku"},
    {"Laothian", "lo"},
    {"Latin", "la"},

    //{"Latvian (Lettish)", "lv"},
    {"Latvian", "lv"},
    {"Lettish", "lv"},

    //{"Limburgish ( Limburger)", "li"},
    {"Limburgish", "li"},
    {"Limburger", "li"},

    {"Lingala", "ln"},
    {"Lithuanian", "lt"},
    {"Macedonian", "mk"},
    {"Malagasy", "mg"},
    {"Malay", "ms"},
    {"Malayalam", "ml"},
    {"Maltese", "mt"},
    {"Maori", "mi"},
    {"Marathi", "mr"},
    {"Moldavian", "mo"},
    {"Mongolian", "mn"},
    {"Nauru", "na"},
    {"Nepali", "ne"},
    {"Norwegian", "no"},
    {"Occitan", "oc"},
    {"Oriya", "or"},

    //{"Oromo (Afan, Galla)", "om"},
    {"Oromo", "om"},
    {"Afan", "om"},
    {"Galla", "om"},

    //{"Pashto (Pushto)", "ps"},
    {"Pashto", "ps"},
    {"Pushto", "ps"},

    {"Polish", "pl"},
    {"Portuguese", "pt"},
    {"Punjabi", "pa"},
    {"Quechua", "qu"},
    {"Rhaeto-Romance", "rm"},
    {"Romanian", "ro"},
    {"Russian", "ru"},
    {"Samoan", "sm"},
    {"Sangro", "sg"},
    {"Sanskrit", "sa"},
    {"Serbian", "sr"},
    {"Serbo-Croatian", "sh"},
    {"Sesotho", "st"},
    {"Setswana", "tn"},
    {"Shona", "sn"},
    {"Sindhi", "sd"},
    {"Sinhalese", "si"},
    {"Siswati", "ss"},
    {"Slovak", "sk"},
    {"Slovenian", "sl"},
    {"Somali", "so"},
    {"Spanish", "es"},
    {"Sundanese", "su"},

    //{"Swahili (Kiswahili)", "sw"},
    {"Swahili", "sw"},
    {"Kiswahili", "sw"},

    {"Swedish", "sv"},
    {"Tagalog", "tl"},
    {"Tajik", "tg"},
    {"Tamil", "ta"},
    {"Tatar", "tt"},
    {"Telugu", "te"},
    {"Thai", "th"},
    {"Tibetan", "bo"},
    {"Tigrinya", "ti"},
    {"Tonga", "to"},
    {"Tsonga", "ts"},
    {"Turkish", "tr"},
    {"Turkmen", "tk"},
    {"Twi", "tw"},
    {"Uighur", "ug"},
    {"Ukrainian", "uk"},
    {"Urdu", "ur"},
    {"Uzbek", "uz"},
    {"Vietnamese", "vi"},
    {"Volapük", "vo"},
    {"Welsh", "cy"},
    {"Wolof", "wo"},
    {"Xhosa", "xh"},
    {"Yiddish", "yi"},
    {"Yoruba", "yo"},
    {"Zulu", "zu"},
}

// Iso639ForLanguage returns the ISO-639 code for a language.  Will return
// "" if the language is not found.
func Iso639ForLanguage(language string) string {
    lower := strings.ToLower(language)
    for _, lc := range iso639Codes {
        if lower == strings.ToLower(lc.language) {
            return lc.iso639
        }
    }
    return ""
}

// LanguageForIso639 return the first applicable language name for the
// given ISO 639 code (some languages have more than one name).  Will
// return "" if the code is not present.
func LanguageForIso639(iso639 string) string {
    lower := strings.ToLower(iso639)
    for _, lc := range iso639Codes {
        if lower == lc.iso639 {
            return lc.language
        }
    }
    return ""
}
