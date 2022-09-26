/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ExternalPaymentResponseDataRelationshipsWallet struct for ExternalPaymentResponseDataRelationshipsWallet
type ExternalPaymentResponseDataRelationshipsWallet struct {
	Links *AddressResponseDataRelationshipsGeocoderLinks      `json:"links,omitempty"`
	Data  *ExternalPaymentResponseDataRelationshipsWalletData `json:"data,omitempty"`
}

// NewExternalPaymentResponseDataRelationshipsWallet instantiates a new ExternalPaymentResponseDataRelationshipsWallet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPaymentResponseDataRelationshipsWallet() *ExternalPaymentResponseDataRelationshipsWallet {
	this := ExternalPaymentResponseDataRelationshipsWallet{}
	return &this
}

// NewExternalPaymentResponseDataRelationshipsWalletWithDefaults instantiates a new ExternalPaymentResponseDataRelationshipsWallet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPaymentResponseDataRelationshipsWalletWithDefaults() *ExternalPaymentResponseDataRelationshipsWallet {
	this := ExternalPaymentResponseDataRelationshipsWallet{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ExternalPaymentResponseDataRelationshipsWallet) GetLinks() AddressResponseDataRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPaymentResponseDataRelationshipsWallet) GetLinksOk() (*AddressResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ExternalPaymentResponseDataRelationshipsWallet) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *ExternalPaymentResponseDataRelationshipsWallet) SetLinks(v AddressResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ExternalPaymentResponseDataRelationshipsWallet) GetData() ExternalPaymentResponseDataRelationshipsWalletData {
	if o == nil || o.Data == nil {
		var ret ExternalPaymentResponseDataRelationshipsWalletData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPaymentResponseDataRelationshipsWallet) GetDataOk() (*ExternalPaymentResponseDataRelationshipsWalletData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ExternalPaymentResponseDataRelationshipsWallet) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ExternalPaymentResponseDataRelationshipsWalletData and assigns it to the Data field.
func (o *ExternalPaymentResponseDataRelationshipsWallet) SetData(v ExternalPaymentResponseDataRelationshipsWalletData) {
	o.Data = &v
}

func (o ExternalPaymentResponseDataRelationshipsWallet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableExternalPaymentResponseDataRelationshipsWallet struct {
	value *ExternalPaymentResponseDataRelationshipsWallet
	isSet bool
}

func (v NullableExternalPaymentResponseDataRelationshipsWallet) Get() *ExternalPaymentResponseDataRelationshipsWallet {
	return v.value
}

func (v *NullableExternalPaymentResponseDataRelationshipsWallet) Set(val *ExternalPaymentResponseDataRelationshipsWallet) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPaymentResponseDataRelationshipsWallet) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPaymentResponseDataRelationshipsWallet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPaymentResponseDataRelationshipsWallet(val *ExternalPaymentResponseDataRelationshipsWallet) *NullableExternalPaymentResponseDataRelationshipsWallet {
	return &NullableExternalPaymentResponseDataRelationshipsWallet{value: val, isSet: true}
}

func (v NullableExternalPaymentResponseDataRelationshipsWallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPaymentResponseDataRelationshipsWallet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}