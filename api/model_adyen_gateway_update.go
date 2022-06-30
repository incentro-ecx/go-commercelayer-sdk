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

// AdyenGatewayUpdate struct for AdyenGatewayUpdate
type AdyenGatewayUpdate struct {
	Data AdyenGatewayUpdateData `json:"data"`
}

// NewAdyenGatewayUpdate instantiates a new AdyenGatewayUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenGatewayUpdate(data AdyenGatewayUpdateData) *AdyenGatewayUpdate {
	this := AdyenGatewayUpdate{}
	this.Data = data
	return &this
}

// NewAdyenGatewayUpdateWithDefaults instantiates a new AdyenGatewayUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenGatewayUpdateWithDefaults() *AdyenGatewayUpdate {
	this := AdyenGatewayUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *AdyenGatewayUpdate) GetData() AdyenGatewayUpdateData {
	if o == nil {
		var ret AdyenGatewayUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AdyenGatewayUpdate) GetDataOk() (*AdyenGatewayUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AdyenGatewayUpdate) SetData(v AdyenGatewayUpdateData) {
	o.Data = v
}

func (o AdyenGatewayUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAdyenGatewayUpdate struct {
	value *AdyenGatewayUpdate
	isSet bool
}

func (v NullableAdyenGatewayUpdate) Get() *AdyenGatewayUpdate {
	return v.value
}

func (v *NullableAdyenGatewayUpdate) Set(val *AdyenGatewayUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenGatewayUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenGatewayUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenGatewayUpdate(val *AdyenGatewayUpdate) *NullableAdyenGatewayUpdate {
	return &NullableAdyenGatewayUpdate{value: val, isSet: true}
}

func (v NullableAdyenGatewayUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenGatewayUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
