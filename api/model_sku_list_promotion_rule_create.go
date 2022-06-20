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

// SkuListPromotionRuleCreate struct for SkuListPromotionRuleCreate
type SkuListPromotionRuleCreate struct {
	Data SkuListPromotionRuleCreateData `json:"data"`
}

// NewSkuListPromotionRuleCreate instantiates a new SkuListPromotionRuleCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListPromotionRuleCreate(data SkuListPromotionRuleCreateData) *SkuListPromotionRuleCreate {
	this := SkuListPromotionRuleCreate{}
	this.Data = data
	return &this
}

// NewSkuListPromotionRuleCreateWithDefaults instantiates a new SkuListPromotionRuleCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListPromotionRuleCreateWithDefaults() *SkuListPromotionRuleCreate {
	this := SkuListPromotionRuleCreate{}
	return &this
}

// GetData returns the Data field value
func (o *SkuListPromotionRuleCreate) GetData() SkuListPromotionRuleCreateData {
	if o == nil {
		var ret SkuListPromotionRuleCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleCreate) GetDataOk() (*SkuListPromotionRuleCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SkuListPromotionRuleCreate) SetData(v SkuListPromotionRuleCreateData) {
	o.Data = v
}

func (o SkuListPromotionRuleCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSkuListPromotionRuleCreate struct {
	value *SkuListPromotionRuleCreate
	isSet bool
}

func (v NullableSkuListPromotionRuleCreate) Get() *SkuListPromotionRuleCreate {
	return v.value
}

func (v *NullableSkuListPromotionRuleCreate) Set(val *SkuListPromotionRuleCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListPromotionRuleCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListPromotionRuleCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListPromotionRuleCreate(val *SkuListPromotionRuleCreate) *NullableSkuListPromotionRuleCreate {
	return &NullableSkuListPromotionRuleCreate{value: val, isSet: true}
}

func (v NullableSkuListPromotionRuleCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListPromotionRuleCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}