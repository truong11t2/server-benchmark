package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

// PageData struct to hold data for template rendering
type PageData struct {
	Results []Country
}

type ListCountries struct {
	Countries []Country
}

type Country struct {
	Id   string
	Name string
}

type Student struct {
	Name  string
	Grade int
}

func main() {
	list := []Country{{Id: "AF", Name: "Afghanistan"}, {Id: "AX", Name: "\u00c5land Islands"}, {Id: "AL", Name: "Albania"}, {Id: "DZ", Name: "Algeria"}, {Id: "AS", Name: "American Samoa"}, {Id: "AD", Name: "Andorra"}, {Id: "AO", Name: "Angola"}, {Id: "AI", Name: "Anguilla"}, {Id: "AQ", Name: "Antarctica"}, {Id: "AG", Name: "Antigua & Barbuda"}, {Id: "AR", Name: "Argentina"}, {Id: "AM", Name: "Armenia"}, {Id: "AW", Name: "Aruba"}, {Id: "AU", Name: "Australia"}, {Id: "AT", Name: "Austria"}, {Id: "AZ", Name: "Azerbaijan"}, {Id: "BS", Name: "Bahamas"}, {Id: "BH", Name: "Bahrain"}, {Id: "BD", Name: "Bangladesh"}, {Id: "BB", Name: "Barbados"}, {Id: "BY", Name: "Belarus"}, {Id: "BE", Name: "Belgium"}, {Id: "BZ", Name: "Belize"}, {Id: "BJ", Name: "Benin"}, {Id: "BM", Name: "Bermuda"}, {Id: "BT", Name: "Bhutan"}, {Id: "BO", Name: "Bolivia"}, {Id: "BA", Name: "Bosnia & Herzegovina"}, {Id: "BW", Name: "Botswana"}, {Id: "BV", Name: "Bouvet Island"}, {Id: "BR", Name: "Brazil"}, {Id: "IO", Name: "British Indian Ocean Territory"}, {Id: "VG", Name: "British Virgin Islands"}, {Id: "BN", Name: "Brunei"}, {Id: "BG", Name: "Bulgaria"}, {Id: "BF", Name: "Burkina Faso"}, {Id: "BI", Name: "Burundi"}, {Id: "KH", Name: "Cambodia"}, {Id: "CM", Name: "Cameroon"}, {Id: "CA", Name: "Canada"}, {Id: "CV", Name: "Cape Verde"}, {Id: "BQ", Name: "Caribbean Netherlands"}, {Id: "KY", Name: "Cayman Islands"}, {Id: "CF", Name: "Central African Republic"}, {Id: "TD", Name: "Chad"}, {Id: "CL", Name: "Chile"}, {Id: "CN", Name: "China"}, {Id: "CX", Name: "Christmas Island"}, {Id: "CC", Name: "Cocos (Keeling) Islands"}, {Id: "CO", Name: "Colombia"}, {Id: "KM", Name: "Comoros"}, {Id: "CG", Name: "Congo - Brazzaville"}, {Id: "CD", Name: "Congo - Kinshasa"}, {Id: "CK", Name: "Cook Islands"}, {Id: "CR", Name: "Costa Rica"}, {Id: "CI", Name: "C\u00f4te d\u2019Ivoire"}, {Id: "HR", Name: "Croatia"}, {Id: "CU", Name: "Cuba"}, {Id: "CW", Name: "Cura\u00e7ao"}, {Id: "CY", Name: "Cyprus"}, {Id: "CZ", Name: "Czechia"}, {Id: "DK", Name: "Denmark"}, {Id: "DJ", Name: "Djibouti"}, {Id: "DM", Name: "Dominica"}, {Id: "DO", Name: "Dominican Republic"}, {Id: "EC", Name: "Ecuador"}, {Id: "EG", Name: "Egypt"}, {Id: "SV", Name: "El Salvador"}, {Id: "GQ", Name: "Equatorial Guinea"}, {Id: "ER", Name: "Eritrea"}, {Id: "EE", Name: "Estonia"}, {Id: "SZ", Name: "Eswatini"}, {Id: "ET", Name: "Ethiopia"}, {Id: "FK", Name: "Falkland Islands"}, {Id: "FO", Name: "Faroe Islands"}, {Id: "FJ", Name: "Fiji"}, {Id: "FI", Name: "Finland"}, {Id: "FR", Name: "France"}, {Id: "GF", Name: "French Guiana"}, {Id: "PF", Name: "French Polynesia"}, {Id: "TF", Name: "French Southern Territories"}, {Id: "GA", Name: "Gabon"}, {Id: "GM", Name: "Gambia"}, {Id: "GE", Name: "Georgia"}, {Id: "DE", Name: "Germany"}, {Id: "GH", Name: "Ghana"}, {Id: "GI", Name: "Gibraltar"}, {Id: "GR", Name: "Greece"}, {Id: "GL", Name: "Greenland"}, {Id: "GD", Name: "Grenada"}, {Id: "GP", Name: "Guadeloupe"}, {Id: "GU", Name: "Guam"}, {Id: "GT", Name: "Guatemala"}, {Id: "GG", Name: "Guernsey"}, {Id: "GN", Name: "Guinea"}, {Id: "GW", Name: "Guinea-Bissau"}, {Id: "GY", Name: "Guyana"}, {Id: "HT", Name: "Haiti"}, {Id: "HM", Name: "Heard & McDonald Islands"}, {Id: "HN", Name: "Honduras"}, {Id: "HK", Name: "Hong Kong SAR China"}, {Id: "HU", Name: "Hungary"}, {Id: "IS", Name: "Iceland"}, {Id: "IN", Name: "India"}, {Id: "ID", Name: "Indonesia"}, {Id: "IR", Name: "Iran"}, {Id: "IQ", Name: "Iraq"}, {Id: "IE", Name: "Ireland"}, {Id: "IM", Name: "Isle of Man"}, {Id: "IL", Name: "Israel"}, {Id: "IT", Name: "Italy"}, {Id: "JM", Name: "Jamaica"}, {Id: "JP", Name: "Japan"}, {Id: "JE", Name: "Jersey"}, {Id: "JO", Name: "Jordan"}, {Id: "KZ", Name: "Kazakhstan"}, {Id: "KE", Name: "Kenya"}, {Id: "KI", Name: "Kiribati"}, {Id: "KW", Name: "Kuwait"}, {Id: "KG", Name: "Kyrgyzstan"}, {Id: "LA", Name: "Laos"}, {Id: "LV", Name: "Latvia"}, {Id: "LB", Name: "Lebanon"}, {Id: "LS", Name: "Lesotho"}, {Id: "LR", Name: "Liberia"}, {Id: "LY", Name: "Libya"}, {Id: "LI", Name: "Liechtenstein"}, {Id: "LT", Name: "Lithuania"}, {Id: "LU", Name: "Luxembourg"}, {Id: "MO", Name: "Macao SAR China"}, {Id: "MG", Name: "Madagascar"}, {Id: "MW", Name: "Malawi"}, {Id: "MY", Name: "Malaysia"}, {Id: "MV", Name: "Maldives"}, {Id: "ML", Name: "Mali"}, {Id: "MT", Name: "Malta"}, {Id: "MH", Name: "Marshall Islands"}, {Id: "MQ", Name: "Martinique"}, {Id: "MR", Name: "Mauritania"}, {Id: "MU", Name: "Mauritius"}, {Id: "YT", Name: "Mayotte"}, {Id: "MX", Name: "Mexico"}, {Id: "FM", Name: "Micronesia"}, {Id: "MD", Name: "Moldova"}, {Id: "MC", Name: "Monaco"}, {Id: "MN", Name: "Mongolia"}, {Id: "ME", Name: "Montenegro"}, {Id: "MS", Name: "Montserrat"}, {Id: "MA", Name: "Morocco"}, {Id: "MZ", Name: "Mozambique"}, {Id: "MM", Name: "Myanmar (Burma)"}, {Id: "NA", Name: "Namibia"}, {Id: "NR", Name: "Nauru"}, {Id: "NP", Name: "Nepal"}, {Id: "NL", Name: "Netherlands"}, {Id: "NC", Name: "New Caledonia"}, {Id: "NZ", Name: "New Zealand"}, {Id: "NI", Name: "Nicaragua"}, {Id: "NE", Name: "Niger"}, {Id: "NG", Name: "Nigeria"}, {Id: "NU", Name: "Niue"}, {Id: "NF", Name: "Norfolk Island"}, {Id: "KP", Name: "North Korea"}, {Id: "MK", Name: "North Macedonia"}, {Id: "MP", Name: "Northern Mariana Islands"}, {Id: "NO", Name: "Norway"}, {Id: "OM", Name: "Oman"}, {Id: "PK", Name: "Pakistan"}, {Id: "PW", Name: "Palau"}, {Id: "PS", Name: "Palestinian Territories"}, {Id: "PA", Name: "Panama"}, {Id: "PG", Name: "Papua New Guinea"}, {Id: "PY", Name: "Paraguay"}, {Id: "PE", Name: "Peru"}, {Id: "PH", Name: "Philippines"}, {Id: "PN", Name: "Pitcairn Islands"}, {Id: "PL", Name: "Poland"}, {Id: "PT", Name: "Portugal"}, {Id: "PR", Name: "Puerto Rico"}, {Id: "QA", Name: "Qatar"}, {Id: "RE", Name: "R\u00e9union"}, {Id: "RO", Name: "Romania"}, {Id: "RU", Name: "Russia"}, {Id: "RW", Name: "Rwanda"}, {Id: "WS", Name: "Samoa"}, {Id: "SM", Name: "San Marino"}, {Id: "ST", Name: "S\u00e3o Tom\u00e9 & Pr\u00edncipe"}, {Id: "SA", Name: "Saudi Arabia"}, {Id: "SN", Name: "Senegal"}, {Id: "RS", Name: "Serbia"}, {Id: "SC", Name: "Seychelles"}, {Id: "SL", Name: "Sierra Leone"}, {Id: "SG", Name: "Singapore"}, {Id: "SX", Name: "Sint Maarten"}, {Id: "SK", Name: "Slovakia"}, {Id: "SI", Name: "Slovenia"}, {Id: "SB", Name: "Solomon Islands"}, {Id: "SO", Name: "Somalia"}, {Id: "ZA", Name: "South Africa"}, {Id: "GS", Name: "South Georgia & South Sandwich Islands"}, {Id: "KR", Name: "South Korea"}, {Id: "SS", Name: "South Sudan"}, {Id: "ES", Name: "Spain"}, {Id: "LK", Name: "Sri Lanka"}, {Id: "BL", Name: "St. Barth\u00e9lemy"}, {Id: "SH", Name: "St. Helena"}, {Id: "KN", Name: "St. Kitts & Nevis"}, {Id: "LC", Name: "St. Lucia"}, {Id: "MF", Name: "St. Martin"}, {Id: "PM", Name: "St. Pierre & Miquelon"}, {Id: "VC", Name: "St. Vincent & Grenadines"}, {Id: "SD", Name: "Sudan"}, {Id: "SR", Name: "Suriname"}, {Id: "SJ", Name: "Svalbard & Jan Mayen"}, {Id: "SE", Name: "Sweden"}, {Id: "CH", Name: "Switzerland"}, {Id: "SY", Name: "Syria"}, {Id: "TW", Name: "Taiwan"}, {Id: "TJ", Name: "Tajikistan"}, {Id: "TZ", Name: "Tanzania"}, {Id: "TH", Name: "Thailand"}, {Id: "TL", Name: "Timor-Leste"}, {Id: "TG", Name: "Togo"}, {Id: "TK", Name: "Tokelau"}, {Id: "TO", Name: "Tonga"}, {Id: "TT", Name: "Trinidad & Tobago"}, {Id: "TN", Name: "Tunisia"}, {Id: "TR", Name: "Turkey"}, {Id: "TM", Name: "Turkmenistan"}, {Id: "TC", Name: "Turks & Caicos Islands"}, {Id: "TV", Name: "Tuvalu"}, {Id: "UM", Name: "U.S. Outlying Islands"}, {Id: "VI", Name: "U.S. Virgin Islands"}, {Id: "UG", Name: "Uganda"}, {Id: "UA", Name: "Ukraine"}, {Id: "AE", Name: "United Arab Emirates"}, {Id: "GB", Name: "United Kingdom"}, {Id: "US", Name: "United States"}, {Id: "UY", Name: "Uruguay"}, {Id: "UZ", Name: "Uzbekistan"}, {Id: "VU", Name: "Vanuatu"}, {Id: "VA", Name: "Vatican City"}, {Id: "VE", Name: "Venezuela"}, {Id: "VN", Name: "Vietnam"}, {Id: "WF", Name: "Wallis & Futuna"}, {Id: "EH", Name: "Western Sahara"}, {Id: "YE", Name: "Yemen"}, {Id: "ZM", Name: "Zambia"}, {Id: "ZW", Name: "Zimbabwe"}}

	// Parse the template file
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		log.Fatal(err)
	}

	// Parse the template file
	tmplCountry, err := template.ParseFiles("db/index.html")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	// Define a handler function for static content
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Define a handler function for template rendering
	http.HandleFunc("/template", func(w http.ResponseWriter, r *http.Request) {
		// Data to be displayed in the template
		//data := PageData{Results: []Country{{Id: "VN", Name: "Vietnam"}, {Id: "US", Name: "United States"}}}
		data := PageData{Results: list}

		// Execute the template with the provided data
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Define a handler function for dynamic content
	http.HandleFunc("/db", func(w http.ResponseWriter, r *http.Request) {

		var results []Country
		// Connect to the PostgreSQL database
		conn, err := pgx.Connect(context.Background(), "postgresql://username:password@127.0.0.1:5432/benchmark")
		if err != nil {
			http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
			return
		}
		defer conn.Close(context.Background())

		// Execute a query to fetch data
		rows, err := conn.Query(context.Background(), "SELECT * FROM countries")
		if err != nil {
			http.Error(w, "Unable to execute query", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var columnValue Country
			err := rows.Scan(&columnValue.Id, &columnValue.Name)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Unable to scan rows", http.StatusInternalServerError)
				return
			}
			// Extract data from the query result
			results = append(results, columnValue)
		}

		// Data to be displayed in the template
		data := ListCountries{Countries: results}

		// Execute the template with the provided data
		err = tmplCountry.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Start the HTTP server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
