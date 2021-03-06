package bownty

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var Domain string

// APIReader main struct.
type APIReader struct {
	Urls        map[string]string
	Domain      string
	Limit       int
	Page        int
	extraParams []string
	CalledURL   string
	Password    string
}

var endPoints = map[string]string{
	"countries":    "/api/countries.json",
	"cities":       "/locations/:country_id/cities.json",
	"deal_sites":   "/locations/:country_id/deal_sites.json",
	"categories":   "/locations/:country_id/categories.json",
	"deals":        "/locations/:city_id/deals.json",
	"transactions": "/transactions.json",
}

// Create creates a new instance of the APIReader
func Create(limit, page int) APIReader {

	if Domain == "" {
		panic("Bownty domain has to be set!")
	}

	return APIReader{
		Urls:        endPoints,
		Limit:       limit,
		Page:        page,
		Domain:      Domain,
		extraParams: []string{},
	}
}

// AddExtraParams appends the given parameters to the api query
func (r *APIReader) AddExtraParams(params ...string) {
	for _, param := range params {
		r.extraParams = append(r.extraParams, param)
	}
}

// GetCountryList returns with *Locations struct
// Returns with the country endpoint data
func (r *APIReader) GetCountryList() (*Locations, error) {
	result := &Locations{}
	content, err := r.get("countries")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetCountryList returns with *Locations struct
// Returns with the cities endpoint data filtered by the country
func (r *APIReader) GetCityList(countryId int) (*Locations, error) {
	result := &Locations{}
	content, err := r.get("cities", ":country_id", strconv.Itoa(countryId))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetDealSitesList returns with *Merchants
// Returns with the deal sites endpoint data filtered by the country
func (r *APIReader) GetDealSitesList(countryId int) (*Merchants, error) {
	result := &Merchants{}
	content, err := r.get("deal_sites", ":country_id", strconv.Itoa(countryId))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetCategoryList returns with *Categories
// Returns with the categories endpoint data filtered by the country
func (r *APIReader) GetCategoryList(countryId int) (*Categories, error) {
	result := &Categories{}
	content, err := r.get("categories", ":country_id", strconv.Itoa(countryId))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetDealList returns with *Deals
// Returns with the deals endpoint data filtered by the city or the country id
// If you want to get all the deals in a country add this extra param "inclusive_location=1"
func (r *APIReader) GetDealList(cityId int) (*Deals, error) {
	result := &Deals{}
	content, err := r.get("deals", ":city_id", strconv.Itoa(cityId))

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *APIReader) GetTransactionList() (*Transactions, error) {
	result := &Transactions{}
	content, err := r.get("transactions")

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *APIReader) get(key string, params ...string) ([]byte, error) {
	url := r.Urls[key]

	if len(params) > 1 {
		for index, param := range params {
			if index%2 == 0 {
				url = strings.Replace(url, param, params[index+1], 1)
			}
		}
	}

	url += "?limit=" + strconv.Itoa(r.Limit) + "&page=" + strconv.Itoa(r.Page)

	if len(r.extraParams) > 0 {
		for _, ep := range r.extraParams {
			url += "&" + ep
		}
	}

	r.CalledURL = url

	req, err := http.NewRequest("GET", r.Domain+url, bytes.NewBuffer([]byte{}))
	req.Header.Set("Content-Type", "application/json")

	if r.Password != "" {
		req.Header.Set("X-Site-Authorization", r.Password)
	}

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return content, nil
}
