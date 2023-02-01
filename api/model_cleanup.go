/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.4.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Cleanup struct for Cleanup
type Cleanup struct {
	Data *CleanupData `json:"data,omitempty"`
}

// NewCleanup instantiates a new Cleanup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCleanup() *Cleanup {
	this := Cleanup{}
	return &this
}

// NewCleanupWithDefaults instantiates a new Cleanup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCleanupWithDefaults() *Cleanup {
	this := Cleanup{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Cleanup) GetData() CleanupData {
	if o == nil || o.Data == nil {
		var ret CleanupData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cleanup) GetDataOk() (*CleanupData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Cleanup) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given CleanupData and assigns it to the Data field.
func (o *Cleanup) SetData(v CleanupData) {
	o.Data = &v
}

func (o Cleanup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCleanup struct {
	value *Cleanup
	isSet bool
}

func (v NullableCleanup) Get() *Cleanup {
	return v.value
}

func (v *NullableCleanup) Set(val *Cleanup) {
	v.value = val
	v.isSet = true
}

func (v NullableCleanup) IsSet() bool {
	return v.isSet
}

func (v *NullableCleanup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCleanup(val *Cleanup) *NullableCleanup {
	return &NullableCleanup{value: val, isSet: true}
}

func (v NullableCleanup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCleanup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}