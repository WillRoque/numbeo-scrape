package numbeo

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type CityInfo struct {
	City    string
	Country string

	NetSalary   float64
	MonthlyCost float64 // Monthly cost for a single person without rent

	SmallAptCentre  float64
	SmallAptOutside float64
	LargeAptCentre  float64
	LargeAptOutside float64

	MoneyAfterMonthlyCost     float64 // NetSalary - MonthlyCost
	MoneyAfterSmallAptCentre  float64 // NetSalary - SmallAptCentre
	MoneyAfterSmallAptOutside float64 // NetSalary - SmallAptOutside
	MoneyAfterLargeAptCentre  float64 // NetSalary - LargeAptCentre
	MoneyAfterLargeAptOutside float64 // NetSalary - LargeAptOutside

	MoneyAfterSmallAptCentreAndMonthlyCost  float64 // NetSalary - SmallAptCentre - MoneyAfterMonthlyCost
	MoneyAfterSmallAptOutsideAndMonthlyCost float64 // NetSalary - SmallAptOutside - MoneyAfterMonthlyCost
	MoneyAfterLargeAptCentreAndMonthlyCost  float64 // NetSalary - LargeAptCentre - MoneyAfterMonthlyCost
	MoneyAfterLargeAptOutsideAndMonthlyCost float64 // NetSalary - LargeAptOutside - MoneyAfterMonthlyCost
}

// GetCityInfo scrapes data of a city from numbeo.com
// and returns CityInfo with its fields calculated and filled.
func GetCityInfo(u url.URL) (CityInfo, error) {
	var ci CityInfo

	res, err := http.Get(u.String())
	if err != nil {
		return ci, err
	}
	defer res.Body.Close()

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return ci, err
	}

	// Find city and country
	doc.Find(".breadcrumb_link").Each(func(i int, el *goquery.Selection) {
		nextContent, exists := el.Next().Attr("content")
		if !exists {
			return
		}
		if nextContent == "2" {
			ci.Country = el.Children().First().Text()
		}
		if nextContent == "3" {
			ci.City = el.Children().First().Text()
		}
	})

	// Find the monthly cost for a single person without rent
	doc.Find(".emp_number").Each(func(i int, el *goquery.Selection) {
		if !strings.Contains(el.Parent().Text(), "A single person estimated monthly costs") {
			return
		}
		el.Find(".in_other_currency").Remove()
		valStr := getFloatString(el.Text())
		if ci.MonthlyCost, err = strconv.ParseFloat(valStr, 64); err != nil {
			log.Printf("error parsing monthly cost for %s: %v", u.String(), err)
		}
	})

	// Find rent prices and net salary
	doc.Find(".priceValue ").Each(func(i int, el *goquery.Selection) {
		var varPtr *float64
		switch {
		case strings.Contains(el.Parent().Text(), "Apartment (1 bedroom) in City Centre"):
			varPtr = &ci.SmallAptCentre
		case strings.Contains(el.Parent().Text(), "Apartment (1 bedroom) Outside of Centre"):
			varPtr = &ci.SmallAptOutside
		case strings.Contains(el.Parent().Text(), "Apartment (3 bedrooms) in City Centre"):
			varPtr = &ci.LargeAptCentre
		case strings.Contains(el.Parent().Text(), "Apartment (3 bedrooms) Outside of Centre"):
			varPtr = &ci.LargeAptOutside
		case strings.Contains(el.Parent().Text(), "Average Monthly Net Salary (After Tax)"):
			varPtr = &ci.NetSalary
		default:
			return
		}
		valStr := getFloatString(el.Text())
		if val, err := strconv.ParseFloat(valStr, 64); err == nil {
			*varPtr = val
		} else {
			log.Printf("error parsing price for %s (original: %s): %v", u.String(), el.Text(), err)
		}
	})

	ci.MoneyAfterMonthlyCost = ci.NetSalary - ci.MonthlyCost
	ci.MoneyAfterSmallAptCentre = ci.NetSalary - ci.SmallAptCentre
	ci.MoneyAfterSmallAptOutside = ci.NetSalary - ci.SmallAptOutside
	ci.MoneyAfterLargeAptCentre = ci.NetSalary - ci.LargeAptCentre
	ci.MoneyAfterLargeAptOutside = ci.NetSalary - ci.LargeAptOutside

	ci.MoneyAfterSmallAptCentreAndMonthlyCost = ci.MoneyAfterSmallAptCentre - ci.MonthlyCost
	ci.MoneyAfterSmallAptOutsideAndMonthlyCost = ci.MoneyAfterSmallAptOutside - ci.MonthlyCost
	ci.MoneyAfterLargeAptCentreAndMonthlyCost = ci.MoneyAfterLargeAptCentre - ci.MonthlyCost
	ci.MoneyAfterLargeAptOutsideAndMonthlyCost = ci.MoneyAfterLargeAptOutside - ci.MonthlyCost

	return ci, nil
}

// GetCitiesURLs scrapes URLs from numbeo.com of cities with relevant amount of data.
func GetCitiesURLs() []url.URL {
	numbeoURL := "https://www.numbeo.com/cost-of-living/rankings_current.jsp"
	res, err := http.Get(numbeoURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the URL of each city
	var urls []url.URL
	doc.Find(".cityOrCountryInIndicesTable a").Each(func(i int, el *goquery.Selection) {
		href, ok := el.Attr("href")
		if !ok {
			return
		}
		href += "?displayCurrency=USD"
		u, err := url.Parse(href)
		if err != nil {
			log.Printf(`url.Parse("%s") error: %v`, href, err)
		}
		urls = append(urls, *u)
	})

	return urls
}
