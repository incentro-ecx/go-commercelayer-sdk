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

// ShippingCategoryCreate struct for ShippingCategoryCreate
type ShippingCategoryCreate struct {
	Data ShippingCategoryCreateData `json:"data"`
}

// NewShippingCategoryCreate instantiates a new ShippingCategoryCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingCategoryCreate(data ShippingCategoryCreateData) *ShippingCategoryCreate {
	this := ShippingCategoryCreate{}
	this.Data = data
	return &this
}

// NewShippingCategoryCreateWithDefaults instantiates a new ShippingCategoryCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingCategoryCreateWithDefaults() *ShippingCategoryCreate {
	this := ShippingCategoryCreate{}
	return &this
}

// GetData returns the Data field value
func (o *ShippingCategoryCreate) GetData() ShippingCategoryCreateData {
	if o == nil {
		var ret ShippingCategoryCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ShippingCategoryCreate) GetDataOk() (*ShippingCategoryCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ShippingCategoryCreate) SetData(v ShippingCategoryCreateData) {
	o.Data = v
}

func (o ShippingCategoryCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableShippingCategoryCreate struct {
	value *ShippingCategoryCreate
	isSet bool
}

func (v NullableShippingCategoryCreate) Get() *ShippingCategoryCreate {
	return v.value
}

func (v *NullableShippingCategoryCreate) Set(val *ShippingCategoryCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingCategoryCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingCategoryCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingCategoryCreate(val *ShippingCategoryCreate) *NullableShippingCategoryCreate {
	return &NullableShippingCategoryCreate{value: val, isSet: true}
}

func (v NullableShippingCategoryCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingCategoryCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}