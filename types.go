package golambda_helper

type Header struct {
	ContentType              string `json:"Content-Type"`
	AccessControlAllowOrigin string `json:"Access-Control-Allow-Origin"`
	Location                 string `json:"Location"`
}

type Response struct {
	Body       string `json:"body"`
	StatusCode int    `json:"statusCode"`
	Header     Header `json:"headers"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type ReturnObjectShopName struct {
	ShopName ShopName `json:"shopname"`
}

type ReturnObjectShopNames struct {
	ShopName []ShopName `json:"shopnames"`
}

type ShopName struct {
	Id           string `json:"id"`
	FriendlyName string `json:"friendly_name"`
	ShopName     string `json:"shop_name"`
	CreateDate   string `json:"create_date"`
	Deleted      string `json:"deleted"`
}

type Oauth struct {
	ShopName     string `json:"shop_name"`
	Code         string `json:"code"`
	Hmac         string `json:"hmac"`
	InstallState string `json:"install_state"`
	OauthToken   string `json:"oauth_token"`
	State        string `json:"state"`
}
