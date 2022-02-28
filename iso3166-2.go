// iso3266-2.go
// source: https://en.wikipedia.org/wiki/ISO_3166-2
// on February 28, 2022
// I has to combine the subdivisions by hand.
package translate

import (
	"bufio"
	"strings"
)

// Entry: ISO 3166-1 alpha-2 code, click to view the ISO 3166-2 codes of the
//	country
// Country name: English short name officially used by the ISO 3166
//	Maintenance Agency (ISO 3166/MA)
// Subdivisions assigned codes: Number and category of subdivisions assigned
//	codes in ISO 3166-2;[1] if there are more than one level of subdivisions,
//	the first-level subdivisions are shown in italics

const (
	RAW_ISO3166_2 = `AD	Andorra	7 parishes
AE	United Arab Emirates	7 emirates
AF	Afghanistan	34 provinces
AG	Antigua and Barbuda	6 parishes 2 dependencies
AI	Anguilla	—
AL	Albania	12 counties
AM	Armenia	1 city 10 regions
AO	Angola	18 provinces
AQ	Antarctica	—
AR	Argentina	1 city 23 provinces
AS	American Samoa	—
AT	Austria	9 states
AU	Australia	6 states 2 territories
AW	Aruba	—
AX	Åland Islands	—
AZ	Azerbaijan	1 autonomous republic 11 municipalities 66 rayons
BA	Bosnia and Herzegovina	2 entities 1 district with special status
BB	Barbados	11 parishes
BD	Bangladesh	8 divisions 64 districts
BE	Belgium	3 regions 10 provinces
BF	Burkina Faso	13 regions 45 provinces
BG	Bulgaria	28 regions
BH	Bahrain	4 governorates
BI	Burundi	18 provinces
BJ	Benin	12 departments
BL	Saint Barthélemy	—
BM	Bermuda	—
BN	Brunei Darussalam	4 districts
BO	Bolivia (Plurinational State of)	9 departments
BQ	Bonaire, Sint Eustatius and Saba	3 special municipalities
BR	Brazil	1 federal district 26 states
BS	Bahamas	31 districts 1 island
BT	Bhutan	20 districts
BV	Bouvet Island	—
BW	Botswana	10 districts 4 towns 2 cities
BY	Belarus	6 oblasts 1 city
BZ	Belize	6 districts
CA	Canada	10 provinces 3 territories
CC	Cocos (Keeling) Islands	—
CD	Congo, Democratic Republic of the	1 city 25 provinces
CF	Central African Republic	1 commune 14 prefectures 2 economic prefectures
CG	Congo	12 departments
CH	Switzerland	26 cantons
CI	Côte d'Ivoire	12 districts 2 autonomous districts
CK	Cook Islands	—
CL	Chile	16 regions
CM	Cameroon	10 regions
CN	China	4 municipalities 23 provinces 5 autonomous regions 2 special administrative regions
CO	Colombia	1 capital district 32 departments
CR	Costa Rica	7 provinces
CU	Cuba	15 provinces 1 special municipality
CV	Cabo Verde	2 geographical regions 22 municipalities
CW	Curaçao	—
CX	Christmas Island	—
CY	Cyprus	6 districts
CZ	Czechia	13 regions and 1 capital city 76 districts
DE	Germany	16 states
DJ	Djibouti	5 regions 1 city
DK	Denmark	5 regions
DM	Dominica	10 parishes
DO	Dominican Republic	10 regions 1 district 31 provinces
DZ	Algeria	48 provinces
EC	Ecuador	24 provinces
EE	Estonia	15 counties 64 rural municipalities 15 urban municipalities
EG	Egypt	27 governorates
EH	Western Sahara	—
ER	Eritrea	6 regions
ES	Spain	17 autonomous communities 2 autonomous cities in North Africa 50 provinces
ET	Ethiopia	2 administrations 10 regional states
FI	Finland	19 regions
FJ	Fiji	4 divisions 1 dependency 14 provinces
FK	Falkland Islands (Malvinas)	—
FM	Micronesia (Federated States of)	4 states
FO	Faroe Islands	—
FR	France	12 metropolitan regions 3 metropolitan collectivities with special status 5 overseas collectivities 1 overseas collectivity with special status 5 overseas departments 95 metropolitan departments 1 dependency 3 overseas departmental collectivity 1 overseas territory 2 overseas unique territorial collectivity 1 European collectivity
GA	Gabon	9 provinces
GB	United Kingdom of Great Britain and Northern Ireland	3 countries 1 province 32 council areas 27 two-tier counties 11 districts 77 unitary authorities 36 metropolitan districts 32 London boroughs 1 city corporation
GD	Grenada	6 parishes 1 dependency
GE	Georgia	2 autonomous republics 1 city 9 regions
GF	French Guiana	—
GG	Guernsey	—
GH	Ghana	16 regions
GI	Gibraltar	—
GL	Greenland	5 municipalities
GM	Gambia	1 city 5 divisions
GN	Guinea	7 administrative regions 1 governorate 33 prefectures
GP	Guadeloupe	—
GQ	Equatorial Guinea	2 regions 8 provinces
GR	Greece	13 administrative regions 1 self-governed part
GS	South Georgia and the South Sandwich Islands	—
GT	Guatemala	22 departments
GU	Guam	—
GW	Guinea-Bissau	3 provinces 1 autonomous sector 8 regions
GY	Guyana	10 regions
HK	Hong Kong	—
HM	Heard Island and McDonald Islands	—
HN	Honduras	18 departments
HR	Croatia	1 city 20 counties
HT	Haiti	10 departments
HU	Hungary	1 capital city 19 counties 23 cities of county right
ID	Indonesia	7 geographical units 32 provinces 1 capital district 1 special region
IE	Ireland	4 provinces 26 counties
IL	Israel	6 districts
IM	Isle of Man	—
IN	India	28 states 8 union territories
IO	British Indian Ocean Territory	—
IQ	Iraq	18 governorates
IR	Iran (Islamic Republic of)	31 provinces
IS	Iceland	8 regions 69 municipalities
IT	Italy	15 regions 80 provinces 5 autonomous regions 2 autonomous provinces 6 free municipal consortiums 14 metropolitan cities 4 decentralized regional entities
JE	Jersey	—
JM	Jamaica	14 parishes
JO	Jordan	12 governorates
JP	Japan	47 prefectures
KE	Kenya	47 counties
KG	Kyrgyzstan	2 cities 7 regions
KH	Cambodia	1 autonomous municipality 24 provinces
KI	Kiribati	3 groups of islands
KM	Comoros	3 islands
KN	Saint Kitts and Nevis	2 states 14 parishes
KP	Korea (Democratic People's Republic of)	1 capital city 1 metropolitan city 1 special city 9 provinces
KR	Korea, Republic of	6 metropolitan cities 1 special city 1 special self-governing city 8 provinces 1 special self-governing province
KW	Kuwait	6 governorates
KY	Cayman Islands	—
KZ	Kazakhstan	3 cities 14 regions
LA	Lao People's Democratic Republic	1 prefecture 17 provinces
LB	Lebanon	8 governorates
LC	Saint Lucia	10 districts
LI	Liechtenstein	11 communes
LK	Sri Lanka	9 provinces 25 districts
LR	Liberia	15 counties
LS	Lesotho	10 districts
LT	Lithuania	10 counties 9 municipalities 7 city municipalities 44 district municipalities
LU	Luxembourg	12 cantons
LV	Latvia	36 municipalities 7 state cities
LY	Libya	22 popularates
MA	Morocco	12 regions 62 provinces 13 prefectures
MC	Monaco	17 quarters
MD	Moldova, Republic of	1 autonomous territorial unit 3 cities 32 districts 1 territorial unit
ME	Montenegro	24 municipalities
MF	Saint Martin (French part)	—
MG	Madagascar	6 provinces
MH	Marshall Islands	2 chains of islands 24 municipalities
MK	North Macedonia	80 municipalities
ML	Mali	1 district 10 regions
MM	Myanmar	7 regions 7 states 1 union territory
MN	Mongolia	1 capital city 21 provinces
MO	Macao	—
MP	Northern Mariana Islands	—
MQ	Martinique	—
MR	Mauritania	15 regions
MS	Montserrat	—
MT	Malta	68 local councils
MU	Mauritius	3 dependencies 9 districts
MV	Maldives	19 administrative atolls 2 cities
MW	Malawi	3 regions 28 districts
MX	Mexico	31 states 1 federal district
MY	Malaysia	3 federal territories 13 states
MZ	Mozambique	1 city 10 provinces
NA	Namibia	14 regions
NC	New Caledonia	—
NE	Niger	1 urban community 7 departments
NF	Norfolk Island	—
NG	Nigeria	1 capital territory 36 states
NI	Nicaragua	15 departments 2 autonomous regions
NL	Netherlands	[note 1] 12 provinces 3 countries3 special municipalities
NO	Norway	11 counties 2 arctic regions
NP	Nepal	5 development regions 7 provinces 14 zones
NR	Nauru	14 districts
NU	Niue	—
NZ	New Zealand	16 regions 1 special island authority
OM	Oman	11 governorates
PA	Panama	10 provinces 4 indigenous regions
PE	Peru	25 regions 1 municipality
PF	French Polynesia	—
PG	Papua New Guinea	1 district 20 provinces 1 autonomous region
PH	Philippines	17 regions 81 provinces
PK	Pakistan	4 provinces 2 autonomous territories 1 federal territory
PL	Poland	16 voivodships
PM	Saint Pierre and Miquelon	—
PN	Pitcairn	—
PR	Puerto Rico	—
PS	Palestine, State of	16 governorates
PT	Portugal	18 districts 2 autonomous regions
PW	Palau	16 states
PY	Paraguay	1 capital 17 departments
QA	Qatar	8 municipalities
RE	Réunion	—
RO	Romania	41 departments 1 municipality
RS	Serbia	2 autonomous provinces 1 city 29 districts
RU	Russian Federation	21 republics 9 administrative territories 46 administrative regions 2 autonomous cities 1 autonomous region 4 autonomous districts
RW	Rwanda	1 town council 4 provinces
SA	Saudi Arabia	13 regions
SB	Solomon Islands	1 capital territory 9 provinces
SC	Seychelles	27 districts
SD	Sudan	18 states
SE	Sweden	21 counties
SG	Singapore	5 districts
SH	Saint Helena, Ascension and Tristan da Cunha	3 geographical entities
SI	Slovenia	212 municipalities
SJ	Svalbard and Jan Mayen	—
SK	Slovakia	8 regions
SL	Sierra Leone	1 area 4 provinces
SM	San Marino	9 municipalities
SN	Senegal	14 regions
SO	Somalia	18 regions
SR	Suriname	10 districts
SS	South Sudan	10 states
ST	Sao Tome and Principe	1 autonomous region 6 districts
SV	El Salvador	14 departments
SX	Sint Maarten (Dutch part)	—
SY	Syrian Arab Republic	14 provinces
SZ	Eswatini	4 regions
TC	Turks and Caicos Islands	—
TD	Chad	23 provinces
TF	French Southern Territories	—
TG	Togo	5 regions
TH	Thailand	1 metropolitan administration 1 special administrative city 76 provinces
TJ	Tajikistan	1 autonomous region 2 regions 1 capital territory 1 district under republic administration
TK	Tokelau	—
TL	Timor-Leste	12 municipalities 1 special administrative region
TM	Turkmenistan	5 regions 1 city
TN	Tunisia	24 governorates
TO	Tonga	5 divisions
TR	Turkey	81 provinces
TT	Trinidad and Tobago	9 regions 3 boroughs 2 cities 1 ward
TV	Tuvalu	1 town council 7 island councils
TW	Taiwan, Province of China[note 2]	13 counties 3 cities 6 special municipalities
TZ	Tanzania, United Republic of	31 regions
UA	Ukraine	24 regions 1 republic 2 cities
UG	Uganda	4 geographical regions 134 districts 1 city
UM	United States Minor Outlying Islands	9 islands, groups of islands
US	United States of America	50 states 1 district 6 outlying areas
UY	Uruguay	19 departments
UZ	Uzbekistan	1 city 12 regions 1 republic
VA	Holy See	—
VC	Saint Vincent and the Grenadines	6 parishes
VE	Venezuela (Bolivarian Republic of)	1 federal dependency 1 federal district 23 states
VG	Virgin Islands (British)	—
VI	Virgin Islands (U.S.)	—
VN	Viet Nam	58 provinces 5 municipalities
VU	Vanuatu	6 provinces
WF	Wallis and Futuna	3 administrative precincts
WS	Samoa	11 districts
YE	Yemen	1 municipality 21 governorates
YT	Mayotte	—
ZA	South Africa	9 provinces
ZM	Zambia	10 provinces
ZW	Zimbabwe	10 provinces`
)

