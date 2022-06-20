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

// ReturnLineItemCreate struct for ReturnLineItemCreate
type ReturnLineItemCreate struct {
	Data ReturnLineItemCreateData `json:"data"`
}

// NewReturnLineItemCreate instantiates a new ReturnLineItemCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnLineItemCreate(data ReturnLineItemCreateData) *ReturnLineItemCreate {
	this := ReturnLineItemCreate{}
	this.Data = data
	return &this
}

// NewReturnLineItemCreateWithDefaults instantiates a new ReturnLineItemCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnLineItemCreateWithDefaults() *ReturnLineItemCreate {
	this := ReturnLineItemCreate{}
	return &this
}

// GetData returns the Data field value
func (o *ReturnLineItemCreate) GetData() ReturnLineItemCreateData {
	if o == nil {
		var ret ReturnLineItemCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ReturnLineItemCreate) GetDataOk() (*ReturnLineItemCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ReturnLineItemCreate) SetData(v ReturnLineItemCreateData) {
	o.Data = v
}

func (o ReturnLineItemCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableReturnLineItemCreate struct {
	value *ReturnLineItemCreate
	isSet bool
}

func (v NullableReturnLineItemCreate) Get() *ReturnLineItemCreate {
	return v.value
}

func (v *NullableReturnLineItemCreate) Set(val *ReturnLineItemCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnLineItemCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnLineItemCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnLineItemCreate(val *ReturnLineItemCreate) *NullableReturnLineItemCreate {
	return &NullableReturnLineItemCreate{value: val, isSet: true}
}

func (v NullableReturnLineItemCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnLineItemCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}