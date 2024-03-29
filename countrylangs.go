package translate

import (
    "strings"
)

type countryLangs struct {
    country, langs string
}

// countryLangs contains an incomplete list of countries and languages.
// from http://www.infoplease.com/ipa/A0855611.html
var countryLangs_ = []countryLangs {

    {"Afghanistan", "Dari Persian, Pashtu (both official), other Turkic and minor languages"},
    {"Albania", "Albanian (Tosk is the official dialect), Greek"},
    {"Algeria", "Arabic (official), French, Berber dialects"},
    {"Andorra", "Catalán (official), French, Castilian, Portuguese"},
    {"Angola", "Portuguese (official), Bantu and other African languages"},
    {"Antigua and Barbuda", "English (official), local dialects"},

    {"Aruba", "Dutch"}, // from Wikipedia

    {"Argentina", "Spanish (official), English, Italian, German, French"},
    {"Armenia", "Armenian 98%, Yezidi, Russian"},
    {"Australia", "English 79%, native and other languages"},
    {"Austria", "German (official nationwide); Slovene, Croatian, Hungarian (each official in one region)"},
    {"Azerbaijan", "Azerbaijani Turkic 89%, Russian 3%, Armenian 2%, other 6% (1995 est.)"},
    {"Bahamas", "English (official), Creole (among Haitian immigrants)"},
    {"Bahrain", "Arabic, English, Farsi, Urdu"},
    {"Bangladesh", "Bangla (official), English"},
    {"Barbados", "English"},

    //{"Belarus", "Belorussian (White Russian), Russian, other"},
    {"Belarus", "Belarusian, Russian"},

    {"Belgium", "Dutch (Flemish) 60%, French 40%, German less than 1% (all official)"},
    {"Belize", "English (official), Spanish, Mayan, Garifuna (Carib), Creole"},
    {"Benin", "French (official), Fon, Yoruba, tribal languages"},
    {"Bhutan", "Dzongkha (official), Tibetan dialects (among Bhotes), Nepalese dialects (among Nepalese)"},
    {"Bolivia", "Spanish, Quechua, Aymara (all official)"},

    // use Croatian
    //{"Bosnia and Herzegovina", "Bosnian, Croatian, Serbian"},
    {"Bosnia and Herzegovina", "Croatian"},

    {"Botswana", "English 2% (official), Setswana 78%, Kalanga 8%, Sekgalagadi 3%, other (2001)"},
    {"Brazil", "Portuguese (official), Spanish, English, French"},
    {"Brunei", "Malay (official), English, Chinese"},
    {"Bulgaria", "Bulgarian 85%, Turkish 10%, Roma 4%"},
    {"Burkina Faso", "French (official); native African (Sudanic) languages 90%"},
    {"Burundi", "Kirundi and French (official), Swahili"},

    //{"Cambodia", "Khmer 95% (official), French, English"},
    {"Cambodia", "Cambodian, French, English"},

    {"Cameroon", "French, English (both official); 24 major African language groups"},
    {"Canada", "English 59.3%, French 23.2% (both official); other 17.5%"},
    {"Cape Verde", "Portuguese, Criuolo"},
    {"Central African Republic", "French (official), Sangho (lingua franca, national), tribal languages"},
    {"Chad", "French, Arabic (both official); Sara; more than 120 languages and dialects"},
    {"Chile", "Spanish"},

    //{"China", "Standard Chinese (Mandarin/Putonghua), Yue (Cantonese), Wu (Shanghaiese), Minbei (Fuzhou), Minnan (Hokkien-Taiwanese), Xiang, Gan, Hakka dialects, minority languages"},
    {"China", "Chinese"},

    {"Colombia", "Spanish"},
    {"Comoros", "Arabic and French (both official), Shikomoro (Swahili/Arabic blend)"},
    {"Congo, Democratic Republic of the", "French (official), Lingala, Kingwana, Kikongo, Tshiluba"},
    {"Congo, Republic of", "French (official), Lingala, Monokutuba, Kikongo, many local languages and dialects"},
    {"Costa Rica", "Spanish (official), English"},

    // Two spellings
    {"Côte d'Ivoire", "French (official) and African languages (Dioula esp.)"},
    {"Cote d'Ivoire", "French (official) and African languages (Dioula esp.)"},


    {"Croatia", "Croatian 96% (official), other 4% (including Italian, Hungarian, Czech, Slovak, German)"},
    {"Cuba", "Spanish"},
    {"Cyprus", "Greek, Turkish (both official); English"},
    {"Czech Republic", "Czech"},
    {"Denmark", "Danish, Faroese, Greenlandic (Inuit dialect), German; English is the predominant second language"},
    {"Djibouti", "French and Arabic (both official), Somali, Afar"},
    {"Dominica", "English (official) and French patois"},
    {"Dominican Republic", "Spanish"},
    {"East Timor", "Tetum, Portuguese (official); Bahasa Indonesia, English; other indigenous languages, including Tetum, Galole, Mambae, and Kemak"},
    {"Ecuador", "Spanish (official), Quechua, other Amerindian languages"},
    {"Egypt", "Arabic (official), English and French widely understood by educated classes"},
    {"El Salvador", "Spanish, Nahua (among some Amerindians)"},
    {"Equatorial Guinea", "Spanish, French (both official); pidgin English, Fang, Bubi, Ibo"},
    {"Eritrea", "Afar, Arabic, Tigre and Kunama, Tigrinya, other Cushitic languages"},
    {"Estonia", "Estonian 67% (official), Russian 30%, other (2000)"},
    {"Ethiopia", "Amharic, Tigrigna, Orominga, Guaragigna, Somali, Arabic, English, over 70 others"},
    {"Fiji", "English (official), Fijian, Hindustani"},
    {"Finland", "Finnish 92%, Swedish 6% (both official); small Sami- (Lapp) and Russian-speaking minorities"},
    {"France", "French 100%, rapidly declining regional dialects (Provençal, Breton, Alsatian, Corsican, Catalan, Basque, Flemish)"},
    {"Gabon", "French (official), Fang, Myene, Nzebi, Bapounou/Eschira, Bandjabi"},
    {"Gambia", "English (official), Mandinka, Wolof, Fula, other indigenous"},
    {"Georgia", "Georgian 71% (official), Russian 9%, Armenian 7%, Azerbaijani 6%, other 7% (Abkhaz is the official language in Abkhazia)"},
    {"Germany", "German"},
    {"Ghana", "English (official), African languages (including Akan, Moshi-Dagomba, Ewe, and Ga)"},
    {"Greece", "Greek 99% (official), English, French"},
    {"Grenada", "English (official), French patois"},
    {"Guatemala", "Spanish 60%, Amerindian languages 40% (23 officially recognized Amerindian languages, including Quiche, Cakchiquel, Kekchi, Mam, Garifuna, and Xinca)"},
    {"Guinea", "French (official), native tongues (Malinké, Susu, Fulani)"},
    {"Guinea-Bissau", "Portuguese (official), Criolo, African languages"},
    {"Guyana", "English (official), Amerindian dialects, Creole, Hindi, Urdu"},

    //{"Haiti", "Creole and French (both official)"},
    {"Haiti", "French, Creole"},

    {"Honduras", "Spanish (official), Amerindian dialects; English widely spoken in business"},

    // Adding Hong Kong
    {"Hong Kong", "Chinese"},

    //{"Hungary", "Magyar (Hungarian) 94%, other 6%"},
    {"Hungary", "Hungarian"},

    {"Iceland", "Icelandic, English, Nordic languages, German widely spoken"},
    {"India", "Hindi 30%, English, Bengali, Gujarati, Kashmiri, Malayalam, Marathi, Oriya, Punjabi, Tamil, Telugu, Urdu, Kannada, Assamese, Sanskrit, Sindhi (all official); Hindi/Urdu; 1,600+ dialects"},

    //{"Indonesia", "Bahasa Indonesia (official), English, Dutch, Javanese, and more than 580 other languages and dialects"},
    {"Indonesia", "Indonesian, English, Dutch, Javanese"},

    //{"Iran", "Persian and Persian dialects 58%, Turkic and Turkic dialects 26%, Kurdish 9%, Luri 2%, Balochi 1%, Arabic 1%, Turkish 1%, other 2%"},
    {"Iran", "Farsi"},

    {"Iraq", "Arabic (official), Kurdish (official in Kurdish regions), Assyrian, Armenian"},
    {"Ireland", "English, Irish (Gaelic) (both official)"},
    {"Israel", "Hebrew (official), Arabic, English"},
    {"Italy", "Italian (official); German-, French-, and Slovene-speaking minorities"},
    {"Jamaica", "English, Jamaican Creole"},
    {"Japan", "Japanese"},
    {"Jordan", "Arabic (official), English"},

    //{"Kazakhstan", "Kazak (Qazaq, state language) 64%; Russian (official, used in everyday business) 95% (2001 est.)"},
    {"Kazakhstan", "Russian, Kazakh"},

    {"Kenya", "English (official), Swahili (national), and numerous indigenous languages"},
    {"Kiribati", "English (official), I-Kiribati (Gilbertese)"},

    // See North Korea/South Korea also
    {"Korea, North", "Korean"},
    {"Korea, South", "Korean, English widely taught"},
    {"Kosovo", "Albanian (official), Serbian (official), Bosnian, Turkish, Roma"},
    {"Kuwait", "Arabic (official), English"},

    //{"Kyrgyzstan", "Kyrgyz, Russian (both official)"},
    {"Kyrgyzstan", "Kirghiz, Russian"},

    //{"Laos", "Lao (official), French, English, various ethnic languages"},
    {"Laos", "Laothian"},

    {"Latvia", "Latvian 58% (official), Russian 38%, Lithuanian, other (2000)"},
    {"Lebanon", "Arabic (official), French, English, Armenian"},
    {"Lesotho", "English, Sesotho (both official); Zulu, Xhosa"},
    {"Liberia", "English 20% (official), some 20 ethnic-group languages"},
    {"Libya", "Arabic, Italian, and English widely understood in major cities"},
    {"Liechtenstein", "German (official), Alemannic dialect"},
    {"Lithuania", "Lithuanian 82% (official), Russian 8%, Polish 6% (2001)"},

    //{"Luxembourg", "Luxermbourgish (national) French, German (both administrative)"},
    {"Luxembourg", "French, German"},

    {"Macedonia", "Macedonian 67%, Albanian 25% (both official); Turkish 4%, Roma 2%, Serbian 1% (2002)"},
    {"Madagascar", "Malagasy and French (both official)"},
    {"Malawi", "Chichewa 57.2% (official), Chinyanja 12.8%, Chiyao 10.1%, Chitumbuka 9.5%, Chisena 2.7%, Chilomwe 2.4%, Chitonga 1.7%, other 3.6% (1998)"},

    //{"Malaysia", "Bahasa Melayu (Malay, official), English, Chinese dialects (Cantonese, Mandarin, Hokkien, Hakka, Hainan, Foochow), Tamil, Telugu, Malayalam, Panjabi, Thai; several indigenous languages (including Iban, Kadazan) in East Malaysia"},
    {"Malaysia", "Malay, English, Chinese"},

    {"Maldives", "Maldivian Dhivehi (official); English spoken by most government officials"},
    {"Mali", "French (official), Bambara 80%, numerous African languages"},
    {"Malta", "Maltese and English (both official)"},
    {"Marshall Islands", "Marshallese 98% (two major dialects from the Malayo-Polynesian family), English widely spoken as a second language (both official); Japanese"},
    {"Mauritania", "Hassaniya Arabic (official), Pulaar, Soninke, French, Wolof"},
    {"Mauritius", "English less than 1% (official), Creole 81%, Bojpoori 12%, French 3% (2000)"},
    {"Mexico", "Spanish, various Mayan, Nahuatl, and other regional indigenous languages"},
    {"Micronesia", "English (official, common), Chukese, Pohnpeian, Yapase, Kosrean, Ulithian, Woleaian, Nukuoro, Kapingamarangi"},

    //{"Moldova", "Moldovan (official; virtually the same as Romanian), Russian, Gagauz (a Turkish dialect)"},
    {"Moldova", "Moldavian, Russian"},

    {"Monaco", "French (official), English, Italian, Monégasque"},
    {"Mongolia", "Mongolian, 90%; also Turkic and Russian (1999)"},
    {"Montenegro", "Serbian/Montenegrin (Ijekavian dialect—official)"},
    {"Morocco", "Arabic (official), Berber dialects, French often used for business, government, and diplomacy"},
    {"Mozambique", "Portuguese 9% (official; second language of 27%), Emakhuwa 26%, Xichangana 11%, Elomwe 8%, Cisena 7%, Echuwabo 6%, other Mozambican languages 32% (1997)"},
    {"Myanmar", "Burmese, minority languages"},
    {"Namibia", "English 7% (official), Afrikaans is common language of most of the population and of about 60% of the white population, German 32%; indigenous languages: Oshivambo, Herero, Nama"},
    {"Nauru", "Nauruan (official), English"},
    {"Nepal", "Nepali 48% (official), Maithali 12%, Bhojpuri 7%, Tharu 6%, Tamang 5%, others. English spoken by many in government and business (2001)"},
    {"Netherlands", "Dutch, Frisian (both official)"},

    // Added
    {"Netherlands Antilles", "Dutch"},

    {"New Zealand", "English, Maori (both official)"},
    {"Nicaragua", "Spanish 98% (official); English and indigenous languages on Atlantic coast (1995)"},
    {"Niger", "French (official), Hausa, Djerma"},
    {"Nigeria", "English (official), Hausa, Yoruba, Ibo, Fulani, and more than 200 others"},

    // also Korea, North
    {"North Korea", "Korean"},

    //{"Norway", "Bokmål Norwegian, Nynorsk Norwegian (both official); small Sami- and Finnish-speaking minorities (Sami is official in six municipalities)"},
    {"Norway", "Norwegian"},

    {"Oman", "Arabic (official), English, Baluchi, Urdu, Indian dialects"},
    {"Pakistan", "Urdu 8%, English (both official); Punjabi 48%, Sindhi 12%, Siraiki (a Punjabi variant) 10%, Pashtu 8%, Balochi 3%, Hindko 2%, Brahui 1%, Burushaski, and others 8%"},
    {"Palau", "Palauan 64.7%, English 9.4%, Sonsoralese, Tobi, Angaur (each official on some islands), Filipino 13.5%, Chinese 5.7%, Carolinian 1.5%, Japanese 1.5%, other Asian 2.3%, other languages 1.5% (2000)"},
    {"Palestinian State (proposed)", "Arabic, Hebrew, English"},
    {"Panama", "Spanish (official), English 14%, many bilingual"},

    //{"Papua New Guinea", "Tok Pisin (Melanesian Pidgin, the lingua franca), Hiri Motu (in Papua region), English 1%–2%; 715 indigenous languages"},
    {"Papua New Guinea", "English"},

    {"Paraguay", "Spanish, Guaraní (both official)"},
    {"Peru", "Spanish, Quéchua (both official); Aymara; many minor Amazonian languages"},

    //{"Philippines", "Filipino (based on Tagalog), English (both official); eight major dialects: Tagalog, Cebuano, Ilocano, Hiligaynon or Ilonggo, Bicol, Waray, Pampango, and Pangasinense"},
    {"Philippines", "English"},

    {"Poland", "Polish 98% (2002)"},
    {"Portugal", "Portuguese (official), Mirandese (official, but locally used)"},
    {"Qatar", "Arabic (official); English a common second language"},
    {"Romania", "Romanian (official), Hungarian, German"},
    {"Russia", "Russian, others"},
    {"Rwanda", "Kinyarwanda, French, and English (all official); Kiswahili in commercial centers"},
    {"St. Kitts and Nevis", "English"},
    {"St. Lucia", "English (official), French patois"},
    {"St. Vincent and the Grenadines", "English, French patois"},
    {"Samoa", "Samoan, English"},
    {"San Marino", "Italian"},
    {"São Tomé and Príncipe", "Portuguese (official)"},
    {"Saudi Arabia", "Arabic"},
    {"Senegal", "French (official); Wolof, Pulaar, Jola, Mandinka"},
    {"Serbia", "Serbian (official); Romanian, Hungarian, Slovak, and Croatian (all official in Vojvodina); Albanian (official in Kosovo)"},
    {"Seychelles", "Seselwa Creole 92%, English 5%, French (all official) (2002)"},
    {"Sierra Leone", "English (official), Mende (southern vernacular), Temne (northern vernacular), Krio (lingua franca)"},

    //{"Singapore", "Mandarin 35%, English 23%, Malay 14.1%, Hokkien 11.4%, Cantonese 5.7%, Teochew 4.9%, Tamil 3.2%, other Chinese dialects 1.8%, other 0.9% (2000)"},
    {"Singapore", "Chinese"},

    {"Slovakia", "Slovak 84% (official), Hungarian 11%, Roma 2%, Ukrainian 1% (2001)"},
    {"Slovenia", "Slovenian 91%, Serbo-Croatian 5% (2002)"},
    {"Solomon Islands", "English 1%–2% (official), Melanesian pidgin (lingua franca), 120 indigenous languages"},
    {"Somalia", "Somali (official), Arabic, English, Italian"},

    //{"South Africa", "IsiZulu 23.8%, IsiXhosa 17.6%, Afrikaans 13.3%, Sepedi 9.4%, English 8.2%, Setswana 8.2%, Sesotho 7.9%, Xitsonga 4.4%, other 7.2%"},
    {"South Africa", "English, Zulu, Afrikaans"},

    // Added
    {"South Korea", "Korean"},

    {"South Sudan", "English (official), Arabic (includes Juba and Sudanese variants) (official), regional languages include Dinka, Nuer, Bari, Zande, Shilluk"},
    //{"Spain", "Castilian Spanish 74% (official nationwide); Catalan 17%, Galician 7%, Basque 2% (each official regionally)"},
    {"Spain", "Spanish"},

    //{"Sri Lanka", "Sinhala 74% (official and national), Tamil 18% (national), other 8%; English is commonly used in government and spoken competently by about 10%"},
    {"Sri Lanka", "Sinhalese, English"},

    {"Sudan", "Arabic (official), Nubian, Ta Bedawie, diverse dialects of Nilotic, Nilo-Hamitic, Sudanic languages, English"},
    {"Suriname", "Dutch (official), Surinamese (lingua franca), English widely spoken, Hindustani, Javanese"},
    {"Swaziland", "English, siSwati (both official)"},
    {"Sweden", "Swedish, small Sami- and Finnish-speaking minorities"},
    {"Switzerland", "German 64%, French 20%, Italian 7% (all official); Romansch 0.5% (national)"},
    {"Syria", "Arabic (official); Kurdish, Armenian, Aramaic, Circassian widely understood; French, English somewhat understood"},
    {"Taiwan", "Chinese (Mandarin, official), Taiwanese (Min), Hakka dialects"},
    {"Tajikistan", "Tajik (official), Russian widely used in government and business"},
    {"Tanzania", "Swahili, English (both official); Arabic; many local languages"},
    {"Thailand", "Thai (Siamese), English (secondary language of the elite), ethnic and regional dialects"},
    {"Togo", "French (official, commerce); Ewé, Mina (south); Kabyé, Dagomba (north); and many dialects"},
    {"Tonga", "Tongan (an Austronesian language), English"},
    {"Trinidad and Tobago", "English (official), Hindi, French, Spanish, Chinese"},
    {"Tunisia", "Arabic (official, commerce), French (commerce)"},
    {"Turkey", "Turkish (official), Kurdish, Dimli, Azeri, Kabardian"},
    {"Turkmenistan", "Turkmen 72%; Russian 12%; Uzbek 9%, other 7%"},
    {"Tuvalu", "Tuvaluan, English, Samoan, Kiribati (on the island of Nui)"},
    {"Uganda", "English (official), Ganda or Luganda, other Niger-Congo languages, Nilo-Saharan languages, Swahili, Arabic"},
    {"Ukraine", "Ukrainian 67%, Russian 24%, Romanian, Polish, Hungarian"},
    {"United Arab Emirates", "Arabic (official), Persian, English, Hindi, Urdu"},
    {"United Kingdom", "English, Welsh, Scots Gaelic"},
    {"United States", "English 82%, Spanish 11% (2000)"},
    {"Uruguay", "Spanish, Portunol, or Brazilero"},
    {"Uzbekistan", "Uzbek 74.3%, Russian 14.2%, Tajik 4.4%, other 7.1%"},
    {"Vanuatu", "Bislama 23% (a Melanesian pidgin English), English 2%, French 1% (all 3 official); more than 100 local languages 73%"},
    {"Vatican City (Holy See)", "Italian, Latin, French, various other languages"},
    {"Venezuela", "Spanish (official), numerous indigenous dialects"},
    {"Vietnam", "Vietnamese (official); English (increasingly favored as a second language); some French, Chinese, Khmer; mountain area languages (Mon-Khmer and Malayo-Polynesian)"},
    {"Western Sahara (proposed state)", "Hassaniya Arabic, Moroccan Arabic"},
    {"Yemen", "Arabic"},
    {"Zambia", "English (official); major vernaculars: Bemba, Kaonda, Lozi, Lunda, Luvale, Nyanja, Tonga; about 70 other indigenous languages"},
    {"Zimbabwe", "English (official), Shona, Ndebele (Sindebele), numerous minor tribal dialects"},

    // Special for Google "Rest of the World"
    {"Rest of the World", "English"},
}

