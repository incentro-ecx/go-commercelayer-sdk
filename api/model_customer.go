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

// Customer struct for Customer
type Customer struct {
	Data CustomerData `json:"data"`
}

// NewCustomer instantiates a new Customer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomer(data CustomerData) *Customer {
	this := Customer{}
	this.Data = data
	return &this
}

// NewCustomerWithDefaults instantiates a new Customer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerWithDefaults() *Customer {
	this := Customer{}
	return &this
}

// GetData returns the Data field value
func (o *Customer) GetData() CustomerData {
	if o == nil {
		var ret CustomerData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Customer) GetDataOk() (*CustomerData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Customer) SetData(v CustomerData) {
	o.Data = v
}

func (o Customer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCustomer struct {
	value *Customer
	isSet bool
}

func (v NullableCustomer) Get() *Customer {
	return v.value
}

func (v *NullableCustomer) Set(val *Customer) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomer(val *Customer) *NullableCustomer {
	return &NullableCustomer{value: val, isSet: true}
}

func (v NullableCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}