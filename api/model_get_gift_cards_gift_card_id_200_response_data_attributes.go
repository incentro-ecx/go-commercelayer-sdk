/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETGiftCardsGiftCardId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETGiftCardsGiftCardId200ResponseDataAttributes{}

// GETGiftCardsGiftCardId200ResponseDataAttributes struct for GETGiftCardsGiftCardId200ResponseDataAttributes
type GETGiftCardsGiftCardId200ResponseDataAttributes struct {
	// The gift card status. One of 'draft' (default), 'inactive', 'active', or 'redeemed'.
	Status interface{} `json:"status,omitempty"`
	// The gift card code UUID. If not set, it's automatically generated.
	Code interface{} `json:"code,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode interface{} `json:"currency_code,omitempty"`
	// The gift card initial balance, in cents.
	InitialBalanceCents interface{} `json:"initial_balance_cents,omitempty"`
	// The gift card initial balance, float.
	InitialBalanceFloat interface{} `json:"initial_balance_float,omitempty"`
	// The gift card initial balance, formatted.
	FormattedInitialBalance interface{} `json:"formatted_initial_balance,omitempty"`
	// The gift card balance, in cents.
	BalanceCents interface{} `json:"balance_cents,omitempty"`
	// The gift card balance, float.
	BalanceFloat interface{} `json:"balance_float,omitempty"`
	// The gift card balance, formatted.
	FormattedBalance interface{} `json:"formatted_balance,omitempty"`
	// The gift card balance max, in cents.
	BalanceMaxCents interface{} `json:"balance_max_cents,omitempty"`
	// The gift card balance max, float.
	BalanceMaxFloat interface{} `json:"balance_max_float,omitempty"`
	// The gift card balance max, formatted.
	FormattedBalanceMax interface{} `json:"formatted_balance_max,omitempty"`
	// The gift card balance log. Tracks all the gift card transactions.
	BalanceLog interface{} `json:"balance_log,omitempty"`
	// The gift card usage log. Tracks all the gift card usage actions by orders.
	UsageLog interface{} `json:"usage_log,omitempty"`
	// Indicates if the gift card can be used only one.
	SingleUse interface{} `json:"single_use,omitempty"`
	// Indicates if the gift card can be recharged.
	Rechargeable interface{} `json:"rechargeable,omitempty"`
	// Indicates if redeemed gift card amount is distributed for tax calculation.
	DistributeDiscount interface{} `json:"distribute_discount,omitempty"`
	// The URL of an image that represents the gift card.
	ImageUrl interface{} `json:"image_url,omitempty"`
	// Time at which the gift card will expire.
	ExpiresAt interface{} `json:"expires_at,omitempty"`
	// The email address of the associated recipient. When creating or updating a gift card, this is a shortcut to find or create the associated recipient by email.
	RecipientEmail interface{} `json:"recipient_email,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewGETGiftCardsGiftCardId200ResponseDataAttributes instantiates a new GETGiftCardsGiftCardId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETGiftCardsGiftCardId200ResponseDataAttributes() *GETGiftCardsGiftCardId200ResponseDataAttributes {
	this := GETGiftCardsGiftCardId200ResponseDataAttributes{}
	return &this
}

// NewGETGiftCardsGiftCardId200ResponseDataAttributesWithDefaults instantiates a new GETGiftCardsGiftCardId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETGiftCardsGiftCardId200ResponseDataAttributesWithDefaults() *GETGiftCardsGiftCardId200ResponseDataAttributes {
	this := GETGiftCardsGiftCardId200ResponseDataAttributes{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetStatus() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasStatus() bool {
	if o != nil && IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given interface{} and assigns it to the Status field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetStatus(v interface{}) {
	o.Status = v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return &o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasCode() bool {
	if o != nil && IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given interface{} and assigns it to the Code field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetCode(v interface{}) {
	o.Code = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetCurrencyCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasCurrencyCode() bool {
	if o != nil && IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given interface{} and assigns it to the CurrencyCode field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetCurrencyCode(v interface{}) {
	o.CurrencyCode = v
}

// GetInitialBalanceCents returns the InitialBalanceCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetInitialBalanceCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.InitialBalanceCents
}

// GetInitialBalanceCentsOk returns a tuple with the InitialBalanceCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetInitialBalanceCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.InitialBalanceCents) {
		return nil, false
	}
	return &o.InitialBalanceCents, true
}

// HasInitialBalanceCents returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasInitialBalanceCents() bool {
	if o != nil && IsNil(o.InitialBalanceCents) {
		return true
	}

	return false
}

// SetInitialBalanceCents gets a reference to the given interface{} and assigns it to the InitialBalanceCents field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetInitialBalanceCents(v interface{}) {
	o.InitialBalanceCents = v
}

// GetInitialBalanceFloat returns the InitialBalanceFloat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetInitialBalanceFloat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.InitialBalanceFloat
}

