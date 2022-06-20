/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CustomerPasswordResetUpdate struct for CustomerPasswordResetUpdate
type CustomerPasswordResetUpdate struct {
	Data CustomerPasswordResetUpdateData `json:"data"`
}

// NewCustomerPasswordResetUpdate instantiates a new CustomerPasswordResetUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPasswordResetUpdate(data CustomerPasswordResetUpdateData) *CustomerPasswordResetUpdate {
	this := CustomerPasswordResetUpdate{}
	this.Data = data
	return &this
}

// NewCustomerPasswordResetUpdateWithDefaults instantiates a new CustomerPasswordResetUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPasswordResetUpdateWithDefaults() *CustomerPasswordResetUpdate {
	this := CustomerPasswordResetUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *CustomerPasswordResetUpdate) GetData() CustomerPasswordResetUpdateData {
	if o == nil {
		var ret CustomerPasswordResetUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomerPasswordResetUpdate) GetDataOk() (*CustomerPasswordResetUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomerPasswordResetUpdate) SetData(v CustomerPasswordResetUpdateData) {
	o.Data = v
}

func (o CustomerPasswordResetUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerPasswordResetUpdate struct {
	value *CustomerPasswordResetUpdate
	isSet bool
}

func (v NullableCustomerPasswordResetUpdate) Get() *CustomerPasswordResetUpdate {
	return v.value
}

func (v *NullableCustomerPasswordResetUpdate) Set(val *CustomerPasswordResetUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPasswordResetUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPasswordResetUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPasswordResetUpdate(val *CustomerPasswordResetUpdate) *NullableCustomerPasswordResetUpdate {
	return &NullableCustomerPasswordResetUpdate{value: val, isSet: true}
}

func (v NullableCustomerPasswordResetUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPasswordResetUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}