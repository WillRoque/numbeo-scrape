package numbeo

import (
	"net/http"
	"net/url"
	"testing"
)

func TestGetCityInfo(t *testing.T) {
	cases := []struct {
		url, city, country string
	}{
		{"https://www.numbeo.com/cost-of-living/in/Basel?displayCurrency=USD", "Basel", "Switzerland"},
		{"https://www.numbeo.com/cost-of-living/in/Munich?displayCurrency=USD", "Munich", "Germany"},
		{"https://www.numbeo.com/cost-of-living/in/Amsterdam?displayCurrency=USD", "Amsterdam", "Netherlands"},
	}
	for _, c := range cases {
		u, err := url.Parse(c.url)
		if err != nil {
			t.Errorf(`url.Parse("%s") error: %v`, c.url, err)
		}
		ci, err := GetCityInfo(*u)
		if err != nil {
			t.Errorf("GetCityInfo error: %v", err)
		}
		if ci.City == "" {
			t.Errorf("Could not find city name for %s", u.String())
		}
		if ci.City != c.city {
			t.Errorf("Did not find correct city name for %s. Got %s, Want %s", u.String(), ci.City, c.city)
		}
		if ci.Country == "" {
			t.Errorf("Could not find country name for %s", u.String())
		}
		if ci.Country != c.country {
			t.Errorf("Did not find correct country name for %s. Got %s, Want %s", u.String(), ci.Country, c.country)
		}
		if ci.NetSalary == 0 {
			t.Errorf("Could not find net salary for %s", u.String())
		}
		if ci.MonthlyCost == 0 {
			t.Errorf("Could not find monthly cost for %s", u.String())
		}
		if ci.SmallAptCentre == 0 {
			t.Errorf("Could not find cost of small apartment in city centre for %s", u.String())
		}
		if ci.SmallAptOutside == 0 {
			t.Errorf("Could not find cost of small apartment outside of city centre for %s", u.String())
		}
		if ci.LargeAptCentre == 0 {
			t.Errorf("Could not find cost of large apartment in city centre for %s", u.String())
		}
		if ci.LargeAptOutside == 0 {
			t.Errorf("Could not find cost of large apartment outside of city centre for %s", u.String())
		}
	}
}

func TestGetCitiesURLs(t *testing.T) {
	citiesURLs := GetCitiesURLs()
	if len(citiesURLs) == 0 {
		t.Errorf("No cities found")
	}
}

func TestNumbeoIsUp(t *testing.T) {
	res, err := http.Get("https://www.numbeo.com")
	if err != nil {
		t.Errorf("Error on reading from numbeo.com: %v", err)
	}
	res.Body.Close()
}
