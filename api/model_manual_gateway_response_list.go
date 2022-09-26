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

// ManualGatewayResponseList struct for ManualGatewayResponseList
type ManualGatewayResponseList struct {
	Data []Data `json:"data,omitempty"`
}

// NewManualGatewayResponseList instantiates a new ManualGatewayResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualGatewayResponseList() *ManualGatewayResponseList {
	this := ManualGatewayResponseList{}
	return &this
}

// NewManualGatewayResponseListWithDefaults instantiates a new ManualGatewayResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualGatewayResponseListWithDefaults() *ManualGatewayResponseList {
	this := ManualGatewayResponseList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ManualGatewayResponseList) GetData() []Data {
	if o == nil || o.Data == nil {
		var ret []Data
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManualGatewayResponseList) GetDataOk() ([]Data, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ManualGatewayResponseList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Data and assigns it to the Data field.
func (o *ManualGatewayResponseList) SetData(v []Data) {
	o.Data = v
}

func (o ManualGatewayResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableManualGatewayResponseList struct {
	value *ManualGatewayResponseList
	isSet bool
}

func (v NullableManualGatewayResponseList) Get() *ManualGatewayResponseList {
	return v.value
}

func (v *NullableManualGatewayResponseList) Set(val *ManualGatewayResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableManualGatewayResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableManualGatewayResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualGatewayResponseList(val *ManualGatewayResponseList) *NullableManualGatewayResponseList {
	return &NullableManualGatewayResponseList{value: val, isSet: true}
}

func (v NullableManualGatewayResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualGatewayResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}