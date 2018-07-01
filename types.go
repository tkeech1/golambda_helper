package golambdahelper

// Header is a generic struct to represent an HTTP header
type Header struct {
	ContentType              string `json:"Content-Type"`
	AccessControlAllowOrigin string `json:"Access-Control-Allow-Origin"`
	Location                 string `json:"Location"`
}

// Response is a generic struct to represent an HTTP response
type Response struct {
	Body       string `json:"body"`
	StatusCode int    `json:"statusCode"`
	Header     Header `json:"headers"`
}

// ErrorResponse is a generic struct to represent an HTTP error response
type ErrorResponse struct {
	Message string `json:"message"`
}

// ReturnObjectShopName is a generic struct to represent an ShopName
type ReturnObjectShopName struct {
	ShopName ShopName `json:"shopname"`
}

// ReturnObjectShopNames is a generic struct to represent a slice of ShopNames
type ReturnObjectShopNames struct {
	ShopName []ShopName `json:"shopnames"`
}

// ShopName is a generic struct to represent a single ShopName
type ShopName struct {
	ID           string `json:"id"`
	FriendlyName string `json:"friendly_name,omitempty"`
	ShopName     string `json:"shop_name,omitempty"`
	CreateDate   string `json:"create_date,omitempty"`
	Deleted      string `json:"deleted,omitempty"`
}
