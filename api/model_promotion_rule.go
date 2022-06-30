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

// PromotionRule struct for PromotionRule
type PromotionRule struct {
	Data PromotionRuleData `json:"data"`
}

// NewPromotionRule instantiates a new PromotionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotionRule(data PromotionRuleData) *PromotionRule {
	this := PromotionRule{}
	this.Data = data
	return &this
}

// NewPromotionRuleWithDefaults instantiates a new PromotionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionRuleWithDefaults() *PromotionRule {
	this := PromotionRule{}
	return &this
}

// GetData returns the Data field value
func (o *PromotionRule) GetData() PromotionRuleData {
	if o == nil {
		var ret PromotionRuleData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PromotionRule) GetDataOk() (*PromotionRuleData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PromotionRule) SetData(v PromotionRuleData) {
	o.Data = v
}

func (o PromotionRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePromotionRule struct {
	value *PromotionRule
	isSet bool
}

func (v NullablePromotionRule) Get() *PromotionRule {
	return v.value
}

func (v *NullablePromotionRule) Set(val *PromotionRule) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotionRule) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotionRule(val *PromotionRule) *NullablePromotionRule {
	return &NullablePromotionRule{value: val, isSet: true}
}

func (v NullablePromotionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
