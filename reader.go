package bownty

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// APIReader main struct.
type APIReader struct {
	Urls        map[string]string
	Domain      string
	Request     *http.Request
	Error       error
	Response    *http.Response
	Content     []byte
	Limit       int
	Page        int
	extraParams []string
	CalledURL   string
}

var endPoints = map[string]string{
	"countries":  "/api/countries.json",
	"cities":     "/locations/:country_id/cities.json",
	"deal_sites":  "/locations/:country_id/deal_sites.json",
	"categories": "/locations/:country_id/categories.json",
	"deals":      "/locations/:city_id/deals.json",
}

// Create creates a new instance of the APIReader
func Create(domain string, limit, page int) APIReader {

	if domain == "" {
		panic("Bownty domain has to be set!")
	}

	return APIReader{
		Urls:        endPoints,
		Limit:       limit,
		Page:        page,
		Domain:      domain,
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

// GetMerchantList returns with *Merchants
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