type Iso3166_2 struct {
	Code, Country, Subdivisions string
}

var iso3166_2s []Iso3166_2

func Iso3166_2s_all() []Iso3166_2 {
	return iso3166_2s_processed()
}

// Iso3166ForCountry returns the ISO-3166 code for a country.  Will return ""
// if the country is not found.
func Iso3166ForCountry(country string) string {
	upper := strings.ToUpper(country)
	for _, iso3166 := range iso3166_2s_processed() {
		if upper == strings.ToUpper(iso3166.Country) {
			return iso3166.Code
		}
	}
	return ""
}

// CountryForIso3166 returns the country name for the given ISO-3166 code.
// Will return "" if the code is not present.
func CountryForIso3166(iso3166code string) string {
	upper := strings.ToUpper(iso3166code)
	for _, iso3166 := range iso3166_2s_processed() {
		if upper == iso3166.Code {
			return iso3166.Country
		}
	}
	return ""
}

// Iso3166Countries returns the list of ISO-3166 countries.
func Iso3166Countries() []string {
	processed := iso3166_2s_processed()
	countries := make([]string, len(processed), len(processed))
	for i, iso3166 := range processed {
		countries[i] = iso3166.Country
	}
	return countries
}

func iso3166_2s_processed() []Iso3166_2 {
	if len(iso3166_2s) > 0 {
		return iso3166_2s
	}
	// We have not processed these yet.
	scanner := bufio.NewScanner(strings.NewReader(RAW_ISO3166_2))
	for scanner.Scan() {
		line := scanner.Text()
		toks := strings.Split(line, "	")
		iso3166 := Iso3166_2{
			Code:         toks[0],
			Country:      toks[1],
			Subdivisions: toks[2]}
		iso3166_2s = append(iso3166_2s, iso3166)
	}
	return iso3166_2s
}

