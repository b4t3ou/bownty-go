package bownty

import (
	"reflect"
	"testing"
)

// TestCreate tests the API reader creation
func TestCreate(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	Domain = "https://bownty.co.uk"
	reader := Create(20, 1)

	if reader.Limit != 20 {
		t.Error("Expected limit 20")
	}

	if reader.Page != 1 {
		t.Error("Expected limit 1")
	}

	if reflect.TypeOf(reader).String() != "bownty.APIReader" {
		t.Error("Faild to create api reader")
	}

	Domain = ""
	_ = Create(2, 1)
}

// TestAPIReader_AddExtraParams tests adding extra params to your query
func TestAPIReader_AddExtraParams(t *testing.T) {
	Domain = "https://bownty.co.uk"
	reader := Create(20, 1)
	reader.AddExtraParams("price_range=100-200")

	if len(reader.extraParams) != 1 {
		t.Error("Expected extra param count 1")
	}

	reader.AddExtraParams("save_range=10-40", "locale=en")

	if len(reader.extraParams) != 3 {
		t.Error("Expected extra param count 3")
	}

	if reader.extraParams[0] != "price_range=100-200" {
		t.Error("Expected first parameter price_range=100-200")
	}

	if reader.extraParams[1] != "save_range=10-40" {
		t.Error("Expected second parameter save_range=10-40")
	}

	if reader.extraParams[2] != "locale=en" {
		t.Error("Expected third parameter locale=en")
	}
}

// TestGet tests the main api caller
func TestGet(t *testing.T) {
	Domain = "https://bownty.co.uk"
	reader := Create(20, 1)
	_, err := reader.get("countries")

	if err != nil {
		t.Error("Expected content bownty country list")
	}
}

// TestAPIReader_GetCountryList tests to get back the country list
func TestAPIReader_GetCountryList(t *testing.T) {
	Domain = "https://bownty.co.uk"
	reader := Create(5, 1)
	countries, _ := reader.GetCountryList()

	if reflect.TypeOf(countries).String() != "*bownty.Locations" {
		t.Error("Expected return value countries list")
	}

	if !countries.Success {
		t.Error("Expected success country list response")
	}

	if len(countries.Data) != 5 {
		t.Error("Expected country count 5")
	}
}

func TestAPIReader_GetCityList(t *testing.T) {
	Domain = "https://bownty.co.uk"
	reader := Create(5, 1)
	countries, _ := reader.GetCityList(252)

	if reflect.TypeOf(countries).String() != "*bownty.Locations" {
		t.Error("Expected return value city list")
	}

	if !countries.Success {
		t.Error("Expected success city list response")
	}

	if len(countries.Data) != 5 {
		t.Error("Expected city count 5")
	}
}

func TestAPIReader_GetDealSitesList(t *testing.T) {
	Domain = "https://bownty.co.uk"
	reader := Create(5, 1)
	countries, _ := reader.GetDealSitesList(252)

	if reflect.TypeOf(countries).String() != "*bownty.Merchants" {
		t.Error("Expected return value deal sites list")
	}

	if !countries.Success {
		t.Error("Expected success deal sites list response")
	}

	if len(countries.Data) != 5 {
		t.Error("Expected deal sites count 5")
	}
}
