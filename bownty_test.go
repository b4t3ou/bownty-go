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

	reader := Create("https://bownty.co.uk", 20, 1)

	if reader.Limit != 20 {
		t.Error("Expected limit 20")
	}

	if reader.Page != 1 {
		t.Error("Expected limit 1")
	}

	if reflect.TypeOf(reader).String() != "bownty.APIReader" {
		t.Error("Faild to create api reader")
	}

	_ = Create("", 2, 1)
}

// TestAPIReader_AddExtraParams tests adding extra params to your query
func TestAPIReader_AddExtraParams(t *testing.T) {
	reader := Create("https://bownty.co.uk", 20, 1)
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
	reader := Create("https://bownty.co.uk", 20, 1)
	_, err := reader.get("countries")

	if err != nil {
		t.Error("Expected content bownty country list")
	}
}

// TestAPIReader_GetCountryList tests to get back the country list
func TestAPIReader_GetCountryList(t *testing.T) {
	reader := Create("https://bownty.co.uk", 5, 1)
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
