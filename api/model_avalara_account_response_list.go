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

// AvalaraAccountResponseList struct for AvalaraAccountResponseList
type AvalaraAccountResponseList struct {
	Data []Data `json:"data,omitempty"`
}

// NewAvalaraAccountResponseList instantiates a new AvalaraAccountResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvalaraAccountResponseList() *AvalaraAccountResponseList {
	this := AvalaraAccountResponseList{}
	return &this
}

// NewAvalaraAccountResponseListWithDefaults instantiates a new AvalaraAccountResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvalaraAccountResponseListWithDefaults() *AvalaraAccountResponseList {
	this := AvalaraAccountResponseList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AvalaraAccountResponseList) GetData() []Data {
	if o == nil || o.Data == nil {
		var ret []Data
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccountResponseList) GetDataOk() ([]Data, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AvalaraAccountResponseList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Data and assigns it to the Data field.
func (o *AvalaraAccountResponseList) SetData(v []Data) {
	o.Data = v
}

func (o AvalaraAccountResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAvalaraAccountResponseList struct {
	value *AvalaraAccountResponseList
	isSet bool
}

func (v NullableAvalaraAccountResponseList) Get() *AvalaraAccountResponseList {
	return v.value
}

func (v *NullableAvalaraAccountResponseList) Set(val *AvalaraAccountResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableAvalaraAccountResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableAvalaraAccountResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvalaraAccountResponseList(val *AvalaraAccountResponseList) *NullableAvalaraAccountResponseList {
	return &NullableAvalaraAccountResponseList{value: val, isSet: true}
}

func (v NullableAvalaraAccountResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvalaraAccountResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}