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

// StripeGatewayResponseList struct for StripeGatewayResponseList
type StripeGatewayResponseList struct {
	Data []Data `json:"data,omitempty"`
}

// NewStripeGatewayResponseList instantiates a new StripeGatewayResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeGatewayResponseList() *StripeGatewayResponseList {
	this := StripeGatewayResponseList{}
	return &this
}

// NewStripeGatewayResponseListWithDefaults instantiates a new StripeGatewayResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeGatewayResponseListWithDefaults() *StripeGatewayResponseList {
	this := StripeGatewayResponseList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *StripeGatewayResponseList) GetData() []Data {
	if o == nil || o.Data == nil {
		var ret []Data
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeGatewayResponseList) GetDataOk() ([]Data, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *StripeGatewayResponseList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Data and assigns it to the Data field.
func (o *StripeGatewayResponseList) SetData(v []Data) {
	o.Data = v
}

func (o StripeGatewayResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStripeGatewayResponseList struct {
	value *StripeGatewayResponseList
	isSet bool
}

func (v NullableStripeGatewayResponseList) Get() *StripeGatewayResponseList {
	return v.value
}

func (v *NullableStripeGatewayResponseList) Set(val *StripeGatewayResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeGatewayResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeGatewayResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeGatewayResponseList(val *StripeGatewayResponseList) *NullableStripeGatewayResponseList {
	return &NullableStripeGatewayResponseList{value: val, isSet: true}
}

func (v NullableStripeGatewayResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeGatewayResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}