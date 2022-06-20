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

// OrderAmountPromotionRuleUpdate struct for OrderAmountPromotionRuleUpdate
type OrderAmountPromotionRuleUpdate struct {
	Data OrderAmountPromotionRuleUpdateData `json:"data"`
}

// NewOrderAmountPromotionRuleUpdate instantiates a new OrderAmountPromotionRuleUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAmountPromotionRuleUpdate(data OrderAmountPromotionRuleUpdateData) *OrderAmountPromotionRuleUpdate {
	this := OrderAmountPromotionRuleUpdate{}
	this.Data = data
	return &this
}

// NewOrderAmountPromotionRuleUpdateWithDefaults instantiates a new OrderAmountPromotionRuleUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAmountPromotionRuleUpdateWithDefaults() *OrderAmountPromotionRuleUpdate {
	this := OrderAmountPromotionRuleUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *OrderAmountPromotionRuleUpdate) GetData() OrderAmountPromotionRuleUpdateData {
	if o == nil {
		var ret OrderAmountPromotionRuleUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrderAmountPromotionRuleUpdate) GetDataOk() (*OrderAmountPromotionRuleUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OrderAmountPromotionRuleUpdate) SetData(v OrderAmountPromotionRuleUpdateData) {
	o.Data = v
}

func (o OrderAmountPromotionRuleUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableOrderAmountPromotionRuleUpdate struct {
	value *OrderAmountPromotionRuleUpdate
	isSet bool
}

func (v NullableOrderAmountPromotionRuleUpdate) Get() *OrderAmountPromotionRuleUpdate {
	return v.value
}

func (v *NullableOrderAmountPromotionRuleUpdate) Set(val *OrderAmountPromotionRuleUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAmountPromotionRuleUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAmountPromotionRuleUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAmountPromotionRuleUpdate(val *OrderAmountPromotionRuleUpdate) *NullableOrderAmountPromotionRuleUpdate {
	return &NullableOrderAmountPromotionRuleUpdate{value: val, isSet: true}
}

func (v NullableOrderAmountPromotionRuleUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAmountPromotionRuleUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}