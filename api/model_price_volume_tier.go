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

// PriceVolumeTier struct for PriceVolumeTier
type PriceVolumeTier struct {
	Data PriceVolumeTierData `json:"data"`
}

// NewPriceVolumeTier instantiates a new PriceVolumeTier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceVolumeTier(data PriceVolumeTierData) *PriceVolumeTier {
	this := PriceVolumeTier{}
	this.Data = data
	return &this
}

// NewPriceVolumeTierWithDefaults instantiates a new PriceVolumeTier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceVolumeTierWithDefaults() *PriceVolumeTier {
	this := PriceVolumeTier{}
	return &this
}

// GetData returns the Data field value
func (o *PriceVolumeTier) GetData() PriceVolumeTierData {
	if o == nil {
		var ret PriceVolumeTierData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PriceVolumeTier) GetDataOk() (*PriceVolumeTierData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PriceVolumeTier) SetData(v PriceVolumeTierData) {
	o.Data = v
}

func (o PriceVolumeTier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePriceVolumeTier struct {
	value *PriceVolumeTier
	isSet bool
}

func (v NullablePriceVolumeTier) Get() *PriceVolumeTier {
	return v.value
}

func (v *NullablePriceVolumeTier) Set(val *PriceVolumeTier) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceVolumeTier) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceVolumeTier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceVolumeTier(val *PriceVolumeTier) *NullablePriceVolumeTier {
	return &NullablePriceVolumeTier{value: val, isSet: true}
}

func (v NullablePriceVolumeTier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceVolumeTier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}