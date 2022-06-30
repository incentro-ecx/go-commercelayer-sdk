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

// InventoryStockLocationUpdate struct for InventoryStockLocationUpdate
type InventoryStockLocationUpdate struct {
	Data InventoryStockLocationUpdateData `json:"data"`
}

// NewInventoryStockLocationUpdate instantiates a new InventoryStockLocationUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryStockLocationUpdate(data InventoryStockLocationUpdateData) *InventoryStockLocationUpdate {
	this := InventoryStockLocationUpdate{}
	this.Data = data
	return &this
}

// NewInventoryStockLocationUpdateWithDefaults instantiates a new InventoryStockLocationUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryStockLocationUpdateWithDefaults() *InventoryStockLocationUpdate {
	this := InventoryStockLocationUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *InventoryStockLocationUpdate) GetData() InventoryStockLocationUpdateData {
	if o == nil {
		var ret InventoryStockLocationUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InventoryStockLocationUpdate) GetDataOk() (*InventoryStockLocationUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InventoryStockLocationUpdate) SetData(v InventoryStockLocationUpdateData) {
	o.Data = v
}

func (o InventoryStockLocationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryStockLocationUpdate struct {
	value *InventoryStockLocationUpdate
	isSet bool
}

func (v NullableInventoryStockLocationUpdate) Get() *InventoryStockLocationUpdate {
	return v.value
}

func (v *NullableInventoryStockLocationUpdate) Set(val *InventoryStockLocationUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryStockLocationUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryStockLocationUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryStockLocationUpdate(val *InventoryStockLocationUpdate) *NullableInventoryStockLocationUpdate {
	return &NullableInventoryStockLocationUpdate{value: val, isSet: true}
}

func (v NullableInventoryStockLocationUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryStockLocationUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
