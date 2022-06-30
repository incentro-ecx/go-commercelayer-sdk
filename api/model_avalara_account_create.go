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

// AvalaraAccountCreate struct for AvalaraAccountCreate
type AvalaraAccountCreate struct {
	Data AvalaraAccountCreateData `json:"data"`
}

// NewAvalaraAccountCreate instantiates a new AvalaraAccountCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvalaraAccountCreate(data AvalaraAccountCreateData) *AvalaraAccountCreate {
	this := AvalaraAccountCreate{}
	this.Data = data
	return &this
}

// NewAvalaraAccountCreateWithDefaults instantiates a new AvalaraAccountCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvalaraAccountCreateWithDefaults() *AvalaraAccountCreate {
	this := AvalaraAccountCreate{}
	return &this
}

// GetData returns the Data field value
func (o *AvalaraAccountCreate) GetData() AvalaraAccountCreateData {
	if o == nil {
		var ret AvalaraAccountCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AvalaraAccountCreate) GetDataOk() (*AvalaraAccountCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AvalaraAccountCreate) SetData(v AvalaraAccountCreateData) {
	o.Data = v
}

func (o AvalaraAccountCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAvalaraAccountCreate struct {
	value *AvalaraAccountCreate
	isSet bool
}

func (v NullableAvalaraAccountCreate) Get() *AvalaraAccountCreate {
	return v.value
}

func (v *NullableAvalaraAccountCreate) Set(val *AvalaraAccountCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAvalaraAccountCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAvalaraAccountCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvalaraAccountCreate(val *AvalaraAccountCreate) *NullableAvalaraAccountCreate {
	return &NullableAvalaraAccountCreate{value: val, isSet: true}
}

func (v NullableAvalaraAccountCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvalaraAccountCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
