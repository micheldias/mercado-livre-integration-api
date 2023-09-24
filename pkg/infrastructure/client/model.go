package client

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
