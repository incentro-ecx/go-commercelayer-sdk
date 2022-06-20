/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ExternalPaymentDataRelationships struct for ExternalPaymentDataRelationships
type ExternalPaymentDataRelationships struct {
	Order          *AdyenPaymentDataRelationshipsOrder              `json:"order,omitempty"`
	PaymentGateway *AdyenPaymentDataRelationshipsPaymentGateway     `json:"payment_gateway,omitempty"`
	Wallet         *CustomerDataRelationshipsCustomerPaymentSources `json:"wallet,omitempty"`
}

// NewExternalPaymentDataRelationships instantiates a new ExternalPaymentDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPaymentDataRelationships() *ExternalPaymentDataRelationships {
	this := ExternalPaymentDataRelationships{}
	return &this
}

// NewExternalPaymentDataRelationshipsWithDefaults instantiates a new ExternalPaymentDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPaymentDataRelationshipsWithDefaults() *ExternalPaymentDataRelationships {
	this := ExternalPaymentDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ExternalPaymentDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder {
	if o == nil || o.Order == nil {
		var ret AdyenPaymentDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPaymentDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ExternalPaymentDataRelationships) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given AdyenPaymentDataRelationshipsOrder and assigns it to the Order field.
func (o *ExternalPaymentDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder) {
	o.Order = &v
}

// GetPaymentGateway returns the PaymentGateway field value if set, zero value otherwise.
func (o *ExternalPaymentDataRelationships) GetPaymentGateway() AdyenPaymentDataRelationshipsPaymentGateway {
	if o == nil || o.PaymentGateway == nil {
		var ret AdyenPaymentDataRelationshipsPaymentGateway
		return ret
	}
	return *o.PaymentGateway
}

// GetPaymentGatewayOk returns a tuple with the PaymentGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPaymentDataRelationships) GetPaymentGatewayOk() (*AdyenPaymentDataRelationshipsPaymentGateway, bool) {
	if o == nil || o.PaymentGateway == nil {
		return nil, false
	}
	return o.PaymentGateway, true
}

// HasPaymentGateway returns a boolean if a field has been set.
func (o *ExternalPaymentDataRelationships) HasPaymentGateway() bool {
	if o != nil && o.PaymentGateway != nil {
		return true
	}

	return false
}

// SetPaymentGateway gets a reference to the given AdyenPaymentDataRelationshipsPaymentGateway and assigns it to the PaymentGateway field.
func (o *ExternalPaymentDataRelationships) SetPaymentGateway(v AdyenPaymentDataRelationshipsPaymentGateway) {
	o.PaymentGateway = &v
}

// GetWallet returns the Wallet field value if set, zero value otherwise.
func (o *ExternalPaymentDataRelationships) GetWallet() CustomerDataRelationshipsCustomerPaymentSources {
	if o == nil || o.Wallet == nil {
		var ret CustomerDataRelationshipsCustomerPaymentSources
		return ret
	}
	return *o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPaymentDataRelationships) GetWalletOk() (*CustomerDataRelationshipsCustomerPaymentSources, bool) {
	if o == nil || o.Wallet == nil {
		return nil, false
	}
	return o.Wallet, true
}

// HasWallet returns a boolean if a field has been set.
func (o *ExternalPaymentDataRelationships) HasWallet() bool {
	if o != nil && o.Wallet != nil {
		return true
	}

	return false
}

// SetWallet gets a reference to the given CustomerDataRelationshipsCustomerPaymentSources and assigns it to the Wallet field.
func (o *ExternalPaymentDataRelationships) SetWallet(v CustomerDataRelationshipsCustomerPaymentSources) {
	o.Wallet = &v
}

func (o ExternalPaymentDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.PaymentGateway != nil {
		toSerialize["payment_gateway"] = o.PaymentGateway
	}
	if o.Wallet != nil {
		toSerialize["wallet"] = o.Wallet
	}
	return json.Marshal(toSerialize)
}

type NullableExternalPaymentDataRelationships struct {
	value *ExternalPaymentDataRelationships
	isSet bool
}

func (v NullableExternalPaymentDataRelationships) Get() *ExternalPaymentDataRelationships {
	return v.value
}

func (v *NullableExternalPaymentDataRelationships) Set(val *ExternalPaymentDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPaymentDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPaymentDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPaymentDataRelationships(val *ExternalPaymentDataRelationships) *NullableExternalPaymentDataRelationships {
	return &NullableExternalPaymentDataRelationships{value: val, isSet: true}
}

func (v NullableExternalPaymentDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPaymentDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}