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

// InventoryModelCreate struct for InventoryModelCreate
type InventoryModelCreate struct {
	Data InventoryModelCreateData `json:"data"`
}

// NewInventoryModelCreate instantiates a new InventoryModelCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryModelCreate(data InventoryModelCreateData) *InventoryModelCreate {
	this := InventoryModelCreate{}
	this.Data = data
	return &this
}

// NewInventoryModelCreateWithDefaults instantiates a new InventoryModelCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryModelCreateWithDefaults() *InventoryModelCreate {
	this := InventoryModelCreate{}
	return &this
}

// GetData returns the Data field value
func (o *InventoryModelCreate) GetData() InventoryModelCreateData {
	if o == nil {
		var ret InventoryModelCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InventoryModelCreate) GetDataOk() (*InventoryModelCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InventoryModelCreate) SetData(v InventoryModelCreateData) {
	o.Data = v
}

func (o InventoryModelCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableInventoryModelCreate struct {
	value *InventoryModelCreate
	isSet bool
}

func (v NullableInventoryModelCreate) Get() *InventoryModelCreate {
	return v.value
}

func (v *NullableInventoryModelCreate) Set(val *InventoryModelCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryModelCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryModelCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryModelCreate(val *InventoryModelCreate) *NullableInventoryModelCreate {
	return &NullableInventoryModelCreate{value: val, isSet: true}
}

func (v NullableInventoryModelCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventoryModelCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
