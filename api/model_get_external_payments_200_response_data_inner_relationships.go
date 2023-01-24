/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETExternalPayments200ResponseDataInnerRelationships struct for GETExternalPayments200ResponseDataInnerRelationships
type GETExternalPayments200ResponseDataInnerRelationships struct {
	Order          *GETAdyenPayments200ResponseDataInnerRelationshipsOrder          `json:"order,omitempty"`
	PaymentGateway *GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway `json:"payment_gateway,omitempty"`
	Wallet         *GETExternalPayments200ResponseDataInnerRelationshipsWallet      `json:"wallet,omitempty"`
}

// NewGETExternalPayments200ResponseDataInnerRelationships instantiates a new GETExternalPayments200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExternalPayments200ResponseDataInnerRelationships() *GETExternalPayments200ResponseDataInnerRelationships {
	this := GETExternalPayments200ResponseDataInnerRelationships{}
	return &this
}

// NewGETExternalPayments200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETExternalPayments200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExternalPayments200ResponseDataInnerRelationshipsWithDefaults() *GETExternalPayments200ResponseDataInnerRelationships {
	this := GETExternalPayments200ResponseDataInnerRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *GETExternalPayments200ResponseDataInnerRelationships) GetOrder() GETAdyenPayments200ResponseDataInnerRelationshipsOrder {
	if o == nil || o.Order == nil {
		var ret GETAdyenPayments200ResponseDataInnerRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalPayments200ResponseDataInnerRelationships) GetOrderOk() (*GETAdyenPayments200ResponseDataInnerRelationshipsOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *GETExternalPayments200ResponseDataInnerRelationships) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given GETAdyenPayments200ResponseDataInnerRelationshipsOrder and assigns it to the Order field.
func (o *GETExternalPayments200ResponseDataInnerRelationships) SetOrder(v GETAdyenPayments200ResponseDataInnerRelationshipsOrder) {
	o.Order = &v
}

// GetPaymentGateway returns the PaymentGateway field value if set, zero value otherwise.
func (o *GETExternalPayments200ResponseDataInnerRelationships) GetPaymentGateway() GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway {
	if o == nil || o.PaymentGateway == nil {
		var ret GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway
		return ret
	}
	return *o.PaymentGateway
}

// GetPaymentGatewayOk returns a tuple with the PaymentGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalPayments200ResponseDataInnerRelationships) GetPaymentGatewayOk() (*GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway, bool) {
	if o == nil || o.PaymentGateway == nil {
		return nil, false
	}
	return o.PaymentGateway, true
}

// HasPaymentGateway returns a boolean if a field has been set.
func (o *GETExternalPayments200ResponseDataInnerRelationships) HasPaymentGateway() bool {
	if o != nil && o.PaymentGateway != nil {
		return true
	}

	return false
}

// SetPaymentGateway gets a reference to the given GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway and assigns it to the PaymentGateway field.
func (o *GETExternalPayments200ResponseDataInnerRelationships) SetPaymentGateway(v GETAdyenPayments200ResponseDataInnerRelationshipsPaymentGateway) {
	o.PaymentGateway = &v
}

// GetWallet returns the Wallet field value if set, zero value otherwise.
func (o *GETExternalPayments200ResponseDataInnerRelationships) GetWallet() GETExternalPayments200ResponseDataInnerRelationshipsWallet {
	if o == nil || o.Wallet == nil {
		var ret GETExternalPayments200ResponseDataInnerRelationshipsWallet
		return ret
	}
	return *o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalPayments200ResponseDataInnerRelationships) GetWalletOk() (*GETExternalPayments200ResponseDataInnerRelationshipsWallet, bool) {
	if o == nil || o.Wallet == nil {
		return nil, false
	}
	return o.Wallet, true
}

// HasWallet returns a boolean if a field has been set.
func (o *GETExternalPayments200ResponseDataInnerRelationships) HasWallet() bool {
	if o != nil && o.Wallet != nil {
		return true
	}

	return false
}

// SetWallet gets a reference to the given GETExternalPayments200ResponseDataInnerRelationshipsWallet and assigns it to the Wallet field.
func (o *GETExternalPayments200ResponseDataInnerRelationships) SetWallet(v GETExternalPayments200ResponseDataInnerRelationshipsWallet) {
	o.Wallet = &v
}

func (o GETExternalPayments200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
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

type NullableGETExternalPayments200ResponseDataInnerRelationships struct {
	value *GETExternalPayments200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETExternalPayments200ResponseDataInnerRelationships) Get() *GETExternalPayments200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETExternalPayments200ResponseDataInnerRelationships) Set(val *GETExternalPayments200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExternalPayments200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExternalPayments200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExternalPayments200ResponseDataInnerRelationships(val *GETExternalPayments200ResponseDataInnerRelationships) *NullableGETExternalPayments200ResponseDataInnerRelationships {
	return &NullableGETExternalPayments200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETExternalPayments200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExternalPayments200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
