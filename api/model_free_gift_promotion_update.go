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

// FreeGiftPromotionUpdate struct for FreeGiftPromotionUpdate
type FreeGiftPromotionUpdate struct {
	Data FreeGiftPromotionUpdateData `json:"data"`
}

// NewFreeGiftPromotionUpdate instantiates a new FreeGiftPromotionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreeGiftPromotionUpdate(data FreeGiftPromotionUpdateData) *FreeGiftPromotionUpdate {
	this := FreeGiftPromotionUpdate{}
	this.Data = data
	return &this
}

// NewFreeGiftPromotionUpdateWithDefaults instantiates a new FreeGiftPromotionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreeGiftPromotionUpdateWithDefaults() *FreeGiftPromotionUpdate {
	this := FreeGiftPromotionUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *FreeGiftPromotionUpdate) GetData() FreeGiftPromotionUpdateData {
	if o == nil {
		var ret FreeGiftPromotionUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FreeGiftPromotionUpdate) GetDataOk() (*FreeGiftPromotionUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *FreeGiftPromotionUpdate) SetData(v FreeGiftPromotionUpdateData) {
	o.Data = v
}

func (o FreeGiftPromotionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableFreeGiftPromotionUpdate struct {
	value *FreeGiftPromotionUpdate
	isSet bool
}

func (v NullableFreeGiftPromotionUpdate) Get() *FreeGiftPromotionUpdate {
	return v.value
}

func (v *NullableFreeGiftPromotionUpdate) Set(val *FreeGiftPromotionUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableFreeGiftPromotionUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableFreeGiftPromotionUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreeGiftPromotionUpdate(val *FreeGiftPromotionUpdate) *NullableFreeGiftPromotionUpdate {
	return &NullableFreeGiftPromotionUpdate{value: val, isSet: true}
}

func (v NullableFreeGiftPromotionUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreeGiftPromotionUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
