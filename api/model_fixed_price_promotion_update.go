/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the FixedPricePromotionUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FixedPricePromotionUpdate{}

// FixedPricePromotionUpdate struct for FixedPricePromotionUpdate
type FixedPricePromotionUpdate struct {
	Data FixedPricePromotionUpdateData `json:"data"`
}

// NewFixedPricePromotionUpdate instantiates a new FixedPricePromotionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFixedPricePromotionUpdate(data FixedPricePromotionUpdateData) *FixedPricePromotionUpdate {
	this := FixedPricePromotionUpdate{}
	this.Data = data
	return &this
}

// NewFixedPricePromotionUpdateWithDefaults instantiates a new FixedPricePromotionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFixedPricePromotionUpdateWithDefaults() *FixedPricePromotionUpdate {
	this := FixedPricePromotionUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *FixedPricePromotionUpdate) GetData() FixedPricePromotionUpdateData {
	if o == nil {
		var ret FixedPricePromotionUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionUpdate) GetDataOk() (*FixedPricePromotionUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *FixedPricePromotionUpdate) SetData(v FixedPricePromotionUpdateData) {
	o.Data = v
}

func (o FixedPricePromotionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FixedPricePromotionUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableFixedPricePromotionUpdate struct {
	value *FixedPricePromotionUpdate
	isSet bool
}

func (v NullableFixedPricePromotionUpdate) Get() *FixedPricePromotionUpdate {
	return v.value
}

func (v *NullableFixedPricePromotionUpdate) Set(val *FixedPricePromotionUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableFixedPricePromotionUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableFixedPricePromotionUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFixedPricePromotionUpdate(val *FixedPricePromotionUpdate) *NullableFixedPricePromotionUpdate {
	return &NullableFixedPricePromotionUpdate{value: val, isSet: true}
}

func (v NullableFixedPricePromotionUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFixedPricePromotionUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
