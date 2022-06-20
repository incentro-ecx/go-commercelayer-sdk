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

// ShippingCategory struct for ShippingCategory
type ShippingCategory struct {
	Data ShippingCategoryData `json:"data"`
}

// NewShippingCategory instantiates a new ShippingCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingCategory(data ShippingCategoryData) *ShippingCategory {
	this := ShippingCategory{}
	this.Data = data
	return &this
}

// NewShippingCategoryWithDefaults instantiates a new ShippingCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingCategoryWithDefaults() *ShippingCategory {
	this := ShippingCategory{}
	return &this
}

// GetData returns the Data field value
func (o *ShippingCategory) GetData() ShippingCategoryData {
	if o == nil {
		var ret ShippingCategoryData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ShippingCategory) GetDataOk() (*ShippingCategoryData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ShippingCategory) SetData(v ShippingCategoryData) {
	o.Data = v
}

func (o ShippingCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableShippingCategory struct {
	value *ShippingCategory
	isSet bool
}

func (v NullableShippingCategory) Get() *ShippingCategory {
	return v.value
}

func (v *NullableShippingCategory) Set(val *ShippingCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingCategory(val *ShippingCategory) *NullableShippingCategory {
	return &NullableShippingCategory{value: val, isSet: true}
}

func (v NullableShippingCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}