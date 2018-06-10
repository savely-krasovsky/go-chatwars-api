package cwapi

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"sync"
)

const (
	CW2 = "amqps://%s:%s@api.chatwars.me:5673/"
	CW3 = "amqps://%s:%s@api.chtwrs.com:5673/"
)

type ActionEnum string

const (
	CreateAuthCode           ActionEnum = "createAuthCode"
	GrantToken               ActionEnum = "grantToken"
	AuthAdditionalOperation  ActionEnum = "authAdditionalOperation"
	GrantAdditionalOperation ActionEnum = "grantAdditionalOperation"
	AuthorizePayment         ActionEnum = "authorizePayment"
	Pay                      ActionEnum = "pay"
	Payout                   ActionEnum = "payout"
	GetInfo                  ActionEnum = "getInfo"
	RequestProfile           ActionEnum = "requestProfile"
	RequestStock             ActionEnum = "requestStock"
	WantToBuy                ActionEnum = "wantToBuy"

	// Unknown action for this lib, check Chat Wars docs
	UnknownAction ActionEnum = "unknownAction"
)

// Returns constant with ActionEnum
func (res *Response) GetActionEnum() ActionEnum {
	switch res.Action {
	case "createAuthCode":
		return CreateAuthCode
	case "grantToken":
		return GrantToken
	case "authAdditionalOperation":
		return AuthAdditionalOperation
	case "grantAdditionalOperation":
		return GrantAdditionalOperation
	case "authorizePayment":
		return AuthorizePayment
	case "pay":
		return Pay
	case "payout":
		return Payout
	case "getInfo":
		return GetInfo
	case "requestProfile":
		return RequestProfile
	case "requestStock":
		return RequestStock
	case "wantToBuy":
		return WantToBuy
	default:
		return UnknownAction
	}
}

type ResultEnum string

const (
	// Everything is Ok
	Ok ResultEnum = "Ok"
	// Amount is either less than or equal zero
	BadAmount ResultEnum = "BadAmount"
	// The currency you chose is not allowed
	BadCurrency ResultEnum = "BadCurrency"
	// Message format is bad. It could be an invalid javascript, or types are wrong, or not all fields are sane
	BadFormat ResultEnum = "BadFormat"
	// The action you have requested is absent. Check spelling
	ActionNotFound ResultEnum = "ActionNotFound"
	// UserID is wrong, or user became inactive
	NoSuchUser ResultEnum = "NoSuchUser"
	// Your app is not yet registered
	NotRegistered ResultEnum = "NotRegistered"
	// Authorization code is incorrect
	InvalidCode ResultEnum = "InvalidCode"
	// If we have some technical difficulties, or bug and are willing for you to repeat request
	TryAgain ResultEnum = "TryAgain"
	// Some field of transaction is bad or confirmation code is wrong
	AuthorizationFailed ResultEnum = "AuthorizationFailed"
	// The player or application balance is insufficient
	InsufficientFunds ResultEnum = "InsufficientFunds"
	// No such token, might be revoked?
	InvalidToken ResultEnum = "InvalidToken"
	// Your app has no rights to execute this action with this token.
	// Payload will contain requiredOperation field.
	// We encourage you to use this field in following authAdditionalOperation, instead of enumerating existing ones
	Forbidden ResultEnum = "Forbidden"

	// Unknown result for this lib, check Chat Wars docs
	UnknownResult ResultEnum = "UnknownResult"
)

// Returns constant with ResultEnum
func (res *Response) GetResultEnum() ResultEnum {
	switch res.Result {
	case "Ok":
		return Ok
	case "BadAmount":
		return BadAmount
	case "BadCurrency":
		return BadCurrency
	case "BadFormat":
		return BadFormat
	case "ActionNotFound":
		return ActionNotFound
	case "NoSuchUser":
		return NoSuchUser
	case "NotRegistered":
		return NotRegistered
	case "InvalidCode":
		return InvalidCode
	case "TryAgain":
		return TryAgain
	case "AuthorizationFailed":
		return AuthorizationFailed
	case "InsufficientFunds":
		return InsufficientFunds
	case "InvalidToken":
		return InvalidToken
	case "Forbidden":
		return Forbidden
	default:
		return UnknownResult
	}
}

type Client struct {
	User      string
	Password  string
	Updates   chan Response
	RabbitUrl string

	YellowPages chan []YellowPage
	Deals       chan Deal
	Offers      chan Offer
	SexDigest   chan []SexDigestItem

	waiters           sync.Map
	connection        *amqp.Connection
	channelForUpdates *amqp.Channel
	channelForPublish *amqp.Channel
}

type Deal struct {
	SellerID     string `json:"sellerId"`
	SellerCastle string `json:"sellerCastle"`
	SellerName   string `json:"sellerName"`
	BuyerID      string `json:"buyerId"`
	BuyerCastle  string `json:"buyerCastle"`
	BuyerName    string `json:"buyerName"`
	Item         string `json:"item"`
	Quantity     int    `json:"qty"`
	Price        int    `json:"price"`
}

