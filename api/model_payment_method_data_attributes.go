/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PaymentMethodDataAttributes struct for PaymentMethodDataAttributes
type PaymentMethodDataAttributes struct {
	// The payment source type, can be one of: 'AdyenPayment', 'BraintreePayment', 'CheckoutComPayment', 'CreditCard', 'ExternalPayment', 'PaypalPayment', 'StripePayment', or 'WireTransfer'.
	PaymentSourceType *string `json:"payment_source_type,omitempty"`
	// Payment source type, titleized
	Name *string `json:"name,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// Send this attribute if you want to mark the payment as MOTO, must be supported by payment gateway.
	Moto *bool `json:"moto,omitempty"`
	// Time at which the payment method was disabled.
	DisabledAt *string `json:"disabled_at,omitempty"`
	// The payment method's price, in cents
	PriceAmountCents *int32 `json:"price_amount_cents,omitempty"`
	// The payment method's price, float
	PriceAmountFloat *float32 `json:"price_amount_float,omitempty"`
	// The payment method's price, formatted
	FormattedPriceAmount *string `json:"formatted_price_amount,omitempty"`
	// Unique identifier for the resource (hash).
	Id *string `json:"id,omitempty"`
	// Time at which the resource was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPaymentMethodDataAttributes instantiates a new PaymentMethodDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodDataAttributes() *PaymentMethodDataAttributes {
	this := PaymentMethodDataAttributes{}
	return &this
}

// NewPaymentMethodDataAttributesWithDefaults instantiates a new PaymentMethodDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodDataAttributesWithDefaults() *PaymentMethodDataAttributes {
	this := PaymentMethodDataAttributes{}
	return &this
}

// GetPaymentSourceType returns the PaymentSourceType field value if set, zero value otherwise.
func (o *PaymentMethodDataAttributes) GetPaymentSourceType() string {
	if o == nil || o.PaymentSourceType == nil {
		var ret string
		return ret
	}
	return *o.PaymentSourceType
}

// GetPaymentSourceTypeOk returns a tuple with the PaymentSourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodDataAttributes) GetPaymentSourceTypeOk() (*string, bool) {
	if o == nil || o.PaymentSourceType == nil {
		return nil, false
	}
	return o.PaymentSourceType, true
}

// HasPaymentSourceType returns a boolean if a field has been set.
func (o *PaymentMethodDataAttributes) HasPaymentSourceType() bool {
	if o != nil && o.PaymentSourceType != nil {
		return true
	}

	return false
}

// SetPaymentSourceType gets a reference to the given string and assigns it to the PaymentSourceType field.
func (o *PaymentMethodDataAttributes) SetPaymentSourceType(v string) {
	o.PaymentSourceType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PaymentMethodDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PaymentMethodDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PaymentMethodDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *PaymentMethodDataAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodDataAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *PaymentMethodDataAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *PaymentMethodDataAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetMoto returns the Moto field value if set, zero value otherwise.
func (o *PaymentMethodDataAttributes) GetMoto() bool {
	if o == nil || o.Moto == nil {
		var ret bool
		return ret
	}
	return *o.Moto
}

// GetMotoOk returns a tuple with the Moto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodDataAttributes) GetMotoOk() (*bool, bool) {
	if o == nil || o.Moto == nil {
		return nil, false
	}
	return o.Moto, true
}

// HasMoto returns a boolean if a field has been set.
func (o *PaymentMethodDataAttributes) HasMoto() bool {
	if o != nil && o.Moto != nil {
		return true
	}

	return false
}

// SetMoto gets a reference to the given bool and assigns it to the Moto field.
func (o *PaymentMethodDataAttributes) SetMoto(v bool) {
	o.Moto = &v
}

// GetDisabledAt returns the DisabledAt field value if set, zero value otherwise.
func (o *PaymentMethodDataAttributes) GetDisabledAt() string {
	if o == nil || o.DisabledAt == nil {
		var ret string
		return ret
	}
	return *o.DisabledAt
}

// GetDisabledAtOk returns a tuple with the DisabledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodDataAttributes) GetDisabledAtOk() (*string, bool) {
	if o == nil || o.DisabledAt == nil {
		return nil, false
	}
	return o.DisabledAt, true
}

// HasDisabledAt returns a boolean if a field has been set.
func (o *PaymentMethodDataAttributes) HasDisabledAt() bool {
	if o != nil && o.DisabledAt != nil {
		return true
	}

	return false
}

// SetDisabledAt gets a reference to the given string and assigns it to the DisabledAt field.
func (o *PaymentMethodDataAttributes) SetDisabledAt(v string) {
	o.DisabledAt = &v
}

// GetPriceAmountCents returns the PriceAmountCents field value if set, zero value otherwise.
func (o *PaymentMethodDataAttributes) GetPriceAmountCents() int32 {
	if o == nil || o.PriceAmountCents == nil {
		var ret int32
		return ret
	}
	return *o.PriceAmountCents
}

// GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodDataAttributes) GetPriceAmountCentsOk() (*int32, bool) {
	if o == nil || o.PriceAmountCents == nil {
		return nil, false
	}
	return o.PriceAmountCents, true
}

// HasPriceAmountCents returns a boolean if a field has been set.
func (o *PaymentMethodDataAttributes) HasPriceAmountCents() bool {
	if o != nil && o.PriceAmountCents != nil {
		return true
	}

	return false
}

// SetPriceAmountCents gets a reference to the given int32 and assigns it to the PriceAmountCents field.
func (o *PaymentMethodDataAttributes) SetPriceAmountCents(v int32) {
	o.PriceAmountCents = &v
}

// GetPriceAmountFloat returns the PriceAmountFloat field value if set, zero value otherwise.
func (o *PaymentMethodDataAttributes) GetPriceAmountFloat() float32 {
	if o == nil || o.PriceAmountFloat == nil {
		var ret float32
		return ret
	}
	return *o.PriceAmountFloat
}

// GetPriceAmountFloatOk returns a tuple with the PriceAmountFloat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodDataAttributes) GetPriceAmountFloatOk() (*float32, bool) {
	if o == nil || o.PriceAmountFloat == nil {
		return nil, false
	}
	return o.PriceAmountFloat, true
}

// HasPriceAmountFloat returns a boolean if a field has been set.
func (o *PaymentMethodDataAttributes) HasPriceAmountFloat() bool {
	if o != nil && o.PriceAmountFloat != nil {
		return true
	}

	return false
}

// SetPriceAmountFloat gets a reference to the given float32 and assigns it to the PriceAmountFloat field.
func (o *PaymentMethodDataAttributes) SetPriceAmountFloat(v float32) {
	o.PriceAmountFloat = &v
}

// GetFormattedPriceAmount returns the FormattedPriceAmount field value if set, zero value otherwise.
func (o *PaymentMethodDataAttributes) GetFormattedPriceAmount() string {
	if o == nil || o.FormattedPriceAmount == nil {
		var ret string
		return ret
	}
	return *o.FormattedPriceAmount
}

// GetFormattedPriceAmountOk returns a tuple with the FormattedPriceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodDataAttributes) GetFormattedPriceAmountOk() (*string, bool) {
	if o == nil || o.FormattedPriceAmount == nil {
		return nil, false
	}
	return o.FormattedPriceAmount, true
}

// HasFormattedPriceAmount returns a boolean if a field has been set.
func (o *PaymentMethodDataAttributes) HasFormattedPriceAmount() bool {
	if o != nil && o.FormattedPriceAmount != nil {
		return true
	}

	return false
}

// SetFormattedPriceAmount gets a reference to the given string and assigns it to the FormattedPriceAmount field.
func (o *PaymentMethodDataAttributes) SetFormattedPriceAmount(v string) {
	o.FormattedPriceAmount = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaymentMethodDataAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodDataAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaymentMethodDataAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PaymentMethodDataAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PaymentMethodDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PaymentMethodDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *PaymentMethodDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PaymentMethodDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PaymentMethodDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *PaymentMethodDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PaymentMethodDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PaymentMethodDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PaymentMethodDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PaymentMethodDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PaymentMethodDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PaymentMethodDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PaymentMethodDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PaymentMethodDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PaymentMethodDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PaymentMethodDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PaymentSourceType != nil {
		toSerialize["payment_source_type"] = o.PaymentSourceType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.Moto != nil {
		toSerialize["moto"] = o.Moto
	}
	if o.DisabledAt != nil {
		toSerialize["disabled_at"] = o.DisabledAt
	}
	if o.PriceAmountCents != nil {
		toSerialize["price_amount_cents"] = o.PriceAmountCents
	}
	if o.PriceAmountFloat != nil {
		toSerialize["price_amount_float"] = o.PriceAmountFloat
	}
	if o.FormattedPriceAmount != nil {
		toSerialize["formatted_price_amount"] = o.FormattedPriceAmount
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
	return json.Marshal(toSerialize)
}

type NullablePaymentMethodDataAttributes struct {
	value *PaymentMethodDataAttributes
	isSet bool
}

func (v NullablePaymentMethodDataAttributes) Get() *PaymentMethodDataAttributes {
	return v.value
}

func (v *NullablePaymentMethodDataAttributes) Set(val *PaymentMethodDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodDataAttributes(val *PaymentMethodDataAttributes) *NullablePaymentMethodDataAttributes {
	return &NullablePaymentMethodDataAttributes{value: val, isSet: true}
}

func (v NullablePaymentMethodDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