// GetInitialBalanceFloatOk returns a tuple with the InitialBalanceFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetInitialBalanceFloatOk() (*interface{}, bool) {
	if o == nil || IsNil(o.InitialBalanceFloat) {
		return nil, false
	}
	return &o.InitialBalanceFloat, true
}

// HasInitialBalanceFloat returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasInitialBalanceFloat() bool {
	if o != nil && IsNil(o.InitialBalanceFloat) {
		return true
	}

	return false
}

// SetInitialBalanceFloat gets a reference to the given interface{} and assigns it to the InitialBalanceFloat field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetInitialBalanceFloat(v interface{}) {
	o.InitialBalanceFloat = v
}

// GetFormattedInitialBalance returns the FormattedInitialBalance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetFormattedInitialBalance() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FormattedInitialBalance
}

// GetFormattedInitialBalanceOk returns a tuple with the FormattedInitialBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetFormattedInitialBalanceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FormattedInitialBalance) {
		return nil, false
	}
	return &o.FormattedInitialBalance, true
}

// HasFormattedInitialBalance returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasFormattedInitialBalance() bool {
	if o != nil && IsNil(o.FormattedInitialBalance) {
		return true
	}

	return false
}

// SetFormattedInitialBalance gets a reference to the given interface{} and assigns it to the FormattedInitialBalance field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetFormattedInitialBalance(v interface{}) {
	o.FormattedInitialBalance = v
}

// GetBalanceCents returns the BalanceCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.BalanceCents
}

// GetBalanceCentsOk returns a tuple with the BalanceCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.BalanceCents) {
		return nil, false
	}
	return &o.BalanceCents, true
}

// HasBalanceCents returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasBalanceCents() bool {
	if o != nil && IsNil(o.BalanceCents) {
		return true
	}

	return false
}

// SetBalanceCents gets a reference to the given interface{} and assigns it to the BalanceCents field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceCents(v interface{}) {
	o.BalanceCents = v
}

// GetBalanceFloat returns the BalanceFloat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceFloat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.BalanceFloat
}

// GetBalanceFloatOk returns a tuple with the BalanceFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceFloatOk() (*interface{}, bool) {
	if o == nil || IsNil(o.BalanceFloat) {
		return nil, false
	}
	return &o.BalanceFloat, true
}

// HasBalanceFloat returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasBalanceFloat() bool {
	if o != nil && IsNil(o.BalanceFloat) {
		return true
	}

	return false
}

// SetBalanceFloat gets a reference to the given interface{} and assigns it to the BalanceFloat field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceFloat(v interface{}) {
	o.BalanceFloat = v
}

// GetFormattedBalance returns the FormattedBalance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetFormattedBalance() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FormattedBalance
}

// GetFormattedBalanceOk returns a tuple with the FormattedBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetFormattedBalanceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FormattedBalance) {
		return nil, false
	}
	return &o.FormattedBalance, true
}

// HasFormattedBalance returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasFormattedBalance() bool {
	if o != nil && IsNil(o.FormattedBalance) {
		return true
	}

	return false
}

// SetFormattedBalance gets a reference to the given interface{} and assigns it to the FormattedBalance field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetFormattedBalance(v interface{}) {
	o.FormattedBalance = v
}

// GetBalanceMaxCents returns the BalanceMaxCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceMaxCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.BalanceMaxCents
}

// GetBalanceMaxCentsOk returns a tuple with the BalanceMaxCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceMaxCentsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.BalanceMaxCents) {
		return nil, false
	}
	return &o.BalanceMaxCents, true
}

