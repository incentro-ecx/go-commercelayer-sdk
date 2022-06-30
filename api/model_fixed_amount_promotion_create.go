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

// FixedAmountPromotionCreate struct for FixedAmountPromotionCreate
type FixedAmountPromotionCreate struct {
	Data FixedAmountPromotionCreateData `json:"data"`
}

// NewFixedAmountPromotionCreate instantiates a new FixedAmountPromotionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFixedAmountPromotionCreate(data FixedAmountPromotionCreateData) *FixedAmountPromotionCreate {
	this := FixedAmountPromotionCreate{}
	this.Data = data
	return &this
}

// NewFixedAmountPromotionCreateWithDefaults instantiates a new FixedAmountPromotionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFixedAmountPromotionCreateWithDefaults() *FixedAmountPromotionCreate {
	this := FixedAmountPromotionCreate{}
	return &this
}

// GetData returns the Data field value
func (o *FixedAmountPromotionCreate) GetData() FixedAmountPromotionCreateData {
	if o == nil {
		var ret FixedAmountPromotionCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FixedAmountPromotionCreate) GetDataOk() (*FixedAmountPromotionCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *FixedAmountPromotionCreate) SetData(v FixedAmountPromotionCreateData) {
	o.Data = v
}

func (o FixedAmountPromotionCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableFixedAmountPromotionCreate struct {
	value *FixedAmountPromotionCreate
	isSet bool
}

func (v NullableFixedAmountPromotionCreate) Get() *FixedAmountPromotionCreate {
	return v.value
}

func (v *NullableFixedAmountPromotionCreate) Set(val *FixedAmountPromotionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableFixedAmountPromotionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableFixedAmountPromotionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFixedAmountPromotionCreate(val *FixedAmountPromotionCreate) *NullableFixedAmountPromotionCreate {
	return &NullableFixedAmountPromotionCreate{value: val, isSet: true}
}

func (v NullableFixedAmountPromotionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFixedAmountPromotionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
