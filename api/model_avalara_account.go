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

// checks if the AvalaraAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvalaraAccount{}

// AvalaraAccount struct for AvalaraAccount
type AvalaraAccount struct {
	Data *AvalaraAccountData `json:"data,omitempty"`
}

// NewAvalaraAccount instantiates a new AvalaraAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvalaraAccount() *AvalaraAccount {
	this := AvalaraAccount{}
	return &this
}

// NewAvalaraAccountWithDefaults instantiates a new AvalaraAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvalaraAccountWithDefaults() *AvalaraAccount {
	this := AvalaraAccount{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AvalaraAccount) GetData() AvalaraAccountData {
	if o == nil || IsNil(o.Data) {
		var ret AvalaraAccountData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvalaraAccount) GetDataOk() (*AvalaraAccountData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AvalaraAccount) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AvalaraAccountData and assigns it to the Data field.
func (o *AvalaraAccount) SetData(v AvalaraAccountData) {
	o.Data = &v
}

func (o AvalaraAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvalaraAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAvalaraAccount struct {
	value *AvalaraAccount
	isSet bool
}

func (v NullableAvalaraAccount) Get() *AvalaraAccount {
	return v.value
}

func (v *NullableAvalaraAccount) Set(val *AvalaraAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableAvalaraAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAvalaraAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvalaraAccount(val *AvalaraAccount) *NullableAvalaraAccount {
	return &NullableAvalaraAccount{value: val, isSet: true}
}

func (v NullableAvalaraAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvalaraAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