// HasBalanceMaxCents returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasBalanceMaxCents() bool {
	if o != nil && IsNil(o.BalanceMaxCents) {
		return true
	}

	return false
}

// SetBalanceMaxCents gets a reference to the given interface{} and assigns it to the BalanceMaxCents field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceMaxCents(v interface{}) {
	o.BalanceMaxCents = v
}

// GetBalanceMaxFloat returns the BalanceMaxFloat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceMaxFloat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.BalanceMaxFloat
}

// GetBalanceMaxFloatOk returns a tuple with the BalanceMaxFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceMaxFloatOk() (*interface{}, bool) {
	if o == nil || IsNil(o.BalanceMaxFloat) {
		return nil, false
	}
	return &o.BalanceMaxFloat, true
}

// HasBalanceMaxFloat returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasBalanceMaxFloat() bool {
	if o != nil && IsNil(o.BalanceMaxFloat) {
		return true
	}

	return false
}

// SetBalanceMaxFloat gets a reference to the given interface{} and assigns it to the BalanceMaxFloat field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceMaxFloat(v interface{}) {
	o.BalanceMaxFloat = v
}

// GetFormattedBalanceMax returns the FormattedBalanceMax field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetFormattedBalanceMax() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FormattedBalanceMax
}

// GetFormattedBalanceMaxOk returns a tuple with the FormattedBalanceMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetFormattedBalanceMaxOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FormattedBalanceMax) {
		return nil, false
	}
	return &o.FormattedBalanceMax, true
}

// HasFormattedBalanceMax returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasFormattedBalanceMax() bool {
	if o != nil && IsNil(o.FormattedBalanceMax) {
		return true
	}

	return false
}

// SetFormattedBalanceMax gets a reference to the given interface{} and assigns it to the FormattedBalanceMax field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetFormattedBalanceMax(v interface{}) {
	o.FormattedBalanceMax = v
}

// GetBalanceLog returns the BalanceLog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceLog() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.BalanceLog
}

// GetBalanceLogOk returns a tuple with the BalanceLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetBalanceLogOk() (*interface{}, bool) {
	if o == nil || IsNil(o.BalanceLog) {
		return nil, false
	}
	return &o.BalanceLog, true
}

// HasBalanceLog returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasBalanceLog() bool {
	if o != nil && IsNil(o.BalanceLog) {
		return true
	}

	return false
}

// SetBalanceLog gets a reference to the given interface{} and assigns it to the BalanceLog field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetBalanceLog(v interface{}) {
	o.BalanceLog = v
}

// GetUsageLog returns the UsageLog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetUsageLog() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UsageLog
}

// GetUsageLogOk returns a tuple with the UsageLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetUsageLogOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UsageLog) {
		return nil, false
	}
	return &o.UsageLog, true
}

// HasUsageLog returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasUsageLog() bool {
	if o != nil && IsNil(o.UsageLog) {
		return true
	}

	return false
}

// SetUsageLog gets a reference to the given interface{} and assigns it to the UsageLog field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetUsageLog(v interface{}) {
	o.UsageLog = v
}

// GetSingleUse returns the SingleUse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetSingleUse() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SingleUse
}

// GetSingleUseOk returns a tuple with the SingleUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetSingleUseOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SingleUse) {
		return nil, false
	}
	return &o.SingleUse, true
}

// HasSingleUse returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasSingleUse() bool {
	if o != nil && IsNil(o.SingleUse) {
		return true
	}

	return false
}

// SetSingleUse gets a reference to the given interface{} and assigns it to the SingleUse field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetSingleUse(v interface{}) {
	o.SingleUse = v
}

// GetRechargeable returns the Rechargeable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetRechargeable() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Rechargeable
}

// GetRechargeableOk returns a tuple with the Rechargeable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetRechargeableOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Rechargeable) {
		return nil, false
	}
	return &o.Rechargeable, true
}

// HasRechargeable returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasRechargeable() bool {
	if o != nil && IsNil(o.Rechargeable) {
		return true
	}

	return false
}

// SetRechargeable gets a reference to the given interface{} and assigns it to the Rechargeable field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetRechargeable(v interface{}) {
	o.Rechargeable = v
}

