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

// ReturnLineItemUpdate struct for ReturnLineItemUpdate
type ReturnLineItemUpdate struct {
	Data ReturnLineItemUpdateData `json:"data"`
}

// NewReturnLineItemUpdate instantiates a new ReturnLineItemUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnLineItemUpdate(data ReturnLineItemUpdateData) *ReturnLineItemUpdate {
	this := ReturnLineItemUpdate{}
	this.Data = data
	return &this
}

// NewReturnLineItemUpdateWithDefaults instantiates a new ReturnLineItemUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnLineItemUpdateWithDefaults() *ReturnLineItemUpdate {
	this := ReturnLineItemUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *ReturnLineItemUpdate) GetData() ReturnLineItemUpdateData {
	if o == nil {
		var ret ReturnLineItemUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ReturnLineItemUpdate) GetDataOk() (*ReturnLineItemUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ReturnLineItemUpdate) SetData(v ReturnLineItemUpdateData) {
	o.Data = v
}

func (o ReturnLineItemUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableReturnLineItemUpdate struct {
	value *ReturnLineItemUpdate
	isSet bool
}

func (v NullableReturnLineItemUpdate) Get() *ReturnLineItemUpdate {
	return v.value
}

func (v *NullableReturnLineItemUpdate) Set(val *ReturnLineItemUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnLineItemUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnLineItemUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnLineItemUpdate(val *ReturnLineItemUpdate) *NullableReturnLineItemUpdate {
	return &NullableReturnLineItemUpdate{value: val, isSet: true}
}

func (v NullableReturnLineItemUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnLineItemUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}