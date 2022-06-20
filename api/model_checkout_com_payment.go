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

// CheckoutComPayment struct for CheckoutComPayment
type CheckoutComPayment struct {
	Data CheckoutComPaymentData `json:"data"`
}

// NewCheckoutComPayment instantiates a new CheckoutComPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutComPayment(data CheckoutComPaymentData) *CheckoutComPayment {
	this := CheckoutComPayment{}
	this.Data = data
	return &this
}

// NewCheckoutComPaymentWithDefaults instantiates a new CheckoutComPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutComPaymentWithDefaults() *CheckoutComPayment {
	this := CheckoutComPayment{}
	return &this
}

// GetData returns the Data field value
func (o *CheckoutComPayment) GetData() CheckoutComPaymentData {
	if o == nil {
		var ret CheckoutComPaymentData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CheckoutComPayment) GetDataOk() (*CheckoutComPaymentData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CheckoutComPayment) SetData(v CheckoutComPaymentData) {
	o.Data = v
}

func (o CheckoutComPayment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCheckoutComPayment struct {
	value *CheckoutComPayment
	isSet bool
}

func (v NullableCheckoutComPayment) Get() *CheckoutComPayment {
	return v.value
}

func (v *NullableCheckoutComPayment) Set(val *CheckoutComPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutComPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutComPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutComPayment(val *CheckoutComPayment) *NullableCheckoutComPayment {
	return &NullableCheckoutComPayment{value: val, isSet: true}
}

func (v NullableCheckoutComPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutComPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}