// GetDistributeDiscount returns the DistributeDiscount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetDistributeDiscount() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DistributeDiscount
}

// GetDistributeDiscountOk returns a tuple with the DistributeDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetDistributeDiscountOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DistributeDiscount) {
		return nil, false
	}
	return &o.DistributeDiscount, true
}

// HasDistributeDiscount returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasDistributeDiscount() bool {
	if o != nil && IsNil(o.DistributeDiscount) {
		return true
	}

	return false
}

// SetDistributeDiscount gets a reference to the given interface{} and assigns it to the DistributeDiscount field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetDistributeDiscount(v interface{}) {
	o.DistributeDiscount = v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetImageUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetImageUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return &o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasImageUrl() bool {
	if o != nil && IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given interface{} and assigns it to the ImageUrl field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetImageUrl(v interface{}) {
	o.ImageUrl = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetExpiresAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasExpiresAt() bool {
	if o != nil && IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given interface{} and assigns it to the ExpiresAt field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetExpiresAt(v interface{}) {
	o.ExpiresAt = v
}

// GetRecipientEmail returns the RecipientEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetRecipientEmail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RecipientEmail
}

// GetRecipientEmailOk returns a tuple with the RecipientEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetRecipientEmailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RecipientEmail) {
		return nil, false
	}
	return &o.RecipientEmail, true
}

// HasRecipientEmail returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasRecipientEmail() bool {
	if o != nil && IsNil(o.RecipientEmail) {
		return true
	}

	return false
}

// SetRecipientEmail gets a reference to the given interface{} and assigns it to the RecipientEmail field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetRecipientEmail(v interface{}) {
	o.RecipientEmail = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETGiftCardsGiftCardId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETGiftCardsGiftCardId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETGiftCardsGiftCardId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.InitialBalanceCents != nil {
		toSerialize["initial_balance_cents"] = o.InitialBalanceCents
	}
	if o.InitialBalanceFloat != nil {
		toSerialize["initial_balance_float"] = o.InitialBalanceFloat
	}
	if o.FormattedInitialBalance != nil {
		toSerialize["formatted_initial_balance"] = o.FormattedInitialBalance
	}
	if o.BalanceCents != nil {
		toSerialize["balance_cents"] = o.BalanceCents
	}
	if o.BalanceFloat != nil {
		toSerialize["balance_float"] = o.BalanceFloat
	}
	if o.FormattedBalance != nil {
		toSerialize["formatted_balance"] = o.FormattedBalance
	}
	if o.BalanceMaxCents != nil {
		toSerialize["balance_max_cents"] = o.BalanceMaxCents
	}
	if o.BalanceMaxFloat != nil {
		toSerialize["balance_max_float"] = o.BalanceMaxFloat
	}
	if o.FormattedBalanceMax != nil {
		toSerialize["formatted_balance_max"] = o.FormattedBalanceMax
	}
	if o.BalanceLog != nil {
		toSerialize["balance_log"] = o.BalanceLog
	}
	if o.UsageLog != nil {
		toSerialize["usage_log"] = o.UsageLog
	}
	if o.SingleUse != nil {
		toSerialize["single_use"] = o.SingleUse
	}
	if o.Rechargeable != nil {
		toSerialize["rechargeable"] = o.Rechargeable
	}
	if o.DistributeDiscount != nil {
		toSerialize["distribute_discount"] = o.DistributeDiscount
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.RecipientEmail != nil {
		toSerialize["recipient_email"] = o.RecipientEmail
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableGETGiftCardsGiftCardId200ResponseDataAttributes struct {
	value *GETGiftCardsGiftCardId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETGiftCardsGiftCardId200ResponseDataAttributes) Get() *GETGiftCardsGiftCardId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETGiftCardsGiftCardId200ResponseDataAttributes) Set(val *GETGiftCardsGiftCardId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETGiftCardsGiftCardId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETGiftCardsGiftCardId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETGiftCardsGiftCardId200ResponseDataAttributes(val *GETGiftCardsGiftCardId200ResponseDataAttributes) *NullableGETGiftCardsGiftCardId200ResponseDataAttributes {
	return &NullableGETGiftCardsGiftCardId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETGiftCardsGiftCardId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETGiftCardsGiftCardId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
