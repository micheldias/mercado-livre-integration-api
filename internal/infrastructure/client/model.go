package client

import (
	"time"
)

type CacheAuth struct {
	ClientID     string
	ClientSecret string
	ExpireIn     time.Time
	AccessToken  string
	RefreshToken string
}
type AuthTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	UserId       int    `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
}

type Error struct {
	Message          string   `json:"message"`
	Error            string   `json:"error"`
	ErrorDescription string   `json:"error_description"`
	Status           int      `json:"status"`
	Cause            []string `json:"cause"`
}

type Address struct {
	State   string `json:"state"`
	City    string `json:"city"`
	Address string `json:"address"`
	ZipCode string `json:"zip_code"`
}
type Identification struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Phone struct {
	AreaCode  string `json:"area_code"`
	Number    string `json:"number"`
	Extension string `json:"extension"`
	Verified  bool   `json:"verified"`
}

type Transaction struct {
	Period    string  `json:"period"`
	Total     int     `json:"total"`
	Completed int     `json:"completed"`
	Canceled  int     `json:"canceled"`
	Ratings   Ratings `json:"ratings"`
}
type Ratings struct {
	Positive int `json:"positive"`
	Negative int `json:"negative"`
	Neutral  int `json:"neutral"`
}
type SalesReputation struct {
	LevelID           string      `json:"level_id"`
	PowerSellerStatus string      `json:"power_seller_status"`
	Transactions      Transaction `json:"transactions"`
}

type BuyerReputation struct {
	CanceledTransactions int         `json:"canceled_transactions"`
	Transactions         Transaction `json:"transactions"`
	Tags                 []string    `json:"tags"`
}

type Credit struct {
	Consumed      int    `json:"consumed"`
	CreditLevelID string `json:"credit_level_id"`
}

type User struct {
	ID               int             `json:"id"`
	Nickname         string          `json:"nickname"`
	RegistrationDate string          `json:"registration_date"`
	FirstName        string          `json:"first_name"`
	LastName         string          `json:"last_name"`
	CountryID        string          `json:"country_id"`
	Email            string          `json:"email"`
	Identification   Identification  `json:"identification"`
	Address          Address         `json:"address"`
	Phone            Phone           `json:"phone"`
	AlternativePhone Phone           `json:"alternative_phone"`
	UserType         string          `json:"user_type"`
	Tags             []string        `json:"tags"`
	Logo             string          `json:"logo"`
	Points           int             `json:"points"`
	SiteID           string          `json:"site_id"`
	Permalink        string          `json:"permalink"`
	SellerExperience string          `json:"seller_experience"`
	SellerReputation SalesReputation `json:"seller_reputation"`
	BuyerReputation  BuyerReputation `json:"buyer_reputation"`
	Status           Status          `json:"status"`
	Credit           Credit          `json:"credit"`
}

type Status struct {
	SiteStatus             string  `json:"site_status"`
	List                   List    `json:"list"`
	Buy                    Buy     `json:"buy"`
	Sell                   Sell    `json:"sell"`
	Billing                Billing `json:"billing"`
	MercadopagoTcAccepted  bool    `json:"mercadopago_tc_accepted"`
	MercadopagoAccountType string  `json:"mercadopago_account_type"`
	Mercadoenvios          string  `json:"mercadoenvios"`
	ImmediatePayment       bool    `json:"immediate_payment"`
	ConfirmedEmail         bool    `json:"confirmed_email"`
	UserType               string  `json:"user_type"`
	RequiredAction         string  `json:"required_action"`
}

type Buy struct {
	Allow            bool             `json:"allow"`
	Codes            []interface{}    `json:"codes"`
	ImmediatePayment ImmediatePayment `json:"immediate_payment"`
}

type Sell struct {
	Allow            bool             `json:"allow"`
	Codes            []interface{}    `json:"codes"`
	ImmediatePayment ImmediatePayment `json:"immediate_payment"`
}

type ImmediatePayment struct {
	Required bool          `json:"required"`
	Reasons  []interface{} `json:"reasons"`
}
type Billing struct {
	Allow bool          `json:"allow"`
	Codes []interface{} `json:"codes"`
}

type List struct {
	Allow            bool             `json:"allow"`
	Codes            []interface{}    `json:"codes"`
	ImmediatePayment ImmediatePayment `json:"immediate_payment"`
}
type Sites []Site

type Site struct {
	DefaultCurrencyID string `json:"default_currency_id"`
	ID                string `json:"id"`
	Name              string `json:"name"`
}
type Categories []Category
type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProductRequest struct {
	Title             string      `json:"title"`
	CategoryID        string      `json:"category_id"`
	Price             int         `json:"price"`
	CurrencyID        string      `json:"currency_id"`
	AvailableQuantity int         `json:"available_quantity"`
	BuyingMode        string      `json:"buying_mode"`
	Condition         string      `json:"condition"`
	ListingTypeID     string      `json:"listing_type_id"`
	SaleTerms         []SaleTerm  `json:"sale_terms"`
	Pictures          []Picture   `json:"pictures"`
	Attributes        []Attribute `json:"attributes"`
}

type ProductResponse struct {
	ID                           string        `json:"id"`
	SiteID                       string        `json:"site_id"`
	Title                        string        `json:"title"`
	SellerID                     int           `json:"seller_id"`
	CategoryID                   string        `json:"category_id"`
	OfficialStoreID              interface{}   `json:"official_store_id"`
	Price                        int           `json:"price"`
	BasePrice                    int           `json:"base_price"`
	OriginalPrice                interface{}   `json:"original_price"`
	InventoryID                  interface{}   `json:"inventory_id"`
	CurrencyID                   string        `json:"currency_id"`
	InitialQuantity              int           `json:"initial_quantity"`
	AvailableQuantity            int           `json:"available_quantity"`
	SoldQuantity                 int           `json:"sold_quantity"`
	SaleTerms                    []SaleTerm    `json:"sale_terms"`
	BuyingMode                   string        `json:"buying_mode"`
	ListingTypeID                string        `json:"listing_type_id"`
	StartTime                    time.Time     `json:"start_time"`
	StopTime                     time.Time     `json:"stop_time"`
	EndTime                      time.Time     `json:"end_time"`
	ExpirationTime               time.Time     `json:"expiration_time"`
	Condition                    string        `json:"condition"`
	Permalink                    string        `json:"permalink"`
	Pictures                     []Picture     `json:"pictures"`
	VideoID                      string        `json:"video_id"`
	Descriptions                 []interface{} `json:"descriptions"`
	AcceptsMercadopago           bool          `json:"accepts_mercadopago"`
	NonMercadoPagoPaymentMethods []interface{} `json:"non_mercado_pago_payment_methods"`
	Shipping                     Shipping      `json:"shipping"`
	InternationalDeliveryMode    string        `json:"international_delivery_mode"`
	SellerAddress                struct {
		ID          int    `json:"id"`
		Comment     string `json:"comment"`
		AddressLine string `json:"address_line"`
		ZipCode     string `json:"zip_code"`
		City        struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"city"`
		State struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"state"`
		Country struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
		Latitude       float64 `json:"latitude"`
		Longitude      float64 `json:"longitude"`
		SearchLocation struct {
			Neighborhood struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"neighborhood"`
			City struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"city"`
			State struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"state"`
		} `json:"search_location"`
	} `json:"seller_address"`
	SellerContact     interface{}   `json:"seller_contact"`
	Location          interface{}   `json:"location"`
	Geolocation       Geolocation   `json:"geolocation"`
	CoverageAreas     []interface{} `json:"coverage_areas"`
	Attributes        []Attribute   `json:"attributes"`
	ListingSource     string        `json:"listing_source"`
	Variations        []interface{} `json:"variations"`
	ThumbnailID       string        `json:"thumbnail_id"`
	Thumbnail         string        `json:"thumbnail"`
	Status            string        `json:"status"`
	SubStatus         []interface{} `json:"sub_status"`
	Tags              []string      `json:"tags"`
	Warranty          string        `json:"warranty"`
	CatalogProductID  interface{}   `json:"catalog_product_id"`
	DomainID          string        `json:"domain_id"`
	SellerCustomField interface{}   `json:"seller_custom_field"`
	ParentItemID      interface{}   `json:"parent_item_id"`
	DealIds           []interface{} `json:"deal_ids"`
	AutomaticRelist   bool          `json:"automatic_relist"`
	DateCreated       time.Time     `json:"date_created"`
	LastUpdated       time.Time     `json:"last_updated"`
	Health            interface{}   `json:"health"`
	CatalogListing    bool          `json:"catalog_listing"`
	ItemRelations     []interface{} `json:"item_relations"`
}