func findCountryLangs(country string) *countryLangs {
    lower := strings.ToLower(country)
    for i, cl := range countryLangs_ {
        if lower == strings.ToLower(cl.country) {
            return &countryLangs_[i]
        }
    }
    return nil
}

// HasLanguagesForCountry tells whether we have language information for
// the country.
func HasLanguagesForCountry(country string) bool {
    cl := findCountryLangs(country)
    if cl != nil {
        return true
    }
    return false
}

// CountryHasLanguage tells whether a language is spoken in the given
// country.
func CountryHasLanguage(country, language string) bool {
    cl := findCountryLangs(country)
    if cl == nil {
        return false
    }
    lowerLangs := strings.ToLower(cl.langs)
    return strings.Contains(lowerLangs, strings.ToLower(language))
}

func langForCountryFieldFunc(r rune) bool {
    return strings.ContainsRune(" ,;:", r)
}

// LanguageForCountry returns the primary language for a country.  Returns
// "" if unknown.
func LanguageForCountry(country string) string {
    cl := findCountryLangs(country)
    if cl == nil {
        return ""
    }
    langs := strings.FieldsFunc(cl.langs, langForCountryFieldFunc)
    return langs[0]
}

var colonialLanguages = []string {
    "Dutch",
    "English",
    "German",
    "French",
    "Portuguese",
    "Spanish",
}
// HasLanguagesForCountry returns whether or not a country has English,
// French, German, Dutch, or Spanish as a language.
func HasColonialLanguage(country string) bool {
    //NOTE: This function is very inefficient (but not used much).
    for _, language := range(colonialLanguages) {
        if CountryHasLanguage(country, language) {
            return true
        }
    }
    return false
}

// List of Latin America with non-Spanish countries/places commented out.
var es419 = []string {
    "Argentina",
    "Bolivia",
    //"Brazil", Portugese
    "Chile",
    "Colombia",
    "Costa Rica",
    "Cuba",
    "Dominican Republic",
    "Ecuador",
    "El Salvador",
    //"French Guiana", TODO: ?
    //"Guadeloupe", TODO ?
    "Guatemala",
    //"Haiti", French
    "Honduras",
    //"Martinque", TODO ?
    "Mexico",
    "Nicaragua",
    "Panama",
    "Paraguay",
    "Peru",
    //"Puerto Rico", America - not country.
    //"Saint Barthelemy" TODO ?
    //"Saint Martin", TODO ?
    "Uruguay",
    "Venezuela",
}

// IsEs419 returns true if the given country is part of Spanish speaking
// Latin America.  This helps with BCP-47 code es-419.
func IsEs419(country string) bool {
    lower := strings.ToLower(country)
    for _, c := range es419 {
        if lower == strings.ToLower(c) {
            return true
        }
    }
    return false
}

