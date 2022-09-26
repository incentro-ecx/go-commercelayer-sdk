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

// VoidResponseList struct for VoidResponseList
type VoidResponseList struct {
	Data []Data `json:"data,omitempty"`
}

// NewVoidResponseList instantiates a new VoidResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoidResponseList() *VoidResponseList {
	this := VoidResponseList{}
	return &this
}

// NewVoidResponseListWithDefaults instantiates a new VoidResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoidResponseListWithDefaults() *VoidResponseList {
	this := VoidResponseList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *VoidResponseList) GetData() []Data {
	if o == nil || o.Data == nil {
		var ret []Data
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoidResponseList) GetDataOk() ([]Data, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *VoidResponseList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Data and assigns it to the Data field.
func (o *VoidResponseList) SetData(v []Data) {
	o.Data = v
}

func (o VoidResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableVoidResponseList struct {
	value *VoidResponseList
	isSet bool
}

func (v NullableVoidResponseList) Get() *VoidResponseList {
	return v.value
}

func (v *NullableVoidResponseList) Set(val *VoidResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableVoidResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableVoidResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoidResponseList(val *VoidResponseList) *NullableVoidResponseList {
	return &NullableVoidResponseList{value: val, isSet: true}
}

func (v NullableVoidResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoidResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}