// Subdivisions included in ISO 3166-1
// CN	China	CN-TW Taiwan (TW) [note 2]
// CN-HK Hong Kong (HK)
// CN-MO Macao (MO)
// FI	Finland	FI-01 Åland (AX)
// FR	France	FR-BL Saint Barthélemy (BL)
// FR-GF French Guiana (GF)
// FR-GP Guadeloupe (GP)
// FR-MF Saint Martin (MF)
// FR-MQ Martinique (MQ)
// FR-NC New Caledonia (NC)
// FR-PF French Polynesia (PF)
// FR-PM Saint Pierre and Miquelon (PM)
// FR-RE Réunion (RE)
// FR-TF French Southern Territories (TF)
// FR-WF Wallis and Futuna (WF)
// FR-YT Mayotte (YT)
// NL	Netherlands[note 1]	NL-AW Aruba (AW)
// NL-BQ1 Bonaire (BQ) [note 3]
// NL-BQ2 Saba (BQ) [note 3]
// NL-BQ3 Sint Eustatius (BQ) [note 3]
// NL-CW Curaçao (CW)
// NL-SX Sint Maarten (SX)
// NO	Norway	NO-21 Svalbard (SJ) [note 4]
// NO-22 Jan Mayen (SJ) [note 4]
// US	United States	US-AS American Samoa (AS)
// US-GU Guam (GU)
// US-MP Northern Mariana Islands (MP)
// US-PR Puerto Rico (PR)
// US-UM United States Minor Outlying Islands (UM)
// US-VI Virgin Islands, U.S. (VI)
