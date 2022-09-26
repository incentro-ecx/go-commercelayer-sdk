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

// CustomerPaymentSourceResponseList struct for CustomerPaymentSourceResponseList
type CustomerPaymentSourceResponseList struct {
	Data []Data `json:"data,omitempty"`
}

// NewCustomerPaymentSourceResponseList instantiates a new CustomerPaymentSourceResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerPaymentSourceResponseList() *CustomerPaymentSourceResponseList {
	this := CustomerPaymentSourceResponseList{}
	return &this
}

// NewCustomerPaymentSourceResponseListWithDefaults instantiates a new CustomerPaymentSourceResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerPaymentSourceResponseListWithDefaults() *CustomerPaymentSourceResponseList {
	this := CustomerPaymentSourceResponseList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomerPaymentSourceResponseList) GetData() []Data {
	if o == nil || o.Data == nil {
		var ret []Data
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerPaymentSourceResponseList) GetDataOk() ([]Data, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomerPaymentSourceResponseList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Data and assigns it to the Data field.
func (o *CustomerPaymentSourceResponseList) SetData(v []Data) {
	o.Data = v
}

func (o CustomerPaymentSourceResponseList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerPaymentSourceResponseList struct {
	value *CustomerPaymentSourceResponseList
	isSet bool
}

func (v NullableCustomerPaymentSourceResponseList) Get() *CustomerPaymentSourceResponseList {
	return v.value
}

func (v *NullableCustomerPaymentSourceResponseList) Set(val *CustomerPaymentSourceResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerPaymentSourceResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerPaymentSourceResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerPaymentSourceResponseList(val *CustomerPaymentSourceResponseList) *NullableCustomerPaymentSourceResponseList {
	return &NullableCustomerPaymentSourceResponseList{value: val, isSet: true}
}

func (v NullableCustomerPaymentSourceResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerPaymentSourceResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}