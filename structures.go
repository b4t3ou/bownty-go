package bownty

type Pagination struct {
	PageCount   int  `json:"page_count"`
	CurrentPage int  `json:"current_page"`
	HasNextPage bool `json:"has_next_page"`
	HasPrevPage bool `json:"has_prev_page"`
	Count       int  `json:"count"`
	Limit       int  `json:"limit"`
}

type Locations struct {
	Success    bool       `json:"success"`
	Data       []Location `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Location struct {
	Location Place `json:"location"`
}

type Place struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	EnglishName string  `json:"name_english"`
	Path        string  `json:"path"`
	Long        float32 `json:"lng,omitempty"`
	Lat         float32 `json:"lat,omitempty"`
	Slug        string  `json:"slug,omitempty"`
}

type Merchants struct {
	Success    bool            `json:"success"`
	Data       []MerchantsMain `json:"data"`
	Pagination Pagination      `json:"pagination"`
}

type MerchantsMain struct {
	DealSite Merchant `json:"DealSite"`
}

type Merchant struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Slug      string `json:"slug"`
	ImagePath string `json:"image_path"`
}

type Categories struct {
	Success    bool           `json:"success"`
	Data       []CategoryMain `json:"data"`
	Pagination Pagination     `json:"pagination"`
}

type CategoryMain struct {
	Category Category       `json:"category"`
	Children []CategoryMain `json:"children,omitempty"`
}

type Category struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	ParentId   string `json:"parent_id,omitempty"`
	Position   string `json:"position"`
	DealsCount int    `json:"deals_count"`
	IsLocal    bool   `json:"is_local"`
}

type Deals struct {
	Success    bool       `json:"success"`
	Data       []Deal     `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Deal struct {
	Id                    int                `json:"id"`
	Identifier            int                `json:"identifier"`
	Url                   string             `json:"url"`
	Slug                  string             `json:"slug"`
	DealTypeId            int                `json:"deal_type_id"`
	DealType              string             `json:"deal_type"`
	Name                  string             `json:"name"`
	Description           string             `json:"description"`
	Currency              string             `json:"currency"`
	CurrencySymbol        string             `json:"currency_symbol"`
	Price                 float32            `json:"price"`
	OriginalPrice         float32            `json:"original_price"`
	PriceEuro             float32            `json:"price_euro"`
	Savings               float32            `json:"savings"`
	SavingsPercent        float32            `json:"savings_percent"`
	StartTime             string             `json:"start_time"`
	EndTime               string             `json:"end_time"`
	ExpiryTime            string             `json:"expiry_time"`
	IsActive              bool               `json:"is_active"`
	IsSponsored           bool               `json:"is_sponsored"`
	IsManuallyCategorized bool               `json:"is_manually_categorized"`
	IsLocal               bool               `json:"is_local"`
	IsExperience          int                `json:"is_experience"`
	IsMerchantProcessed   bool               `json:"is_merchant_processed"`
	IsRedeemableNow       bool               `json:"is_redeemable_now"`
	IsMainDeal            bool               `json:"is_main_deal"`
	IsVisibleOnMap        bool               `json:"is_visible_on_map"`
	HasUnsafeWord         bool               `json:"has_unsafe_word"`
	RedeemableDays        DealRedeemableDays `json:"redeemable_days"`
	Merchant              interface{}        `json:"merchant"` //DealMerchant
	MerchantRating        interface{}        `json:"merchant"` //DealMerchantRate
	DealSite              DealSite           `json:"deal_site"`
	Categories            []DealCategory     `json:"categories"`
	Category              DealCategory       `json:"category"`
	Locations             []DealLocation     `json:"locations"`
	Location              DealLocation       `json:"location"`
	Tags                  []DealTag          `json:"tags"`
	GeoPoint              GeoPoint           `json:"geo_point"`
	Images                []DealImage        `json:"images"`
	Themes                interface{}        `json:"themes"`
	Score                 float32            `json:"score"`
	Created               string             `json:"created"`
	Links                 Links              `json:"_links"`
	Localized             DealLocalized      `json:"localized"`
}

type DealRedeemableDays struct {
	Monday    bool `json:"monday"`
	Tuesday   bool `json:"tuesday"`
	Wednesday bool `json:"wednesday"`
	Thursday  bool `json:"thursday"`
	Friday    bool `json:"friday"`
	Saturday  bool `json:"saturday"`
	Sunday    bool `json:"sunday"`
}

type DealMerchant struct {
	Id             int                   `json:"id"`
	ParentId       int                   `json:"parent_id"`
	AllowIndexing  bool                  `json:"allow_indexing"`
	Slug           string                `json:"slug"`
	Name           string                `json:"name"`
	YelpIdentifier string                `json:"yelp_identifier"`
	FacebookId     string                `json:"facebook_id"`
	FoursquareId   string                `json:"foursquare_id"`
	GoogleId       string                `json:"google_id"`
	Address        string                `json:"address"`
	Zip            string                `json:"zip"`
	City           string                `json:"city"`
	Neighborhood   string                `json:"neighborhood"`
	Phone          string                `json:"phone"`
	Website        string                `json:"website"`
	GeoPoint       GeoPoint              `json:"geo_point"`
	Ratings        []DealMerchantRatings `json:"ratings"`
	Links          Links                 `json:"_links"`
}

type GeoPoint struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

type DealMerchantRatings struct {
	Provider     string `json:"provider"`
	Likes        string `json:"likes"`
	TalkingAbout string `json:"talking_about"`
	Url          string `json:"url"`
}

type Links struct {
	Page   LinksContent `json:"page,omitempty"`
	Self   LinksContent `json:"_self,omitempty"`
	Target LinksContent `json:"target,omitempty"`
}

type LinksContent struct {
	Href string `json:"href"`
}

type DealMerchantRate struct {
	Provider         string  `json:"provider"`
	Rating           string  `json:"rating"`
	ReviewCount      string  `json:"review_count"`
	Url              string  `json:"url"`
	NormalizedRating float32 `json:"normalized_rating"`
	DisplayRating    float32 `json:"display_rating"`
}

type DealSite struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Links Links  `json:"_links"`
}

type DealCategory struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	NameEn string `json:"name_en"`
	Slug   string `json:"slug"`
	SlugEN string `json:"slug_en"`
}

type DealLocation struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type DealTag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DealImage struct {
	OriginalSize string `json:"original_size"`
	LargeSize    string `json:"large_size"`
	SmallSize    string `json:"small_size"`
	Sponsored    string `json:"sponsored"`
	Mobile       string `json:"mobile"`
	Mobile2      string `json:"mobile2"`
	Map          string `json:"map"`
}

type DealLocalized struct {
	Price                  string `json:"price"`
	FormattedPrice         string `json:"formatted_price"`
	OriginalPrice          string `json:"original_price"`
	FormattedOriginalPrice string `json:"formatted_original_price"`
	Savings                string `json:"savings"`
	FormattedSavings       string `json:"formatted_savings"`
	SavingsPercent         string `json:"savings_percent"`
}
