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

// KlarnaGateway struct for KlarnaGateway
type KlarnaGateway struct {
	Data KlarnaGatewayData `json:"data"`
}

// NewKlarnaGateway instantiates a new KlarnaGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKlarnaGateway(data KlarnaGatewayData) *KlarnaGateway {
	this := KlarnaGateway{}
	this.Data = data
	return &this
}

// NewKlarnaGatewayWithDefaults instantiates a new KlarnaGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKlarnaGatewayWithDefaults() *KlarnaGateway {
	this := KlarnaGateway{}
	return &this
}

// GetData returns the Data field value
func (o *KlarnaGateway) GetData() KlarnaGatewayData {
	if o == nil {
		var ret KlarnaGatewayData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *KlarnaGateway) GetDataOk() (*KlarnaGatewayData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *KlarnaGateway) SetData(v KlarnaGatewayData) {
	o.Data = v
}

func (o KlarnaGateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableKlarnaGateway struct {
	value *KlarnaGateway
	isSet bool
}

func (v NullableKlarnaGateway) Get() *KlarnaGateway {
	return v.value
}

func (v *NullableKlarnaGateway) Set(val *KlarnaGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableKlarnaGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableKlarnaGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKlarnaGateway(val *KlarnaGateway) *NullableKlarnaGateway {
	return &NullableKlarnaGateway{value: val, isSet: true}
}

func (v NullableKlarnaGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKlarnaGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}