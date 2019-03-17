// Package function contains an HTTP Cloud Function.
package function

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sort"
	"time"

	"github.com/WillRoque/numbeo-scrape/numbeo"
)

// ScrapeNumbeo scrapes numbeo.com to find cities with relevant amount of data,
// and then calculates how much money one would have left at the end of the month
// after paying living expenses and rent.
func ScrapeNumbeo(w http.ResponseWriter, r *http.Request) {
	citiesURLs := numbeo.GetCitiesURLs()
	cities := make([]numbeo.CityInfo, len(citiesURLs))
	c := make(chan numbeo.CityInfo)

	log.Printf("Cities: %d", len(cities))

	for i, u := range citiesURLs {
		log.Printf("Getting city info %d/%d %s", i+1, len(citiesURLs), u.String())
		go getCityInfo(u, c)
	}

	for i := 0; i < len(cities); i++ {
		log.Printf("Receiving city info %d/%d", i+1, len(cities))
		cities[i] = <-c
	}

	sort.Sort(byMoneyAfterLargeAptOutsideAndMonthlyCost(cities))

	spreadsheet := numbeo.GenerateSpreadsheet(cities)
	var buf bytes.Buffer
	spreadsheet.Write(&buf)
	fileBytes := buf.Bytes()

	fileName := "remaining-money.xlsx"
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	http.ServeContent(w, r, fileName, time.Now(), bytes.NewReader(fileBytes))
}

// getCityInfo calculates the information of a city after its scraping
// and sends it through the channel c.
func getCityInfo(u url.URL, c chan numbeo.CityInfo) {
	ci, err := numbeo.GetCityInfo(u)
	if err != nil {
		log.Printf("numbeo.GetCityInfo(%q) error: %v", u.String(), err)
	}
	c <- ci
}

// byMoneyAfterLargeAptOutsideAndMonthlyCost is used to sort cities
// by the amount of money left after paying living expenses and
// rent of a large apartment outside the city centre in descending order.
type byMoneyAfterLargeAptOutsideAndMonthlyCost []numbeo.CityInfo

func (s byMoneyAfterLargeAptOutsideAndMonthlyCost) Len() int      { return len(s) }
func (s byMoneyAfterLargeAptOutsideAndMonthlyCost) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s byMoneyAfterLargeAptOutsideAndMonthlyCost) Less(i, j int) bool {
	return s[i].MoneyAfterLargeAptOutsideAndMonthlyCost > s[j].MoneyAfterLargeAptOutsideAndMonthlyCost
}
