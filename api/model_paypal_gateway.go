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

// PaypalGateway struct for PaypalGateway
type PaypalGateway struct {
	Data PaypalGatewayData `json:"data"`
}

// NewPaypalGateway instantiates a new PaypalGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaypalGateway(data PaypalGatewayData) *PaypalGateway {
	this := PaypalGateway{}
	this.Data = data
	return &this
}

// NewPaypalGatewayWithDefaults instantiates a new PaypalGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaypalGatewayWithDefaults() *PaypalGateway {
	this := PaypalGateway{}
	return &this
}

// GetData returns the Data field value
func (o *PaypalGateway) GetData() PaypalGatewayData {
	if o == nil {
		var ret PaypalGatewayData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PaypalGateway) GetDataOk() (*PaypalGatewayData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PaypalGateway) SetData(v PaypalGatewayData) {
	o.Data = v
}

func (o PaypalGateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePaypalGateway struct {
	value *PaypalGateway
	isSet bool
}

func (v NullablePaypalGateway) Get() *PaypalGateway {
	return v.value
}

func (v *NullablePaypalGateway) Set(val *PaypalGateway) {
	v.value = val
	v.isSet = true
}

func (v NullablePaypalGateway) IsSet() bool {
	return v.isSet
}

func (v *NullablePaypalGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaypalGateway(val *PaypalGateway) *NullablePaypalGateway {
	return &NullablePaypalGateway{value: val, isSet: true}
}

func (v NullablePaypalGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaypalGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