type Offer struct {
	SellerID     string `json:"sellerId"`
	SellerCastle string `json:"sellerCastle"`
	SellerName   string `json:"sellerName"`
	Item         string `json:"item"`
	Quantity     int    `json:"qty"`
	Price        int    `json:"price"`
}

type SexDigestItem struct {
	Name   string `json:"name"`
	Prices []int  `json:"prices"`
}

type YellowPage struct {
	Link        string      `json:"link"`
	Name        string      `json:"name"`
	OwnerName   string      `json:"ownerName"`
	OwnerCastle string      `json:"ownerCastle"`
	Kind        string      `json:"kind"`
	Mana        int         `json:"mana"`
	Offers      []OfferItem `json:"offers"`
}

type OfferItem struct {
	Item  string `json:"item"`
	Price int    `json:"price"`
	Mana  int    `json:"mana"`
}

type Request struct {
	Token   string          `json:"token"`
	Action  string          `json:"action"`
	Payload json.RawMessage `json:"payload"`
}

type reqPayload struct {
	*reqCreateAuthCode
	*reqGrantToken
	*reqAuthAdditionalOperation
	*reqGrantAdditionalOperation
	*reqAuthorizePayment
	*reqPay
	*reqPayout
	*reqWantToBuy
}

type Response struct {
	UUID    string     `json:"uuid"`
	Action  string     `json:"action"`
	Result  string     `json:"result"`
	Payload resPayload `json:"payload"`
}

type resPayload struct {
	RequiredOperation string `json:"requiredOperation"`
	*ResCreateAuthCode
	*ResGrantToken
	*ResAuthAdditionalOperation
	*ResGrantAdditionalOperation
	*ResAuthorizePayment
	*ResPay
	*ResPayout
	*ResGetInfo
	*ResRequestProfile
	*ResRequestStock
	*ResWantToBuy
}

type reqCreateAuthCode struct {
	UserID int `json:"userId"`
}

type ResCreateAuthCode struct {
	UserID int `json:"userId"`
}

type reqGrantToken struct {
	UserID   int    `json:"userId"`
	AuthCode string `json:"authCode"`
}

type ResGrantToken struct {
	UserID int    `json:"userId"`
	ID     string `json:"id"`
	Token  string `json:"token"`
}

type reqAuthAdditionalOperation struct {
	Operation string `json:"operation"`
}

type ResAuthAdditionalOperation struct {
	Operation string `json:"operation"`
	UserID    int    `json:"userId"`
}

type reqGrantAdditionalOperation struct {
	RequestID string `json:"requestId"`
	AuthCode  string `json:"authCode"`
}

type ResGrantAdditionalOperation struct {
	RequestID string `json:"requestId"`
	UserID    int    `json:"userId"`
}

type reqAuthorizePayment struct {
	TransactionID string         `json:"transactionId"`
	Amount        map[string]int `json:"amount"`
}

type ResAuthorizePayment struct {
	Fee    map[string]int `json:"fee"`
	Debit  map[string]int `json:"debit"`
	UserID int            `json:"userId"`
}

type reqPay struct {
	TransactionID    string         `json:"transactionId"`
	Amount           map[string]int `json:"amount"`
	ConfirmationCode string         `json:"confirmationCode"`
}

type ResPay struct {
	Fee    map[string]int `json:"fee"`
	Debit  map[string]int `json:"debit"`
	UserID int            `json:"userId"`
}

type reqPayout struct {
	TransactionID string         `json:"transactionId"`
	Amount        map[string]int `json:"amount"`
	Message       string         `json:"message"`
}

type ResPayout struct {
	UserID int `json:"userId"`
}

type ResGetInfo struct {
	Balance int `json:"balance"`
}

type ResRequestProfile struct {
	Profile Profile `json:"profile"`
	UserID  int     `json:"userId"`
	Token   string  `json:"token"`
}

type Profile struct {
	UserName   string `json:"userName"`
	Castle     string `json:"castle"`
	Level      int    `json:"lvl"`
	Experience int    `json:"exp"`
	Attack     int    `json:"atk"`
	Defense    int    `json:"def"`
	Gold       int    `json:"gold"`
	Pouches    int    `json:"pouches"`
	Guild      string `json:"guild"`
	Class      string `json:"class"`
	Mana       int    `json:"mana"`
	Stamina    int    `json:"stamina"`
}

type ResRequestStock struct {
	Stock  map[string]int `json:"stock"`
	UserID int            `json:"userId"`
}

type reqWantToBuy struct {
	ItemCode   string `json:"itemCode"`
	Quantity   int    `json:"quantity"`
	Price      int    `json:"price"`
	ExactPrice bool   `json:"exactPrice"`
}

type ResWantToBuy struct {
	ItemCode string `json:"itemCode"`
	Quantity int    `json:"quantity"`
	UserID   int    `json:"userId"`
}
