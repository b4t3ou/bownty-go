package bownty

import (
	"reflect"
	"testing"
)

// TestCreate tests the APIReader reader failed creation
func TestCreate(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	Domain = ""
	_ = Create(2, 1)
}
// TestCreate2 tests success APIReader creation
func TestCreate2(t *testing.T) {
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

// TestGet tests the main api response getter
func TestGet(t *testing.T) {
	Domain = "https://bownty.co.uk"
	reader := Create(20, 1)
	_, err := reader.get("countries")

	if err != nil {
		t.Error("Expected content bownty country list")
	}
}

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
	cities, _ := reader.GetCityList(252)

	if reflect.TypeOf(cities).String() != "*bownty.Locations" {
		t.Error("Expected return value city list")
	}

	if !cities.Success {
		t.Error("Expected success city list response")
	}

	if len(cities.Data) != 5 {
		t.Error("Expected city count 5")
	}
}

func TestAPIReader_GetDealSitesList(t *testing.T) {
	Domain = "https://bownty.co.uk"
	reader := Create(5, 1)
	dealSites, _ := reader.GetDealSitesList(252)

	if reflect.TypeOf(dealSites).String() != "*bownty.Merchants" {
		t.Error("Expected return value deal sites list")
	}

	if !dealSites.Success {
		t.Error("Expected success deal sites list response")
	}

	if len(dealSites.Data) != 5 {
		t.Error("Expected deal sites count 5")
	}
}

func TestAPIReader_GetCategoryList(t *testing.T) {
	Domain = "https://bownty.co.uk"
	reader := Create(5, 1)
	categories, _ := reader.GetCategoryList(252)

	if reflect.TypeOf(categories).String() != "*bownty.Categories" {
		t.Error("Expected return value category list")
	}

	if !categories.Success {
		t.Error("Expected success category list response")
	}

	if len(categories.Data) != 5 {
		t.Error("Expected category count 5")
	}
}

func TestAPIReader_GetDealList(t *testing.T) {
	Domain = "https://bownty.co.uk"
	reader := Create(100, 2)
	deals, _ := reader.GetDealList(473)

	if reflect.TypeOf(deals).String() != "*bownty.Deals" {
		t.Error("Expected return value deal list")
	}

	if !deals.Success {
		t.Error("Expected success deal list response")
	}

	if len(deals.Data) != 100 {
		t.Error("Expected deal count 100, get", len(deals.Data))
	}


	cityDealCount := deals.Pagination.Count

	reader.AddExtraParams("inclusive_location=1")
	deals, _ = reader.GetDealList(252)

	if cityDealCount > deals.Pagination.Count {
		t.Error("Expected country deal count has to be bigger, city deals", cityDealCount, "country deals", deals.Pagination.Count)
	}
}
