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

// CustomerPaymentSource struct for CustomerPaymentSource
type CustomerPaymentSource struct {
	Data CustomerPaymentSourceData `json:"data"`
}

// NewCustomerPaymentSource instantiates a new CustomerPaymentSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPaymentSource(data CustomerPaymentSourceData) *CustomerPaymentSource {
	this := CustomerPaymentSource{}
	this.Data = data
	return &this
}

// NewCustomerPaymentSourceWithDefaults instantiates a new CustomerPaymentSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPaymentSourceWithDefaults() *CustomerPaymentSource {
	this := CustomerPaymentSource{}
	return &this
}

// GetData returns the Data field value
func (o *CustomerPaymentSource) GetData() CustomerPaymentSourceData {
	if o == nil {
		var ret CustomerPaymentSourceData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSource) GetDataOk() (*CustomerPaymentSourceData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CustomerPaymentSource) SetData(v CustomerPaymentSourceData) {
	o.Data = v
}

func (o CustomerPaymentSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerPaymentSource struct {
	value *CustomerPaymentSource
	isSet bool
}

func (v NullableCustomerPaymentSource) Get() *CustomerPaymentSource {
	return v.value
}

func (v *NullableCustomerPaymentSource) Set(val *CustomerPaymentSource) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPaymentSource) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPaymentSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPaymentSource(val *CustomerPaymentSource) *NullableCustomerPaymentSource {
	return &NullableCustomerPaymentSource{value: val, isSet: true}
}

func (v NullableCustomerPaymentSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPaymentSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}