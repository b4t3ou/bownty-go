# bownty-go
Bownty Golang API Library

[![Build Status](https://travis-ci.org/b4t3ou/bownty-go.svg?branch=master)](https://travis-ci.org/b4t3ou/bownty-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/b4t3ou/bownty-go)](https://goreportcard.com/report/github.com/b4t3ou/bownty-go)

Small Golang library for the Bownty API <https://wiki.bownty.net/display/API/Deals+API>

####Install

go get github.com/b4t3ou/bownty-go

####Usage

```go
Domain = "https://bownty.co.uk" //or your subdomain

First parameter limit second parameter page
reader := Create(20, 1)

// Get country list
countries, _ := reader.GetCountryList()

// Get city list, parameter country code
countries, _ := reader.GetCityList(252)

// Get deal site list, parameter country code
countries, _ := GetDealSitesList(252)

// Get category list, parameter country code
countries, _ := reader.GetCategoryList(252)

// Get deal list, parameter country or city code
countries, _ := reader.GetDealList(473)
```

####Sample

```Go
package main

import (
	"github.com/b4t3ou/bownty-go"
	"fmt"
)

func main() {
	bownty.Domain = "https://bownty.co.uk"
	reader := bownty.Create(20, 1)
	reader.AddExtraParams("inclusive_location=1")
	deals, _ := reader.GetDealList(252)

	// Deal count in the UK
	fmt.Println(deals.Pagination.Count)
}
````
