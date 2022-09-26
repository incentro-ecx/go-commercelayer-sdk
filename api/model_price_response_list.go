/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PriceResponseList struct for PriceResponseList
type PriceResponseList struct {
	Data []Data `json:"data,omitempty"`
}

// NewPriceResponseList instantiates a new PriceResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceResponseList() *PriceResponseList {
	this := PriceResponseList{}
	return &this
}

// NewPriceResponseListWithDefaults instantiates a new PriceResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceResponseListWithDefaults() *PriceResponseList {
	this := PriceResponseList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PriceResponseList) GetData() []Data {
	if o == nil || o.Data == nil {
		var ret []Data
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceResponseList) GetDataOk() ([]Data, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PriceResponseList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Data and assigns it to the Data field.
func (o *PriceResponseList) SetData(v []Data) {
	o.Data = v
}

func (o PriceResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePriceResponseList struct {
	value *PriceResponseList
	isSet bool
}

func (v NullablePriceResponseList) Get() *PriceResponseList {
	return v.value
}

func (v *NullablePriceResponseList) Set(val *PriceResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceResponseList(val *PriceResponseList) *NullablePriceResponseList {
	return &NullablePriceResponseList{value: val, isSet: true}
}

func (v NullablePriceResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}