type Attribute struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ValueID   string `json:"value_id"`
	ValueName string `json:"value_name"`
}

type Picture struct {
	Source    string `json:"source"`
	ID        string `json:"id"`
	URL       string `json:"url"`
	SecureURL string `json:"secure_url"`
	Size      string `json:"size"`
	MaxSize   string `json:"max_size"`
	Quality   string `json:"quality"`
}

type SaleTerm struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ValueID   string `json:"value_id"`
	ValueName string `json:"value_name"`
}

type Shipping struct {
	Mode         string        `json:"mode"`
	LocalPickUp  bool          `json:"local_pick_up"`
	FreeShipping bool          `json:"free_shipping"`
	Methods      []interface{} `json:"methods"`
	Dimensions   interface{}   `json:"dimensions"`
	Tags         []interface{} `json:"tags"`
	LogisticType string        `json:"logistic_type"`
	StorePickUp  bool          `json:"store_pick_up"`
}

type Geolocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type SellerAddress struct {
	ID             int           `json:"id"`
	Comment        string        `json:"comment"`
	AddressLine    string        `json:"address_line"`
	ZipCode        string        `json:"zip_code"`
	City           SellerCity    `json:"city"`
	State          SellerState   `json:"state"`
	Country        SellerCountry `json:"country"`
	Latitude       float64       `json:"latitude"`
	Longitude      float64       `json:"longitude"`
	SearchLocation struct {
		Neighborhood SellerNeighborhood `json:"neighborhood"`
		City         SellerCity         `json:"city"`
		State        SellerState        `json:"state"`
	} `json:"search_location"`
}

type SellerNeighborhood struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type SellerCity struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SellerState struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SellerCountry struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
